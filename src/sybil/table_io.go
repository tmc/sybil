package sybil

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime/debug"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
)

var DEBUG_TIMING = false
var CHUNKS_BEFORE_GC = 16
var INGEST_DIR = "ingest"
var TEMP_INGEST_DIR = ".ingest.temp"
var CACHE_DIR = "cache"

var HOLD_MATCHES = false
var BLOCKS_PER_CACHE_FILE = 64

func (t *Table) saveTableInfo(fname string) error {
	if err := t.GrabInfoLock(); err != nil {
		return errors.Wrap(err, "t.GrabInfoLock")
	}

	defer t.ReleaseInfoLock()
	var network bytes.Buffer // Stand-in for the network.
	dirname := path.Join(FLAGS.DIR, t.Name)
	filename := path.Join(dirname, fmt.Sprintf("%s.db", fname))
	backup := path.Join(dirname, fmt.Sprintf("%s.bak", fname))

	flagfile := path.Join(dirname, fmt.Sprintf("%s.db.exists", fname))

	// Create a backup file
	cp(backup, filename)

	// Create an encoder and send a value.
	enc := gob.NewEncoder(&network)
	err := enc.Encode(t)

	if err != nil {
		return err
	}

	Debug("SERIALIZED TABLE INFO", fname, "INTO ", network.Len(), "BYTES")

	tempfile, err := ioutil.TempFile(dirname, "info.db")
	if err != nil {
		return errors.Wrap(err, "error creating temp file for table info")
	}

	_, err = network.WriteTo(tempfile)
	if err != nil {
		return errors.Wrap(err, "error saving table info into tempfile")
	}

	RenameAndMod(tempfile.Name(), filename)
	_, err = os.Create(flagfile)
	return err
}

func (t *Table) SaveTableInfo(fname string) error {
	return getSaveTable(t).saveTableInfo(fname)
}

func getSaveTable(t *Table) *Table {
	return &Table{Name: t.Name,
		KeyTable: t.KeyTable,
		KeyTypes: t.KeyTypes,
		IntInfo:  t.IntInfo,
		StrInfo:  t.StrInfo}
}

func (t *Table) saveRecordList(records RecordList) error {
	if len(records) == 0 {
		return nil
	}

	Debug("SAVING RECORD LIST", len(records), t.Name)

	chunkSize := CHUNK_SIZE
	chunks := len(records) / chunkSize

	if chunks == 0 {
		filename, err := t.getNewIngestBlockName()
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("error saving block %v", filename))
		}
		if err := t.SaveRecordsToBlock(records, filename); err != nil {
			return err
		}
	} else {
		for j := 0; j < chunks; j++ {
			filename, err := t.getNewIngestBlockName()
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error saving block %v", filename))
			}
			if err := t.SaveRecordsToBlock(records[j*chunkSize:(j+1)*chunkSize], filename); err != nil {
				return err
			}
		}

		// SAVE THE REMAINDER
		if len(records) > chunks*chunkSize {
			filename, err := t.getNewIngestBlockName()
			if err != nil {
				return errors.Wrap(err, "error creating new ingestion block")
			}

			if err := t.SaveRecordsToBlock(records[chunks*chunkSize:], filename); err != nil {
				return err
			}
		}
	}

	return nil
}

func (t *Table) SaveRecordsToColumns() error {
	os.MkdirAll(path.Join(FLAGS.DIR, t.Name), 0777)
	sort.Sort(SortRecordsByTime{t.newRecords})

	if err := t.FillPartialBlock(); err != nil {
		return err
	}
	ret := t.saveRecordList(t.newRecords)
	t.newRecords = make(RecordList, 0)
	t.SaveTableInfo("info")
	return ret
}

func (t *Table) LoadTableInfo() error {
	tablename := t.Name
	filename := path.Join(FLAGS.DIR, tablename, "info.db")
	if err := t.GrabInfoLock(); err == nil {
		defer t.ReleaseInfoLock()
	} else {
		Debug("LOAD TABLE INFO LOCK TAKEN")
		return err
	}

	return t.LoadTableInfoFrom(filename)
}

func (t *Table) LoadTableInfoFrom(filename string) error {
	savedTable := Table{Name: t.Name}
	savedTable.initDataStructures()

	start := time.Now()

	Debug("OPENING TABLE INFO FROM FILENAME", filename)
	err := decodeInto(filename, &savedTable)
	end := time.Now()
	if err != nil {
		Debug("TABLE INFO DECODE:", err)
		return err
	}

	if DEBUG_TIMING {
		Debug("TABLE INFO OPEN TOOK", end.Sub(start))
	}

	if len(savedTable.KeyTable) > 0 {
		t.KeyTable = savedTable.KeyTable
	}

	if len(savedTable.KeyTypes) > 0 {
		t.KeyTypes = savedTable.KeyTypes
	}

	if savedTable.IntInfo != nil {
		t.IntInfo = savedTable.IntInfo
	}
	if savedTable.StrInfo != nil {
		t.StrInfo = savedTable.StrInfo
	}

	// If we are recovering the INFO lock, we won't necessarily have
	// all fields filled out
	if t.stringIDMu != nil {
		t.populateStringIDLookup()
	}

	return nil
}

// Remove our pointer to the blocklist so a GC is triggered and
// a bunch of new memory becomes available
func (t *Table) ReleaseRecords() {
	t.BlockList = make(map[string]*TableBlock)
	debug.FreeOSMemory()
}

func (t *Table) HasFlagFile() bool {
	// Make a determination of whether this is a new table or not. if it is a
	// new table, we are fine, but if it's not - we are in trouble!
	flagfile := path.Join(FLAGS.DIR, t.Name, "info.db.exists")
	_, err := os.Open(flagfile)
	// If the flagfile exists and we couldn't read the file info, we are in trouble!
	if err == nil {
		t.ReleaseInfoLock()
		Warn("Table info missing, but flag file exists!")
		return true
	}

	return false

}

func fileLooksLikeBlock(v os.FileInfo) bool {

	switch {

	case v.Name() == INGEST_DIR || v.Name() == TEMP_INGEST_DIR:
		return false
	case v.Name() == CACHE_DIR:
		return false
	case strings.HasPrefix(v.Name(), STOMACHE_DIR):
		return false
	case strings.HasSuffix(v.Name(), "info.db"):
		return false
	case strings.HasSuffix(v.Name(), "old"):
		return false
	case strings.HasSuffix(v.Name(), "broken"):
		return false
	case strings.HasSuffix(v.Name(), "lock"):
		return false
	case strings.HasSuffix(v.Name(), "export"):
		return false
	case strings.HasSuffix(v.Name(), "partial"):
		return false
	}

	return true

}

func (t *Table) LoadBlockCache() error {
	_, done := t.trace()
	defer done()
	if err := t.GrabCacheLock(); err != nil {
		return nil
	}

	defer t.ReleaseCacheLock()
	files, err := ioutil.ReadDir(path.Join(FLAGS.DIR, t.Name, CACHE_DIR))

	if err != nil {
		return err
	}

	for _, blockFile := range files {
		filename := path.Join(FLAGS.DIR, t.Name, CACHE_DIR, blockFile.Name())
		blockCache := SavedBlockCache{}

		err = decodeInto(filename, &blockCache)
		if err != nil {
			continue
		}

		for k, v := range blockCache {
			t.BlockInfoCache[k] = v
		}
	}

	Debug("FILLED BLOCK CACHE WITH", len(t.BlockInfoCache), "ITEMS")
	return nil
}

func (t *Table) ResetBlockCache() {
	t.BlockInfoCache = make(map[string]*SavedColumnInfo)
}

func (t *Table) WriteQueryCache(toCacheSpecs map[string]*QuerySpec) {

	// NOW WE SAVE OUR QUERY CACHE HERE...
	savestart := time.Now()
	var wg sync.WaitGroup

	saved := 0

	for blockName, blockQuery := range toCacheSpecs {

		if blockName == INGEST_DIR || len(blockQuery.Results) > 5000 {
			continue
		}
		thisQuery := blockQuery
		thisName := blockName

		wg.Add(1)
		saved++
		go func() {

			thisQuery.SaveCachedResults(thisName)
			if FLAGS.DEBUG {
				fmt.Fprint(os.Stderr, "s")
			}

			wg.Done()
		}()

		wg.Wait()

		saveend := time.Now()

		if saved > 0 {
			if FLAGS.DEBUG {
				fmt.Fprint(os.Stderr, "\n")
			}
			Debug("SAVING CACHED QUERIES TOOK", saveend.Sub(savestart))
		}
	}

	// END QUERY CACHE SAVING

}

func (t *Table) WriteBlockCache() error {
	if len(t.NewBlockInfos) == 0 {
		return nil
	}

	if err := t.GrabCacheLock(); err != nil {
		return err
	}

	defer t.ReleaseCacheLock()

	Debug("WRITING BLOCK CACHE, OUTSTANDING", len(t.NewBlockInfos))

	var numBlocks = len(t.NewBlockInfos) / BLOCKS_PER_CACHE_FILE

	for i := 0; i < numBlocks; i++ {
		cachedInfo := t.NewBlockInfos[i*BLOCKS_PER_CACHE_FILE : (i+1)*BLOCKS_PER_CACHE_FILE]

		blockFile, err := t.getNewCacheBlockFile()
		if err != nil {
			Debug("TROUBLE CREATING CACHE BLOCK FILE")
			break
		}
		blockCache := SavedBlockCache{}

		for _, blockName := range cachedInfo {
			blockCache[blockName] = t.BlockInfoCache[blockName]
		}

		enc := gob.NewEncoder(blockFile)
		err = enc.Encode(&blockCache)
		if err != nil {
			Debug("ERROR ENCODING BLOCK CACHE", err)
		}

		pathname := fmt.Sprintf("%s.db", blockFile.Name())

		Debug("RENAMING", blockFile.Name(), pathname)
		RenameAndMod(blockFile.Name(), pathname)

	}

	t.NewBlockInfos = t.NewBlockInfos[:0]

	return nil
}

func (t *Table) LoadRecords(loadSpec *LoadSpec) (int, error) {
	_, done := t.trace()
	defer done()
	t.LoadBlockCache()

	return t.LoadAndQueryRecords(loadSpec, nil)
}

func (t *Table) ChunkAndSave() error {

	if len(t.newRecords) >= CHUNK_SIZE {
		os.MkdirAll(path.Join(FLAGS.DIR, t.Name), 0777)
		name, err := t.getNewIngestBlockName()
		if err == nil {
			if err := t.SaveRecordsToBlock(t.newRecords, name); err != nil {
				return err
			}
			if err := t.SaveTableInfo("info"); err != nil {
				return err
			}
			t.newRecords = make(RecordList, 0)
			t.ReleaseRecords()
		} else {
			return errors.Wrap(err, "error saving block")
		}
	}

	return nil
}

func (t *Table) IsNotExist() bool {
	// TODO: consider using os.Stat and os.IsNotExist
	tableDir := path.Join(FLAGS.DIR, t.Name)
	_, err := ioutil.ReadDir(tableDir)
	return err != nil
}

package sybil

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"os"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"go.opencensus.io/trace"
)

var GZIP_EXT = ".gz"

func (t *Table) SaveRecordsToBlock(records RecordList, filename string) error {
	if len(records) == 0 {
		return nil
	}

	tempBlock := newTableBlock()
	tempBlock.RecordList = records
	tempBlock.table = t

	return tempBlock.SaveToColumns(filename)
}

func (t *Table) FindPartialBlocks() ([]*TableBlock, error) {
	if _, err := t.LoadRecords(nil); err != nil {
		return nil, err
	}

	ret := make([]*TableBlock, 0)

	t.blockMu.Lock()
	for _, v := range t.BlockList {
		if v.Name == ROW_STORE_BLOCK {
			continue
		}

		if v.Info.NumRecords < int32(CHUNK_SIZE) {
			ret = append(ret, v)
		}
	}
	t.blockMu.Unlock()

	return ret, nil
}

// TODO: find any open blocks and then fill them...
func (t *Table) FillPartialBlock() error {
	if len(t.newRecords) == 0 {
		return nil
	}

	openBlocks, err := t.FindPartialBlocks()
	if err != nil {
		return err
	}

	Debug("OPEN BLOCKS", openBlocks)
	var filename string

	if len(openBlocks) == 0 {
		return nil
	}

	for _, b := range openBlocks {
		filename = b.Name
	}

	Debug("OPENING PARTIAL BLOCK", filename)

	if err := t.GrabBlockLock(filename); err != nil {
		Debug("CANT FILL PARTIAL BLOCK DUE TO LOCK", filename)
		return errors.Wrap(err, "failed to grab block lock")
	}

	defer t.ReleaseBlockLock(filename)

	// open up our last record block, see how full it is
	delete(t.BlockInfoCache, filename)

	block, err := t.LoadBlockFromDir(context.TODO(), filename, nil, true /* LOAD ALL RECORDS */, nil)
	if err != nil {
		return err
	}
	// TODO add error handling
	if block == nil {
		return nil
	}

	partialRecords := block.RecordList
	Debug("LAST BLOCK HAS", len(partialRecords), "RECORDS")

	if len(partialRecords) < CHUNK_SIZE {
		delta := CHUNK_SIZE - len(partialRecords)
		if delta > len(t.newRecords) {
			delta = len(t.newRecords)
		}

		Debug("SAVING PARTIAL RECORDS", delta, "TO", filename)
		partialRecords = append(partialRecords, t.newRecords[0:delta]...)
		if err := t.SaveRecordsToBlock(partialRecords, filename); err != nil {
			Debug("COULDNT SAVE PARTIAL RECORDS TO", filename)
			return errors.Wrap(err, "save records to block")
		}

		if delta < len(t.newRecords) {
			t.newRecords = t.newRecords[delta:]
		} else {
			t.newRecords = make(RecordList, 0)
		}
	}

	return nil
}

// optimizing for integer pre-cached info
func (t *Table) ShouldLoadBlockFromDir(dirname string, querySpec *QuerySpec) bool {
	if querySpec == nil {
		return true
	}

	info := t.LoadBlockInfo(dirname)

	maxRecord := Record{Ints: IntArr{}, Strs: StrArr{}}
	minRecord := Record{Ints: IntArr{}, Strs: StrArr{}}

	if len(info.IntInfoMap) == 0 {
		return true
	}

	for fieldName := range info.StrInfoMap {
		fieldID := t.getKeyID(fieldName)
		minRecord.ResizeFields(fieldID)
		maxRecord.ResizeFields(fieldID)
	}

	for fieldName, fieldInfo := range info.IntInfoMap {
		fieldID := t.getKeyID(fieldName)
		minRecord.ResizeFields(fieldID)
		maxRecord.ResizeFields(fieldID)

		minRecord.Ints[fieldID] = IntField(fieldInfo.Min)
		maxRecord.Ints[fieldID] = IntField(fieldInfo.Max)

		minRecord.Populated[fieldID] = INT_VAL
		maxRecord.Populated[fieldID] = INT_VAL
	}

	add := true
	for _, f := range querySpec.Filters {
		// make the minima record and the maxima records...
		switch fil := f.(type) {
		case IntFilter:
			if fil.Op == "gt" || fil.Op == "lt" {
				if !f.Filter(&minRecord) && !f.Filter(&maxRecord) {
					add = false
					break
				}
			}
		}
	}

	return add
}

func (t *Table) LoadBlockInfo(dirname string) *SavedColumnInfo {
	info := SavedColumnInfo{}
	if dirname == NULL_BLOCK {
		return &info
	}

	t.blockMu.Lock()
	cachedInfo, ok := t.BlockInfoCache[dirname]
	t.blockMu.Unlock()
	if ok {
		return cachedInfo
	}

	// find out how many records are kept in this dir...
	istart := time.Now()
	filename := fmt.Sprintf("%s/info.db", dirname)

	err := decodeInto(filename, &info)

	if err != nil {
		Warn("ERROR DECODING COLUMN BLOCK INFO!", dirname, err)
		return &info
	}
	iend := time.Now()

	if DEBUG_TIMING {
		Debug("LOAD BLOCK INFO TOOK", iend.Sub(istart))
	}

	t.blockMu.Lock()
	t.BlockInfoCache[dirname] = &info
	if info.NumRecords >= int32(CHUNK_SIZE) {
		t.NewBlockInfos = append(t.NewBlockInfos, dirname)
	}
	t.blockMu.Unlock()

	return &info
}

// TODO: have this only pull the blocks into column format and not materialize
// the columns immediately
func (t *Table) LoadBlockFromDir(ctx context.Context, dirname string, loadSpec *LoadSpec, loadRecords bool, replacements map[string]StrReplace) (*TableBlock, error) {
	ctx, done := t.trace(ctx)
	defer done()
	tb := newTableBlock()

	tb.Name = dirname

	tb.table = t

	info := t.LoadBlockInfo(dirname)

	if info == nil {
		return nil, nil
	}

	if info.NumRecords <= 0 {
		return nil, nil
	}

	t.blockMu.Lock()
	t.BlockList[dirname] = &tb
	t.blockMu.Unlock()

	tb.allocateRecords(loadSpec, *info, loadRecords)
	tb.Info = info

	file, _ := os.Open(dirname)
	files, _ := file.Readdir(-1)

	size := int64(0)

	for _, f := range files {
		_, span := trace.StartSpan(ctx, "load file")
		span.AddAttributes(trace.StringAttribute("file", f.Name()))
		span.AddAttributes(trace.Int64Attribute("size", f.Size()))
		fname := f.Name()
		fsize := f.Size()
		size += fsize

		// over here, we have to accomodate .gz extension, i guess
		if loadSpec != nil {
			// we cut off extensions to check our loadSpec
			cname := strings.TrimRight(fname, GZIP_EXT)

			if !loadSpec.files[cname] && !loadRecords {
				continue
			}
		} else if !loadRecords {
			continue
		}

		filename := fmt.Sprintf("%s/%s", dirname, fname)

		dec, err := GetFileDecoder(filename)
		if err != nil {
			return nil, err
		}

		switch {
		case strings.HasPrefix(fname, "str"):
			err = tb.unpackStrCol(dec, *info, replacements)
		case strings.HasPrefix(fname, "set"):
			err = tb.unpackSetCol(dec, *info)
		case strings.HasPrefix(fname, "int"):
			err = tb.unpackIntCol(dec, *info)
		}

		dec.CloseFile()

		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("issue unpacking %v", fname))
		}
		span.End()
	}

	tb.Size = size

	file.Close()
	return &tb, nil
}

type AfterLoadQueryCB struct {
	querySpec *QuerySpec
	wg        *sync.WaitGroup
	records   RecordList

	count int
}

func (cb *AfterLoadQueryCB) CB(digestname string, records RecordList) {
	if digestname == NO_MORE_BLOCKS {
		// TODO: add sessionization call over here, too
		count := FilterAndAggRecords(cb.querySpec, &cb.records)
		cb.count += count

		cb.wg.Done()
		return
	}

	querySpec := cb.querySpec

	for _, r := range records {
		add := true
		// FILTERING
		for j := 0; j < len(querySpec.Filters); j++ {
			// returns True if the record matches!
			ret := !querySpec.Filters[j].Filter(r)
			if ret {
				add = false
				break
			}
		}

		if !add {
			continue
		}

		cb.records = append(cb.records, r)
	}

	if FLAGS.DEBUG {
		fmt.Fprint(os.Stderr, "+")
	}
}

func (b *TableBlock) ExportBlockData() {
	if len(b.RecordList) == 0 {
		return
	}

	tsvData := make([]string, 0)

	for _, r := range b.RecordList {
		sample := r.toTSVRow()
		tsvData = append(tsvData, strings.Join(sample, "\t"))

	}

	exportName := path.Base(b.Name)
	dirName := path.Dir(b.Name)
	fName := path.Join(dirName, "export", exportName+".tsv.gz")

	os.MkdirAll(path.Join(dirName, "export"), 0755)

	tsvHeader := strings.Join(b.RecordList[0].sampleHeader(), "\t")
	tsvStr := strings.Join(tsvData, "\n")
	Debug("SAVING TSV ", len(tsvStr), "RECORDS", len(tsvData), fName)

	allData := strings.Join([]string{tsvHeader, tsvStr}, "\n")
	// Need to save these to a file.
	//	Print(tsv_headers)
	//	Print(tsv_str)

	// GZIPPING
	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)
	w.Write([]byte(allData))
	w.Close() // You must close this first to flush the bytes to the buffer.

	f, _ := os.Create(fName)
	_, err := f.Write(buf.Bytes())
	f.Close()

	if err != nil {
		Warn("COULDNT SAVE TSV FOR", fName, err)
	}

}

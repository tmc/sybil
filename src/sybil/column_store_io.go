package sybil

import "fmt"
import "bytes"

import "os"
import "encoding/gob"
import "runtime/debug"
import "time"
import "regexp"

import "github.com/DeDiS/protobuf"

type ValueMap map[int64][]uint32

var CARDINALITY_THRESHOLD = 4
var DEBUG_RECORD_CONSISTENCY = false
var PROTOBUF_ENABLED = true

func deltaEncodeCol(col ValueMap) {
	for _, records := range col {
		prev := uint32(0)
		for i, v := range records {
			records[int32(i)] = v - prev
			prev = v

		}
	}
}

func deltaEncode(sameMap map[int16]ValueMap) {
	for _, col := range sameMap {
		if len(col) <= CHUNK_SIZE/CARDINALITY_THRESHOLD {
			deltaEncodeCol(col)
		}
	}
}

// this is used to record the buckets when building the column
// blobs
func recordValue(sameMap map[int16]ValueMap, index int32, name int16, value int64) {
	s, ok := sameMap[name]
	if !ok {
		sameMap[name] = ValueMap{}
		s = sameMap[name]
	}

	vi := value

	s[vi] = append(s[vi], uint32(index))
}

func (tb *TableBlock) GetColumnInfo(nameId int16) *TableColumn {
	col, ok := tb.columns[nameId]
	if !ok {
		col = tb.newTableColumn()
		tb.columns[nameId] = col
	}

	return col
}

func (tb *TableBlock) SaveIntsToColumns(dirname string, sameInts map[int16]ValueMap) {
	// now make the dir and shoot each blob out into a separate file

	// SAVED TO A SINGLE BLOCK ON DISK, NOW TO SAVE IT OUT TO SEPARATE VALUES
	os.MkdirAll(dirname, 0777)
	for k, v := range sameInts {
		colName := tb.getStringForKey(k)
		if colName == "" {
			Debug("CANT FIGURE OUT FIELD NAME FOR", k, "SOMETHING IS PROBABLY AWRY")
			continue
		}
		intCol := NewSavedIntColumn()

		intCol.Name = colName
		intCol.DeltaEncodedIDs = OPTS.DELTA_ENCODE_RECORD_IDS

		maxR := 0
		recordToValue := make(map[uint32]int64)
		for bucket, records := range v {
			si := SavedIntBucket{Value: bucket, Records: records}
			intCol.Bins = append(intCol.Bins, si)
			for _, r := range records {
				recordToValue[r] = bucket
				if int(r) >= maxR {
					maxR = int(r) + 1
				}
			}

			// bookkeeping for info.db
			tb.updateIntInfo(k, bucket)
			tb.table.updateIntInfo(k, bucket)
		}

		intCol.BucketEncoded = true
		// the column is high cardinality?
		if len(intCol.Bins) > CHUNK_SIZE/CARDINALITY_THRESHOLD {
			intCol.BucketEncoded = false
			intCol.Bins = nil
			intCol.Values = make([]int64, maxR)
			intCol.ValueEncoded = OPTS.DELTA_ENCODE_INT_VALUES

			for r, val := range recordToValue {
				intCol.Values[r] = val
			}

			prev := int64(0)
			for r, val := range intCol.Values {
				if OPTS.DELTA_ENCODE_INT_VALUES {
					intCol.Values[r] = val - prev
					prev = val
				} else {
					intCol.Values[r] = val
				}
			}
		}

		var network bytes.Buffer
		colFname := fmt.Sprintf("%s/int_%s.db", dirname, tb.getStringForKey(k))
		if PROTOBUF_ENABLED {
			colFname = fmt.Sprintf("%s/int_%s.pb", dirname, tb.getStringForKey(k))

			buf, err := protobuf.Encode(&intCol)
			if err != nil {
				Error("encode:", err)
			}

			network = *bytes.NewBuffer(buf)
		} else {
			// Create an encoder and send a value.
			enc := gob.NewEncoder(&network)
			err := enc.Encode(intCol)
			if err != nil {
				Error("encode:", err)
			}
		}

		action := "SERIALIZED"
		if intCol.BucketEncoded {
			action = "BUCKETED  "
		}

		Debug(action, "COLUMN BLOCK", colFname, network.Len(), "BYTES", "( PER RECORD", network.Len()/len(tb.RecordList), ")")

		w, _ := os.Create(colFname)

		network.WriteTo(w)
	}

}

func (tb *TableBlock) SaveSetsToColumns(dirname string, sameSets map[int16]ValueMap) {
	for k, v := range sameSets {
		colName := tb.getStringForKey(k)
		if colName == "" {
			// TODO: validate what this means. I think it means reading 'null' values off disk
			// when pulling off incomplete records
			Debug("CANT FIGURE OUT FIELD NAME FOR", k, "PROBABLY AN ERRONEOUS FIELD")
			continue
		}
		setCol := SavedSetColumn{}
		setCol.Name = colName
		setCol.DeltaEncodedIDs = OPTS.DELTA_ENCODE_RECORD_IDS
		tempBlock := newTableBlock()

		tbCol := tb.GetColumnInfo(k)
		tempCol := tempBlock.GetColumnInfo(k)
		recordToValue := make(map[uint32][]int32)
		maxR := 0
		for bucket, records := range v {
			// migrating string definitions from column definitions
			strVal := tbCol.getStringForVal(int32(bucket))
			strId := tempCol.getValId(strVal)
			si := SavedSetBucket{Value: int32(strId), Records: records}
			setCol.Bins = append(setCol.Bins, si)
			for _, r := range records {
				_, ok := recordToValue[r]
				if int(r) >= maxR {
					maxR = int(r) + 1
				}

				if !ok {
					recordToValue[r] = make([]int32, 0)

				}

				recordToValue[r] = append(recordToValue[r], strId)
			}
		}

		setCol.StringTable = make([]string, len(tempCol.StringTable))
		for str, id := range tempCol.StringTable {
			setCol.StringTable[id] = str
		}

		// the column is high cardinality?
		setCol.BucketEncoded = true
		if len(setCol.Bins) > CHUNK_SIZE/CARDINALITY_THRESHOLD {
			setCol.BucketEncoded = false
			setCol.Bins = nil
			setCol.Values = make([][]int32, maxR)
			for k, v := range recordToValue {
				setCol.Values[k] = v
			}
		}

		colFname := fmt.Sprintf("%s/set_%s.db", dirname, tb.getStringForKey(k))

		var network bytes.Buffer // Stand-in for the network.

		// Create an encoder and send a value.
		enc := gob.NewEncoder(&network)
		err := enc.Encode(setCol)

		if err != nil {
			Error("encode:", err)
		}

		action := "SERIALIZED"
		if setCol.BucketEncoded {
			action = "BUCKETED  "
		}

		Debug(action, "COLUMN BLOCK", colFname, network.Len(), "BYTES", "( PER RECORD", network.Len()/len(tb.RecordList), ")")

		w, _ := os.Create(colFname)
		network.WriteTo(w)

	}
}

func (tb *TableBlock) SaveStrsToColumns(dirname string, sameStrs map[int16]ValueMap) {
	for k, v := range sameStrs {
		colName := tb.getStringForKey(k)
		if colName == "" {
			// TODO: validate what this means. I think it means reading 'null' values off disk
			// when pulling off incomplete records
			Debug("CANT FIGURE OUT FIELD NAME FOR", k, "PROBABLY AN ERRONEOUS FIELD")
			continue
		}
		strCol := NewSavedStrColumn()
		strCol.Name = colName
		strCol.DeltaEncodedIDs = OPTS.DELTA_ENCODE_RECORD_IDS
		tempBlock := newTableBlock()

		tempCol := tempBlock.GetColumnInfo(k)
		tbCol := tb.GetColumnInfo(k)
		recordToValue := make(map[uint32]int32)
		maxR := 0
		for bucket, records := range v {

			// migrating string definitions from column definitions
			strId := tempCol.getValId(tbCol.getStringForVal(int32(bucket)))

			si := SavedStrBucket{Value: strId, Records: records}
			strCol.Bins = append(strCol.Bins, si)
			for _, r := range records {
				recordToValue[r] = strId
				if r >= uint32(maxR) {
					maxR = int(r) + 1
				}
			}

			// also bookkeeping to be used later inside the block info.db, IMO
			tb.updateStrInfo(k, int(bucket), len(records))
		}

		strCol.BucketEncoded = true
		// the column is high cardinality?
		if len(strCol.Bins) > CHUNK_SIZE/CARDINALITY_THRESHOLD {
			strCol.BucketEncoded = false
			strCol.Bins = nil
			strCol.Values = make([]int32, maxR)
			for k, v := range recordToValue {
				strCol.Values[k] = v
			}
		}

		for _, bucket := range strCol.Bins {
			firstVal := bucket.Records[0]
			if firstVal > 1000 && DEBUG_RECORD_CONSISTENCY {
				Warn(k, bucket.Value, "FIRST RECORD IS", firstVal)
			}
		}

		tb.getStrInfo(k).prune()

		strCol.StringTable = make([]string, len(tempCol.StringTable))
		for str, id := range tempCol.StringTable {
			strCol.StringTable[id] = str
		}

		colFname := fmt.Sprintf("%s/str_%s.db", dirname, tb.getStringForKey(k))

		var network bytes.Buffer // Stand-in for the network.

		// Create an encoder and send a value.
		enc := gob.NewEncoder(&network)
		err := enc.Encode(strCol)

		if err != nil {
			Error("encode:", err)
		}

		action := "SERIALIZED"
		if strCol.BucketEncoded {
			action = "BUCKETED  "
		}

		Debug(action, "COLUMN BLOCK", colFname, network.Len(), "BYTES", "( PER RECORD", network.Len()/len(tb.RecordList), ")")

		w, _ := os.Create(colFname)
		network.WriteTo(w)

	}
}

type SavedIntInfo map[string]*IntInfo
type SavedStrInfo map[string]*StrInfo

func (tb *TableBlock) SaveInfoToColumns(dirname string) {
	records := tb.RecordList

	// Now to save block info...
	colFname := fmt.Sprintf("%s/info.db", dirname)

	var network bytes.Buffer // Stand-in for the network.

	// Create an encoder and send a value.
	enc := gob.NewEncoder(&network)

	savedIntInfo := SavedIntInfo{}
	savedStrInfo := SavedStrInfo{}
	if tb.Info != nil {
		if tb.Info.IntInfoMap != nil {
			savedIntInfo = tb.Info.IntInfoMap
		}
		if tb.Info.StrInfoMap != nil {
			savedStrInfo = tb.Info.StrInfoMap
		}
	}

	for k, v := range tb.IntInfo {
		name := tb.getStringForKey(k)
		savedIntInfo[name] = v
	}

	for k, v := range tb.StrInfo {
		name := tb.getStringForKey(k)
		savedStrInfo[name] = v
	}

	colInfo := SavedColumnInfo{NumRecords: int32(len(records)), IntInfoMap: savedIntInfo, StrInfoMap: savedStrInfo}
	err := enc.Encode(colInfo)

	if err != nil {
		Error("encode:", err)
	}

	length := len(records)
	if length == 0 {
		length = 1
	}

	if DEBUG_TIMING {
		Debug("SERIALIZED BLOCK INFO", colFname, network.Len(), "BYTES", "( PER RECORD", network.Len()/length, ")")
	}

	w, _ := os.Create(colFname)
	network.WriteTo(w)
}

type SeparatedColumns struct {
	ints map[int16]ValueMap
	strs map[int16]ValueMap
	sets map[int16]ValueMap
}

func (tb *TableBlock) SeparateRecordsIntoColumns() SeparatedColumns {
	records := tb.RecordList

	// making a cross section of records that share values
	// goes from fieldname{} -> value{} -> record
	sameInts := make(map[int16]ValueMap)
	sameStrs := make(map[int16]ValueMap)
	sameSets := make(map[int16]ValueMap)

	// parse record list and transfer book keeping data into the current
	// table block, as well as separate record values by column type
	for i, r := range records {
		for k, v := range r.Ints {
			if r.Populated[k] == INT_VAL {
				recordValue(sameInts, int32(i), int16(k), int64(v))
			}
		}
		for k, v := range r.Strs {
			// transition key from the
			col := r.block.GetColumnInfo(int16(k))
			newCol := tb.GetColumnInfo(int16(k))

			vName := col.getStringForVal(int32(v))
			vId := newCol.getValId(vName)

			// record the transitioned key
			if r.Populated[k] == STR_VAL {
				recordValue(sameStrs, int32(i), int16(k), int64(vId))
			}
		}
		for k, v := range r.SetMap {
			col := r.block.GetColumnInfo(int16(k))
			newCol := tb.GetColumnInfo(int16(k))
			if r.Populated[k] == SET_VAL {
				for _, iv := range v {
					vName := col.getStringForVal(int32(iv))
					vId := newCol.getValId(vName)
					recordValue(sameSets, int32(i), int16(k), int64(vId))
				}
			}
		}
	}

	if OPTS.DELTA_ENCODE_RECORD_IDS {
		deltaEncode(sameInts)
		deltaEncode(sameStrs)
		deltaEncode(sameSets)
	}

	ret := SeparatedColumns{ints: sameInts, strs: sameStrs, sets: sameSets}
	return ret

}

func (tb *TableBlock) SaveToColumns(filename string) bool {
	dirname := filename

	// Important to set the BLOCK's dirName so we can keep track
	// of the various block infos
	tb.Name = dirname

	defer tb.table.ReleaseBlockLock(filename)
	if tb.table.GrabBlockLock(filename) == false {
		Debug("Can't grab lock to save block", filename)
		return false
	}

	partialname := fmt.Sprintf("%s.partial", dirname)
	oldblock := fmt.Sprintf("%s.old", dirname)

	start := time.Now()
	oldPercent := debug.SetGCPercent(-1)
	separatedColumns := tb.SeparateRecordsIntoColumns()
	end := time.Now()
	Debug("COLLATING BLOCKS TOOK", end.Sub(start))

	tb.SaveIntsToColumns(partialname, separatedColumns.ints)
	tb.SaveStrsToColumns(partialname, separatedColumns.strs)
	tb.SaveSetsToColumns(partialname, separatedColumns.sets)
	tb.SaveInfoToColumns(partialname)

	end = time.Now()
	Debug("FINISHED BLOCK", partialname, "RELINKING TO", dirname, "TOOK", end.Sub(start))

	debug.SetGCPercent(oldPercent)

	// TODO: Add a stronger consistency check here
	// For now, we load info.db and check NumRecords inside it to prevent
	// catastrophics, but we could load everything potentially
	start = time.Now()
	nb := tb.table.LoadBlockFromDir(partialname, nil, false)
	end = time.Now()

	// TODO:
	if nb == nil || nb.Info.NumRecords != int32(len(tb.RecordList)) {
		Error("COULDNT VALIDATE CONSISTENCY FOR RECENTLY SAVED BLOCK!", filename)
	}

	if DEBUG_RECORD_CONSISTENCY {
		nb = tb.table.LoadBlockFromDir(partialname, nil, true)
		if nb == nil || len(nb.RecordList) != len(tb.RecordList) {
			Error("DEEP VALIDATION OF BLOCK FAILED CONSISTENCY CHECK!", filename)
		}
	}

	Debug("VALIDATED NEW BLOCK HAS", nb.Info.NumRecords, "RECORDS, TOOK", end.Sub(start))

	os.RemoveAll(oldblock)
	err := RenameAndMod(dirname, oldblock)
	if err != nil {
		Error("ERROR RENAMING BLOCK", dirname, oldblock, err)
	}
	err = RenameAndMod(partialname, dirname)
	if err != nil {
		Error("ERROR RENAMING PARTIAL", partialname, dirname, err)
	}

	if err == nil {
		os.RemoveAll(oldblock)
	} else {
		Error("ERROR SAVING BLOCK", partialname, dirname, err)
	}

	Debug("RELEASING BLOCK", tb.Name)
	return true

}

func (tb *TableBlock) unpackStrCol(dec FileDecoder, info SavedColumnInfo) {
	records := tb.RecordList[:]

	into := &SavedStrColumn{}
	err := dec.Decode(into)
	if err != nil {
		Debug("DECODE COL ERR:", err)
		return
	}

	stringLookup := make([]string, info.NumRecords)
	keyTableLen := len(tb.table.KeyTable)
	colId := tb.table.getKeyId(into.Name)

	if int(colId) >= keyTableLen {
		Debug("IGNORING STR COLUMN", into.Name, "SINCE ITS NOT IN KEY TABLE IN BLOCK", tb.Name)
		return
	}

	col := tb.GetColumnInfo(colId)
	// unpack the string table

	// Run our replacements!
	strReplace, ok := OPTS.STR_REPLACEMENTS[into.Name]
	bucketReplace := make(map[int32]int32)
	var re *regexp.Regexp
	if ok {
		re, err = regexp.Compile(strReplace.Pattern)
	}

	for k, v := range into.StringTable {
		var nv = v
		if re != nil {
			nv = re.ReplaceAllString(v, strReplace.Replace)
		}

		existingKey, exists := col.StringTable[nv]

		v = nv

		if exists {
			bucketReplace[int32(k)] = existingKey
		} else {
			bucketReplace[int32(k)] = int32(k)
			col.StringTable[v] = int32(k)
		}

		stringLookup[int32(k)] = v
	}

	col.valStringIdLookup = stringLookup

	isPathCol := false
	if FLAGS.PATH_KEY != nil {
		isPathCol = into.Name == *FLAGS.PATH_KEY
	}
	var record *Record
	var r uint32

	if into.BucketEncoded {
		prev := uint32(0)
		did := into.DeltaEncodedIDs

		for _, bucket := range into.Bins {
			prev = 0
			value := bucket.Value
			newValue, shouldReplace := bucketReplace[value]
			if shouldReplace {
				value = newValue
			}

			castValue := StrField(newValue)
			for _, r = range bucket.Records {

				if did {
					r = prev + r
				}

				prev = r
				record = records[r]

				if DEBUG_RECORD_CONSISTENCY {
					if record.Populated[colId] != _NO_VAL {
						Error("OVERWRITING RECORD VALUE", record, into.Name, colId, bucket.Value)
					}
				}

				records[r].Populated[colId] = STR_VAL
				records[r].Strs[colId] = castValue

				if isPathCol {
					record.Path = stringLookup[newValue]
				}
			}
		}

	} else {
		for r, v := range into.Values {
			newValue, shouldReplace := bucketReplace[v]
			if shouldReplace {
				v = newValue
			}

			records[r].Strs[colId] = StrField(v)
			records[r].Populated[colId] = STR_VAL
		}

	}
}

func (tb *TableBlock) unpackSetCol(dec FileDecoder, info SavedColumnInfo) {
	records := tb.RecordList

	savedCol := NewSavedSetColumn()
	into := &savedCol
	err := dec.Decode(into)
	if err != nil {
		Debug("DECODE COL ERR:", err)
	}

	keyTableLen := len(tb.table.KeyTable)
	colId := tb.table.getKeyId(into.Name)
	stringLookup := make(map[int32]string)

	if int(colId) >= keyTableLen {
		Debug("IGNORING SET COLUMN", into.Name, "SINCE ITS NOT IN KEY TABLE IN BLOCK", tb.Name)
		return
	}

	col := tb.GetColumnInfo(colId)
	// unpack the string table
	for k, v := range into.StringTable {
		col.StringTable[v] = int32(k)
		stringLookup[int32(k)] = v
	}

	trStringLookup := make([]string, len(stringLookup))
	for k, v := range stringLookup {
		trStringLookup[k] = v
	}

	col.valStringIdLookup = trStringLookup

	if into.BucketEncoded {
		for _, bucket := range into.Bins {
			// DONT FORGET TO DELTA UNENCODE THE RECORD VALUES
			prev := uint32(0)
			for _, r := range bucket.Records {
				if into.DeltaEncodedIDs {
					r = r + prev
				}

				curSet, ok := records[r].SetMap[colId]
				if !ok {
					curSet = make(SetField, 0)
				}

				curSet = append(curSet, bucket.Value)
				records[r].SetMap[colId] = curSet

				records[r].Populated[colId] = SET_VAL
				prev = r
			}

		}
	} else {
		for r, v := range into.Values {
			curSet, ok := records[r].SetMap[colId]
			if !ok {
				curSet = make(SetField, 0)
				records[r].SetMap[colId] = curSet
			}

			records[r].SetMap[colId] = SetField(v)
			records[r].Populated[colId] = SET_VAL
		}
	}
}

func (tb *TableBlock) unpackIntCol(dec FileDecoder, info SavedColumnInfo) {
	records := tb.RecordList[:]

	into := &SavedIntColumn{}
	err := dec.Decode(into)
	if err != nil {
		Debug("DECODE COL ERR:", err)
	}

	keyTableLen := len(tb.table.KeyTable)
	colId := tb.table.getKeyId(into.Name)
	if int(colId) >= keyTableLen {
		Debug("IGNORING INT COLUMN", into.Name, "SINCE ITS NOT IN KEY TABLE IN BLOCK", tb.Name)
		return
	}

	isTimeCol := false
	if FLAGS.TIME_COL != nil {
		isTimeCol = into.Name == *FLAGS.TIME_COL
	}

	if into.BucketEncoded {
		for _, bucket := range into.Bins {
			if *FLAGS.UPDATE_TABLE_INFO {
				tb.updateIntInfo(colId, bucket.Value)
				tb.table.updateIntInfo(colId, bucket.Value)
			}

			// DONT FORGET TO DELTA UNENCODE THE RECORD VALUES
			prev := uint32(0)
			for _, r := range bucket.Records {
				if into.DeltaEncodedIDs {
					r = r + prev
				}

				if DEBUG_RECORD_CONSISTENCY {
					if records[r].Populated[colId] != _NO_VAL {
						Error("OVERWRITING RECORD VALUE", records[r], into.Name, colId, bucket.Value)
					}
				}

				records[r].Ints[colId] = IntField(bucket.Value)
				records[r].Populated[colId] = INT_VAL
				prev = r

				if isTimeCol {
					records[r].Timestamp = bucket.Value
				}

			}

		}
	} else {

		prev := int64(0)
		for r, v := range into.Values {
			if *FLAGS.UPDATE_TABLE_INFO {
				tb.updateIntInfo(colId, v)
				tb.table.updateIntInfo(colId, v)
			}

			if into.ValueEncoded {
				v = v + prev
			}

			records[r].Ints[colId] = IntField(v)
			records[r].Populated[colId] = INT_VAL

			if isTimeCol {
				records[r].Timestamp = v
			}

			if into.ValueEncoded {
				prev = v
			}

		}
	}
}

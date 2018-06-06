package sybil

import (
	"crypto/md5"
	"encoding/json"
	"fmt"

	hll "github.com/logv/loglogbeta"
)

type ResultMap map[string]*Result

// This info gets cached when we use
// the query cache. anything in the main
// QuerySpec will not get cached
type savedQueryResults struct {
	Cumulative   *Result
	Results      ResultMap
	TimeResults  map[int]ResultMap
	MatchedCount int
	Sorted       []*Result
	Matched      RecordList
}

type savedQueryParams struct {
	Filters      []Filter              `json:",omitempty"`
	Groups       []Grouping            `json:",omitempty"`
	Aggregations []Aggregation         `json:",omitempty"`
	Distincts    []Grouping            `json:",omitempty"` // list of columns we are creating a count distinct query on
	StrReplace   map[string]StrReplace `json:",omitempty"`

	OrderBy    string `json:",omitempty"`
	PruneBy    string `json:",omitempty"`
	Limit      int    `json:",omitempty"`
	TimeBucket int    `json:",omitempty"`

	Samples       bool `json:",omitempty"`
	CachedQueries bool `json:",omitempty"`

	HistogramParameters HistogramParameters `json:",omitempty"`
}

func Min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

// For outside consumption
type QueryParams savedQueryParams
type QueryResults savedQueryResults

type QuerySpec struct {
	QueryParams
	QueryResults

	BlockList map[string]TableBlock
	Table     *Table
}

type Filter interface {
	Filter(*Record) bool
}

type Grouping struct {
	Name   string
	nameID int16
}

type Aggregation struct {
	Op       string
	opID     int
	Name     string
	nameID   int16
	HistType string
}

type Result struct {
	Hists    map[string]Histogram
	Distinct *hll.LogLogBeta

	GroupByKey  string
	BinaryByKey string
	Count       int64
	Samples     int64
}

func (qs *QuerySpec) NewResult() *Result {
	addedRecord := &Result{}
	addedRecord.Hists = make(map[string]Histogram)

	if len(qs.Distincts) > 0 {
		addedRecord.Distinct = hll.New()
	}

	addedRecord.Count = 0
	return addedRecord
}

func (master_result *ResultMap) Combine(params HistogramParameters, mergeTable *Table, results *ResultMap) {
	for k, v := range *results {
		mval, ok := (*master_result)[k]
		if !ok {
			(*master_result)[k] = v
		} else {
			mval.Combine(params, mergeTable, v)
		}
	}
}

func fullMergeHist(params HistogramParameters, mergeTable *Table, h, ph Histogram) Histogram {
	l1, r1 := h.Range()
	l2, r2 := ph.Range()

	info := IntInfo{Min: Min(l1, l2), Max: Max(r1, r2)}

	nh := mergeTable.NewHist(params, &info, h.IsWeighted())

	for bucket, count := range h.GetIntBuckets() {
		nh.AddWeightedValue(bucket, count)
	}

	for bucket, count := range ph.GetIntBuckets() {
		nh.AddWeightedValue(bucket, count)
	}

	return nh
}

// This does an in place combine of the next_result into this one...
func (rs *Result) Combine(params HistogramParameters, mergeTable *Table, nextResult *Result) {
	if nextResult == nil {
		return
	}

	if nextResult.Count == 0 {
		return
	}

	totalSamples := rs.Samples + nextResult.Samples
	totalCount := rs.Count + nextResult.Count

	// combine histograms...
	for k, h := range nextResult.Hists {

		// If we are doing a node aggregation, we have a MERGE_TABLE
		// set, which means we should go the slow route when merging
		// histograms because we can't be sure they were created with
		// the same extents (being that they may originate from different
		// nodes)
		if OPTS.MERGE_TABLE != nil {
			ph, ok := rs.Hists[k]

			if ok {
				rs.Hists[k] = fullMergeHist(params, mergeTable, h, ph)
			} else {
				rs.Hists[k] = h
			}

		} else {
			_, ok := rs.Hists[k]
			if !ok {
				nh := h.NewHist(params)

				nh.Combine(h)
				rs.Hists[k] = nh
			} else {
				rs.Hists[k].Combine(h)
			}
		}
	}

	// combine count distincts
	if nextResult.Distinct != nil {

		if rs.Distinct == nil {
			rs.Distinct = nextResult.Distinct
		} else {
			rs.Distinct.Merge(nextResult.Distinct)

		}
	}

	rs.Samples = totalSamples
	rs.Count = totalCount
}

// Punctuate resets results on the QuerySpec.
func (qs *QuerySpec) Punctuate() {
	qs.Results = make(ResultMap)
	qs.TimeResults = make(map[int]ResultMap)
}

// ResetResults resets results on the QuerySpec and clears table caches.
func (qs *QuerySpec) ResetResults() {
	qs.Punctuate()

	if qs.Table != nil && qs.Table.BlockList != nil {
		// Reach into all our table blocks and reset their REGEX CACHE
		for _, b := range qs.Table.BlockList {
			for _, c := range b.columns {
				if len(c.RCache) > 0 {
					c.RCache = make(map[int]bool)
				}
			}
		}
	}
}
func (t *Table) Grouping(name string) Grouping {
	colID := t.getKeyID(name)
	return Grouping{name, colID}
}

func (t *Table) Aggregation(histogramType HistogramType, name string, op string) Aggregation {
	colID := t.getKeyID(name)

	agg := Aggregation{Name: name, nameID: colID, Op: op}
	if op == "avg" {
		agg.opID = OP_AVG
	}

	if op == "hist" {
		agg.opID = OP_HIST
		agg.HistType = "basic"
		if histogramType == HistogramTypeLog {
			agg.HistType = "multi"
		}

		if histogramType == HistogramTypeHDR {
			agg.HistType = "hdr"
		}
	}

	if op == DISTINCT_STR {
		agg.opID = OP_DISTINCT
	}

	_, ok := t.IntInfo[colID]
	if !ok {
		// TODO: tell our table we need to load all records!
		Debug("MISSING CACHED INFO FOR", agg)
	}
	return agg
}

// cacheKey returns a stable identifier.
func (qp QueryParams) cacheKey() string {
	buf, err := json.Marshal(qp)
	if err != nil {
		panic(err)
	}

	h := md5.New()
	h.Write(buf)

	ret := fmt.Sprintf("%x", h.Sum(nil))
	return ret

}

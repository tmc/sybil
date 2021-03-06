package sybil

import "testing"
import "math/rand"
import "math"
import "strconv"

func TestCachedQueries(t *testing.T) {
	tableName := getTestTableName(t)
	deleteTestDb(tableName)

	blockCount := 5

	var thisAddRecords = func(block_count int) {
		addRecords(tableName, func(r *Record, i int) {
			age := int64(rand.Intn(20)) + 10

			ageStr := strconv.FormatInt(int64(age), 10)
			r.AddIntField("id", int64(i))
			r.AddIntField("age", age)
			r.AddStrField("age_str", ageStr)
			r.AddSetField("age_set", []string{ageStr})

		}, block_count)
		saveAndReloadTable(t, tableName, block_count)

	}

	thisAddRecords(blockCount)
	testCachedQueryFiles(t, tableName)
	deleteTestDb(tableName)

	thisAddRecords(blockCount)
	testCachedQueryConsistency(t, tableName)
	deleteTestDb(tableName)

	thisAddRecords(blockCount)
	testCachedBasicHist(t, tableName)
	deleteTestDb(tableName)
}

func testCachedQueryFiles(t *testing.T, tableName string) {
	nt := GetTable(tableName)
	filters := []Filter{}
	filters = append(filters, nt.IntFilter("age", "lt", 20))

	aggs := []Aggregation{}
	aggs = append(aggs, nt.Aggregation("age", "hist"))

	querySpec := QuerySpec{Table: nt,
		QueryParams: QueryParams{
			Filters:       filters,
			Aggregations:  aggs,
			CachedQueries: true,
		},
	}
	loadSpec := NewLoadSpec()
	loadSpec.LoadAllColumns = true
	loadSpec.SkipDeleteBlocksAfterQuery = true

	// test that the cached query doesnt already exist
	nt.LoadAndQueryRecords(&loadSpec, nil)
	for _, b := range nt.BlockList {
		loaded := querySpec.LoadCachedResults(b.Name)
		if loaded {
			t.Error("Test DB started with saved query results")
		}
	}

	// test that the cached query is saved
	nt.LoadAndQueryRecords(&loadSpec, &querySpec)
	for _, b := range nt.BlockList {
		loaded := querySpec.LoadCachedResults(b.Name)
		if !loaded {
			t.Error("Did not correctly save and load query results")
		}
	}

	querySpec.CachedQueries = false
	for _, b := range nt.BlockList {
		loaded := querySpec.LoadCachedResults(b.Name)
		if loaded {
			t.Error("Used query cache when flag was not provided")
		}
	}
	querySpec.CachedQueries = true

	// test that a new and slightly different query isnt cached for us
	nt.LoadAndQueryRecords(&loadSpec, nil)
	querySpec.Aggregations = append(aggs, nt.Aggregation("id", "hist"))
	for _, b := range nt.BlockList {
		loaded := querySpec.LoadCachedResults(b.Name)
		if loaded {
			t.Error("Test DB has query results for new query")
		}
	}

}

func testCachedQueryConsistency(t *testing.T, tableName string) {
	nt := GetTable(tableName)
	filters := []Filter{}
	filters = append(filters, nt.IntFilter("age", "lt", 20))

	aggs := []Aggregation{}
	aggs = append(aggs, nt.Aggregation("age", "hist"))

	querySpec := QuerySpec{Table: nt,
		QueryParams: QueryParams{
			Filters:       filters,
			Aggregations:  aggs,
			CachedQueries: true,
		},
	}
	loadSpec := NewLoadSpec()
	loadSpec.LoadAllColumns = true
	loadSpec.SkipDeleteBlocksAfterQuery = true

	nt.LoadAndQueryRecords(&loadSpec, &querySpec)
	copySpec := CopyQuerySpec(&querySpec)

	nt = GetTable(tableName)

	// clear the copied query spec result map and look
	// at the cached query results

	copySpec.Results = make(ResultMap)
	nt.LoadAndQueryRecords(&loadSpec, copySpec)

	if len(querySpec.Results) == 0 {
		t.Error("No Results for Query")
	}

	for k, v := range querySpec.Results {
		v2, ok := copySpec.Results[k]
		if !ok {
			t.Error("Result Mismatch!", k, v)
		}

		if v.Count != v2.Count {
			t.Error("Count Mismatch", v, v2, v.Count, v2.Count)
		}

		if v.Samples != v2.Samples {
			Debug(v, v2)
			t.Error("Samples Mismatch", v, v2, v.Samples, v2.Samples)
		}

	}

	for _, b := range nt.BlockList {
		loaded := querySpec.LoadCachedResults(b.Name)
		if !loaded {
			t.Error("Did not correctly save and load query results")
		}
	}

}

func testCachedBasicHist(t *testing.T, tableName string) {
	nt := GetTable(tableName)

	for _, histType := range []HistogramType{HistogramTypeBasic, HistogramTypeLog} {
		filters := []Filter{}
		filters = append(filters, nt.IntFilter("age", "lt", 20))
		aggs := []Aggregation{}
		aggs = append(aggs, nt.Aggregation("age", "hist"))

		querySpec := QuerySpec{Table: nt,
			QueryParams: QueryParams{
				Filters:       filters,
				Aggregations:  aggs,
				CachedQueries: true,
				HistogramParameters: HistogramParameters{
					Type: histType,
				},
			},
		}

		loadSpec := NewLoadSpec()
		loadSpec.LoadAllColumns = true
		loadSpec.SkipDeleteBlocksAfterQuery = true

		nt.LoadAndQueryRecords(&loadSpec, &querySpec)
		copySpec := CopyQuerySpec(&querySpec)

		nt = GetTable(tableName)

		// clear the copied query spec result map and look
		// at the cached query results

		copySpec.Results = make(ResultMap)
		nt.LoadAndQueryRecords(&loadSpec, copySpec)

		if len(querySpec.Results) == 0 {
			t.Error("No Results for Query")
		}

		for k, v := range querySpec.Results {
			v2, ok := copySpec.Results[k]
			if !ok {
				t.Error("Result Mismatch!", histType, k, v)
			}

			if v.Count != v2.Count {
				t.Error("Count Mismatch", histType, v, v2, v.Count, v2.Count)
			}

			if v.Samples != v2.Samples {
				Debug(v, v2)
				t.Error("Samples Mismatch", histType, v, v2, v.Samples, v2.Samples)
			}

			for k, h := range v.Hists {
				h2, ok := v2.Hists[k]
				if !ok {
					t.Error("Missing Histogram", histType, v, v2)
				}

				if h.StdDev() <= 0 {
					t.Error("Missing StdDev", histType, h, h.StdDev())
				}

				if math.Abs(h.StdDev()-h2.StdDev()) > 0.1 {
					t.Error("StdDev MisMatch", histType, h, h2)
				}

			}

		}

		for _, b := range nt.BlockList {
			loaded := querySpec.LoadCachedResults(b.Name)
			if !loaded {
				t.Error("Did not correctly save and load query results")
			}
		}
	}

}

func TestCacheKeyGeneration(t *testing.T) {
	tests := []struct {
		name string
		qp   QueryParams
		want string
	}{
		{
			"empty",
			QueryParams{},
			"f34409debd75d424ca268a26bc517440",
		},
		{
			"with-replacements",
			QueryParams{StrReplace: map[string]StrReplace{
				"a": StrReplace{},
				"b": StrReplace{},
			}},
			"3e260f59f6db10b9e96ca8c05319985d",
		},
	}

	// TODO: once we're on go1.7 use t.Run
	for _, tt := range tests {
		if got := tt.qp.cacheKey(); got != tt.want {
			t.Errorf("%q. cacheKey = \n%v, want \n%v", tt.name, got, tt.want)

		}
	}
}

package sybil

import (
	"math/rand"
	"strconv"
	"testing"
)

func BenchmarkGetTableUnloadTable(b *testing.B) {
	tableName := getTestTableName(nil)
	deleteTestDb(tableName)
	defer deleteTestDb(tableName)
	defer b.StopTimer()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetTable(tableName)
		UnloadTable(tableName)
	}
}

func BenchmarkAddRecordsA(b *testing.B) {
	tableName := getTestTableName(nil)
	deleteTestDb(tableName)
	defer deleteTestDb(tableName)
	defer b.StopTimer()

	tbl := GetTable(tableName)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r := tbl.NewRecord()
		r.AddIntField("id", int64(i))
	}
}

func BenchmarkAddRecordsB(b *testing.B) {
	tableName := getTestTableName(nil)
	deleteTestDb(tableName)
	defer deleteTestDb(tableName)
	defer b.StopTimer()

	tbl := GetTable(tableName)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r := tbl.NewRecord()
		r.AddIntField("id", int64(i))
		age := int64(rand.Intn(20)) + 10
		r.AddIntField("age", age)
		r.AddStrField("age_str", strconv.FormatInt(int64(age), 10))
	}
}

func BenchmarkSaveRecordsToColumnsA(b *testing.B) {
	tableName := getTestTableName(nil)
	deleteTestDb(tableName)
	defer deleteTestDb(tableName)
	defer b.StopTimer()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		tbl := GetTable(tableName)
		r := tbl.NewRecord()
		r.AddIntField("id", int64(i))
		b.StartTimer()
		tbl.SaveRecordsToColumns()
		b.StopTimer()
		unloadTestTable(tableName)
	}
}

func BenchmarkSaveRecordsToColumnsB(b *testing.B) {
	tableName := getTestTableName(nil)
	deleteTestDb(tableName)
	defer deleteTestDb(tableName)
	defer b.StopTimer()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		tbl := GetTable(tableName)
		r := tbl.NewRecord()
		r.AddIntField("id", int64(i))
		age := int64(rand.Intn(20)) + 10
		r.AddIntField("age", age)
		r.AddStrField("age_str", strconv.FormatInt(int64(age), 10))

		b.StartTimer()
		tbl.SaveRecordsToColumns()
		b.StopTimer()

		unloadTestTable(tableName)
	}
}

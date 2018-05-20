// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sybild.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	sybild.proto

It has these top-level messages:
	IngestRequest
	IngestResponse
	QueryFilter
	QueryRequest
	Histogram
	QueryResult
	ResultMap
	SetField
	Record
	QueryResults
	QueryResponse
	ListTablesRequest
	ListTablesResponse
	GetTableRequest
	Table
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/struct"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// QueryType
type QueryType int32

const (
	QueryType_QUERY_TYPE_UNKNOWN QueryType = 0
	QueryType_TABLE              QueryType = 1
	QueryType_TIME_SERIES        QueryType = 2
	QueryType_DISTRIBUTION       QueryType = 3
	QueryType_SAMPLES            QueryType = 4
)

var QueryType_name = map[int32]string{
	0: "QUERY_TYPE_UNKNOWN",
	1: "TABLE",
	2: "TIME_SERIES",
	3: "DISTRIBUTION",
	4: "SAMPLES",
}
var QueryType_value = map[string]int32{
	"QUERY_TYPE_UNKNOWN": 0,
	"TABLE":              1,
	"TIME_SERIES":        2,
	"DISTRIBUTION":       3,
	"SAMPLES":            4,
}

func (x QueryType) String() string {
	return proto.EnumName(QueryType_name, int32(x))
}
func (QueryType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// QueryOp
type QueryOp int32

const (
	QueryOp_QUERY_OP_UNKNOWN QueryOp = 0
	QueryOp_AVERAGE          QueryOp = 1
	QueryOp_HISTOGRAM        QueryOp = 2
)

var QueryOp_name = map[int32]string{
	0: "QUERY_OP_UNKNOWN",
	1: "AVERAGE",
	2: "HISTOGRAM",
}
var QueryOp_value = map[string]int32{
	"QUERY_OP_UNKNOWN": 0,
	"AVERAGE":          1,
	"HISTOGRAM":        2,
}

func (x QueryOp) String() string {
	return proto.EnumName(QueryOp_name, int32(x))
}
func (QueryOp) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// QueryFilterOp
type QueryFilterOp int32

const (
	QueryFilterOp_QUERY_FILTER_OP_UNKNOWN QueryFilterOp = 0
	QueryFilterOp_RE                      QueryFilterOp = 1
	QueryFilterOp_NRE                     QueryFilterOp = 2
	QueryFilterOp_EQ                      QueryFilterOp = 3
	QueryFilterOp_NEQ                     QueryFilterOp = 4
	QueryFilterOp_GT                      QueryFilterOp = 5
	QueryFilterOp_LT                      QueryFilterOp = 6
	QueryFilterOp_IN                      QueryFilterOp = 7
	QueryFilterOp_NIN                     QueryFilterOp = 8
)

var QueryFilterOp_name = map[int32]string{
	0: "QUERY_FILTER_OP_UNKNOWN",
	1: "RE",
	2: "NRE",
	3: "EQ",
	4: "NEQ",
	5: "GT",
	6: "LT",
	7: "IN",
	8: "NIN",
}
var QueryFilterOp_value = map[string]int32{
	"QUERY_FILTER_OP_UNKNOWN": 0,
	"RE":  1,
	"NRE": 2,
	"EQ":  3,
	"NEQ": 4,
	"GT":  5,
	"LT":  6,
	"IN":  7,
	"NIN": 8,
}

func (x QueryFilterOp) String() string {
	return proto.EnumName(QueryFilterOp_name, int32(x))
}
func (QueryFilterOp) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// IngestRequest
type IngestRequest struct {
	// dataset is the name of the dataset.
	Dataset string `protobuf:"bytes,1,opt,name=dataset" json:"dataset,omitempty"`
	// records is the set of records to insert.
	Records []*google_protobuf.Struct `protobuf:"bytes,2,rep,name=records" json:"records,omitempty"`
}

func (m *IngestRequest) Reset()                    { *m = IngestRequest{} }
func (m *IngestRequest) String() string            { return proto.CompactTextString(m) }
func (*IngestRequest) ProtoMessage()               {}
func (*IngestRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IngestRequest) GetDataset() string {
	if m != nil {
		return m.Dataset
	}
	return ""
}

func (m *IngestRequest) GetRecords() []*google_protobuf.Struct {
	if m != nil {
		return m.Records
	}
	return nil
}

// IngestResponse
type IngestResponse struct {
	NumberInserted int64 `protobuf:"varint,1,opt,name=number_inserted,json=numberInserted" json:"number_inserted,omitempty"`
}

func (m *IngestResponse) Reset()                    { *m = IngestResponse{} }
func (m *IngestResponse) String() string            { return proto.CompactTextString(m) }
func (*IngestResponse) ProtoMessage()               {}
func (*IngestResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *IngestResponse) GetNumberInserted() int64 {
	if m != nil {
		return m.NumberInserted
	}
	return 0
}

// QueryFilter
type QueryFilter struct {
	Column string        `protobuf:"bytes,1,opt,name=column" json:"column,omitempty"`
	Op     QueryFilterOp `protobuf:"varint,2,opt,name=op,enum=pb.QueryFilterOp" json:"op,omitempty"`
	Value  string        `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
}

func (m *QueryFilter) Reset()                    { *m = QueryFilter{} }
func (m *QueryFilter) String() string            { return proto.CompactTextString(m) }
func (*QueryFilter) ProtoMessage()               {}
func (*QueryFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *QueryFilter) GetColumn() string {
	if m != nil {
		return m.Column
	}
	return ""
}

func (m *QueryFilter) GetOp() QueryFilterOp {
	if m != nil {
		return m.Op
	}
	return QueryFilterOp_QUERY_FILTER_OP_UNKNOWN
}

func (m *QueryFilter) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// QueryRequest
type QueryRequest struct {
	// dataset is the name of the dataset.
	Dataset                string         `protobuf:"bytes,1,opt,name=dataset" json:"dataset,omitempty"`
	Type                   QueryType      `protobuf:"varint,2,opt,name=type,enum=pb.QueryType" json:"type,omitempty"`
	Limit                  int64          `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
	Ints                   []string       `protobuf:"bytes,4,rep,name=ints" json:"ints,omitempty"`
	Strs                   []string       `protobuf:"bytes,5,rep,name=strs" json:"strs,omitempty"`
	Groups                 []string       `protobuf:"bytes,6,rep,name=groups" json:"groups,omitempty"`
	DistinctBy             []string       `protobuf:"bytes,7,rep,name=distinct_by,json=distinctBy" json:"distinct_by,omitempty"`
	Sort                   string         `protobuf:"bytes,8,opt,name=sort" json:"sort,omitempty"`
	TimeColumn             string         `protobuf:"bytes,9,opt,name=time_column,json=timeColumn" json:"time_column,omitempty"`
	TimeBucket             int64          `protobuf:"varint,10,opt,name=time_bucket,json=timeBucket" json:"time_bucket,omitempty"`
	Op                     QueryOp        `protobuf:"varint,11,opt,name=op,enum=pb.QueryOp" json:"op,omitempty"`
	LogHistogram           bool           `protobuf:"varint,12,opt,name=log_histogram,json=logHistogram" json:"log_histogram,omitempty"`
	IntFilters             []*QueryFilter `protobuf:"bytes,13,rep,name=int_filters,json=intFilters" json:"int_filters,omitempty"`
	StrFilters             []*QueryFilter `protobuf:"bytes,14,rep,name=str_filters,json=strFilters" json:"str_filters,omitempty"`
	SetFilters             []*QueryFilter `protobuf:"bytes,15,rep,name=set_filters,json=setFilters" json:"set_filters,omitempty"`
	IntHistogramBucketSize int64          `protobuf:"varint,16,opt,name=int_histogram_bucket_size,json=intHistogramBucketSize" json:"int_histogram_bucket_size,omitempty"`
}

func (m *QueryRequest) Reset()                    { *m = QueryRequest{} }
func (m *QueryRequest) String() string            { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()               {}
func (*QueryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *QueryRequest) GetDataset() string {
	if m != nil {
		return m.Dataset
	}
	return ""
}

func (m *QueryRequest) GetType() QueryType {
	if m != nil {
		return m.Type
	}
	return QueryType_QUERY_TYPE_UNKNOWN
}

func (m *QueryRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *QueryRequest) GetInts() []string {
	if m != nil {
		return m.Ints
	}
	return nil
}

func (m *QueryRequest) GetStrs() []string {
	if m != nil {
		return m.Strs
	}
	return nil
}

func (m *QueryRequest) GetGroups() []string {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *QueryRequest) GetDistinctBy() []string {
	if m != nil {
		return m.DistinctBy
	}
	return nil
}

func (m *QueryRequest) GetSort() string {
	if m != nil {
		return m.Sort
	}
	return ""
}

func (m *QueryRequest) GetTimeColumn() string {
	if m != nil {
		return m.TimeColumn
	}
	return ""
}

func (m *QueryRequest) GetTimeBucket() int64 {
	if m != nil {
		return m.TimeBucket
	}
	return 0
}

func (m *QueryRequest) GetOp() QueryOp {
	if m != nil {
		return m.Op
	}
	return QueryOp_QUERY_OP_UNKNOWN
}

func (m *QueryRequest) GetLogHistogram() bool {
	if m != nil {
		return m.LogHistogram
	}
	return false
}

func (m *QueryRequest) GetIntFilters() []*QueryFilter {
	if m != nil {
		return m.IntFilters
	}
	return nil
}

func (m *QueryRequest) GetStrFilters() []*QueryFilter {
	if m != nil {
		return m.StrFilters
	}
	return nil
}

func (m *QueryRequest) GetSetFilters() []*QueryFilter {
	if m != nil {
		return m.SetFilters
	}
	return nil
}

func (m *QueryRequest) GetIntHistogramBucketSize() int64 {
	if m != nil {
		return m.IntHistogramBucketSize
	}
	return 0
}

type Histogram struct {
	Mean         float64          `protobuf:"fixed64,1,opt,name=mean" json:"mean,omitempty"`
	Max          int64            `protobuf:"varint,2,opt,name=max" json:"max,omitempty"`
	Min          int64            `protobuf:"varint,3,opt,name=min" json:"min,omitempty"`
	TotalCount   int64            `protobuf:"varint,4,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
	Percentiles  []int64          `protobuf:"varint,5,rep,packed,name=percentiles" json:"percentiles,omitempty"`
	StrBuckets   map[string]int64 `protobuf:"bytes,6,rep,name=str_buckets,json=strBuckets" json:"str_buckets,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	IntBuckets   map[int64]int64  `protobuf:"bytes,7,rep,name=int_buckets,json=intBuckets" json:"int_buckets,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	RangeStart   int64            `protobuf:"varint,8,opt,name=range_start,json=rangeStart" json:"range_start,omitempty"`
	RangeEnd     int64            `protobuf:"varint,9,opt,name=range_end,json=rangeEnd" json:"range_end,omitempty"`
	StdDeviation float64          `protobuf:"fixed64,10,opt,name=std_deviation,json=stdDeviation" json:"std_deviation,omitempty"`
}

func (m *Histogram) Reset()                    { *m = Histogram{} }
func (m *Histogram) String() string            { return proto.CompactTextString(m) }
func (*Histogram) ProtoMessage()               {}
func (*Histogram) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Histogram) GetMean() float64 {
	if m != nil {
		return m.Mean
	}
	return 0
}

func (m *Histogram) GetMax() int64 {
	if m != nil {
		return m.Max
	}
	return 0
}

func (m *Histogram) GetMin() int64 {
	if m != nil {
		return m.Min
	}
	return 0
}

func (m *Histogram) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *Histogram) GetPercentiles() []int64 {
	if m != nil {
		return m.Percentiles
	}
	return nil
}

func (m *Histogram) GetStrBuckets() map[string]int64 {
	if m != nil {
		return m.StrBuckets
	}
	return nil
}

func (m *Histogram) GetIntBuckets() map[int64]int64 {
	if m != nil {
		return m.IntBuckets
	}
	return nil
}

func (m *Histogram) GetRangeStart() int64 {
	if m != nil {
		return m.RangeStart
	}
	return 0
}

func (m *Histogram) GetRangeEnd() int64 {
	if m != nil {
		return m.RangeEnd
	}
	return 0
}

func (m *Histogram) GetStdDeviation() float64 {
	if m != nil {
		return m.StdDeviation
	}
	return 0
}

type QueryResult struct {
	Histograms  map[string]*Histogram `protobuf:"bytes,1,rep,name=histograms" json:"histograms,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	GroupByKey  string                `protobuf:"bytes,3,opt,name=GroupByKey" json:"GroupByKey,omitempty"`
	BinaryByKey string                `protobuf:"bytes,4,opt,name=BinaryByKey" json:"BinaryByKey,omitempty"`
	Count       int64                 `protobuf:"varint,5,opt,name=Count" json:"Count,omitempty"`
	Samples     int64                 `protobuf:"varint,6,opt,name=Samples" json:"Samples,omitempty"`
}

func (m *QueryResult) Reset()                    { *m = QueryResult{} }
func (m *QueryResult) String() string            { return proto.CompactTextString(m) }
func (*QueryResult) ProtoMessage()               {}
func (*QueryResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *QueryResult) GetHistograms() map[string]*Histogram {
	if m != nil {
		return m.Histograms
	}
	return nil
}

func (m *QueryResult) GetGroupByKey() string {
	if m != nil {
		return m.GroupByKey
	}
	return ""
}

func (m *QueryResult) GetBinaryByKey() string {
	if m != nil {
		return m.BinaryByKey
	}
	return ""
}

func (m *QueryResult) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *QueryResult) GetSamples() int64 {
	if m != nil {
		return m.Samples
	}
	return 0
}

type ResultMap struct {
	Values map[string]*QueryResult `protobuf:"bytes,1,rep,name=values" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ResultMap) Reset()                    { *m = ResultMap{} }
func (m *ResultMap) String() string            { return proto.CompactTextString(m) }
func (*ResultMap) ProtoMessage()               {}
func (*ResultMap) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ResultMap) GetValues() map[string]*QueryResult {
	if m != nil {
		return m.Values
	}
	return nil
}

type SetField struct {
	Values []string `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
}

func (m *SetField) Reset()                    { *m = SetField{} }
func (m *SetField) String() string            { return proto.CompactTextString(m) }
func (*SetField) ProtoMessage()               {}
func (*SetField) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *SetField) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type Record struct {
	Timestamp int64                `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Strs      map[string]string    `protobuf:"bytes,2,rep,name=strs" json:"strs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Ints      map[string]int64     `protobuf:"bytes,3,rep,name=ints" json:"ints,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	Sets      map[string]*SetField `protobuf:"bytes,4,rep,name=sets" json:"sets,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Record) Reset()                    { *m = Record{} }
func (m *Record) String() string            { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()               {}
func (*Record) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Record) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Record) GetStrs() map[string]string {
	if m != nil {
		return m.Strs
	}
	return nil
}

func (m *Record) GetInts() map[string]int64 {
	if m != nil {
		return m.Ints
	}
	return nil
}

func (m *Record) GetSets() map[string]*SetField {
	if m != nil {
		return m.Sets
	}
	return nil
}

// QueryResults
type QueryResults struct {
	Cumulative   *QueryResult         `protobuf:"bytes,1,opt,name=cumulative" json:"cumulative,omitempty"`
	Results      *ResultMap           `protobuf:"bytes,2,opt,name=results" json:"results,omitempty"`
	TimeResults  map[int64]*ResultMap `protobuf:"bytes,3,rep,name=time_results,json=timeResults" json:"time_results,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	MatchedCount int64                `protobuf:"varint,4,opt,name=matched_count,json=matchedCount" json:"matched_count,omitempty"`
	Sorted       []*QueryResult       `protobuf:"bytes,5,rep,name=sorted" json:"sorted,omitempty"`
	Matched      []*Record            `protobuf:"bytes,6,rep,name=matched" json:"matched,omitempty"`
}

func (m *QueryResults) Reset()                    { *m = QueryResults{} }
func (m *QueryResults) String() string            { return proto.CompactTextString(m) }
func (*QueryResults) ProtoMessage()               {}
func (*QueryResults) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *QueryResults) GetCumulative() *QueryResult {
	if m != nil {
		return m.Cumulative
	}
	return nil
}

func (m *QueryResults) GetResults() *ResultMap {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *QueryResults) GetTimeResults() map[int64]*ResultMap {
	if m != nil {
		return m.TimeResults
	}
	return nil
}

func (m *QueryResults) GetMatchedCount() int64 {
	if m != nil {
		return m.MatchedCount
	}
	return 0
}

func (m *QueryResults) GetSorted() []*QueryResult {
	if m != nil {
		return m.Sorted
	}
	return nil
}

func (m *QueryResults) GetMatched() []*Record {
	if m != nil {
		return m.Matched
	}
	return nil
}

// QueryResponse
type QueryResponse struct {
	Results *QueryResults `protobuf:"bytes,1,opt,name=results" json:"results,omitempty"`
}

func (m *QueryResponse) Reset()                    { *m = QueryResponse{} }
func (m *QueryResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()               {}
func (*QueryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *QueryResponse) GetResults() *QueryResults {
	if m != nil {
		return m.Results
	}
	return nil
}

// ListTablesRequest
type ListTablesRequest struct {
}

func (m *ListTablesRequest) Reset()                    { *m = ListTablesRequest{} }
func (m *ListTablesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListTablesRequest) ProtoMessage()               {}
func (*ListTablesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

// ListTablesResponse
type ListTablesResponse struct {
	Tables []string `protobuf:"bytes,1,rep,name=tables" json:"tables,omitempty"`
}

func (m *ListTablesResponse) Reset()                    { *m = ListTablesResponse{} }
func (m *ListTablesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListTablesResponse) ProtoMessage()               {}
func (*ListTablesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *ListTablesResponse) GetTables() []string {
	if m != nil {
		return m.Tables
	}
	return nil
}

// GetTableRequest
type GetTableRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GetTableRequest) Reset()                    { *m = GetTableRequest{} }
func (m *GetTableRequest) String() string            { return proto.CompactTextString(m) }
func (*GetTableRequest) ProtoMessage()               {}
func (*GetTableRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *GetTableRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Table
type Table struct {
	Name              string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	StrColumns        []string `protobuf:"bytes,2,rep,name=str_columns,json=strColumns" json:"str_columns,omitempty"`
	IntColumns        []string `protobuf:"bytes,3,rep,name=int_columns,json=intColumns" json:"int_columns,omitempty"`
	SetColumns        []string `protobuf:"bytes,4,rep,name=set_columns,json=setColumns" json:"set_columns,omitempty"`
	Count             int64    `protobuf:"varint,5,opt,name=count" json:"count,omitempty"`
	StorageSize       int64    `protobuf:"varint,6,opt,name=storage_size,json=storageSize" json:"storage_size,omitempty"`
	AverageObjectSize int64    `protobuf:"varint,7,opt,name=average_object_size,json=averageObjectSize" json:"average_object_size,omitempty"`
}

func (m *Table) Reset()                    { *m = Table{} }
func (m *Table) String() string            { return proto.CompactTextString(m) }
func (*Table) ProtoMessage()               {}
func (*Table) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *Table) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Table) GetStrColumns() []string {
	if m != nil {
		return m.StrColumns
	}
	return nil
}

func (m *Table) GetIntColumns() []string {
	if m != nil {
		return m.IntColumns
	}
	return nil
}

func (m *Table) GetSetColumns() []string {
	if m != nil {
		return m.SetColumns
	}
	return nil
}

func (m *Table) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Table) GetStorageSize() int64 {
	if m != nil {
		return m.StorageSize
	}
	return 0
}

func (m *Table) GetAverageObjectSize() int64 {
	if m != nil {
		return m.AverageObjectSize
	}
	return 0
}

func init() {
	proto.RegisterType((*IngestRequest)(nil), "pb.IngestRequest")
	proto.RegisterType((*IngestResponse)(nil), "pb.IngestResponse")
	proto.RegisterType((*QueryFilter)(nil), "pb.QueryFilter")
	proto.RegisterType((*QueryRequest)(nil), "pb.QueryRequest")
	proto.RegisterType((*Histogram)(nil), "pb.Histogram")
	proto.RegisterType((*QueryResult)(nil), "pb.QueryResult")
	proto.RegisterType((*ResultMap)(nil), "pb.ResultMap")
	proto.RegisterType((*SetField)(nil), "pb.SetField")
	proto.RegisterType((*Record)(nil), "pb.Record")
	proto.RegisterType((*QueryResults)(nil), "pb.QueryResults")
	proto.RegisterType((*QueryResponse)(nil), "pb.QueryResponse")
	proto.RegisterType((*ListTablesRequest)(nil), "pb.ListTablesRequest")
	proto.RegisterType((*ListTablesResponse)(nil), "pb.ListTablesResponse")
	proto.RegisterType((*GetTableRequest)(nil), "pb.GetTableRequest")
	proto.RegisterType((*Table)(nil), "pb.Table")
	proto.RegisterEnum("pb.QueryType", QueryType_name, QueryType_value)
	proto.RegisterEnum("pb.QueryOp", QueryOp_name, QueryOp_value)
	proto.RegisterEnum("pb.QueryFilterOp", QueryFilterOp_name, QueryFilterOp_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Sybil service

type SybilClient interface {
	// Ingest inserts new data into a sybil dataset.
	Ingest(ctx context.Context, in *IngestRequest, opts ...grpc.CallOption) (*IngestResponse, error)
	// Query
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
	// ListTables
	ListTables(ctx context.Context, in *ListTablesRequest, opts ...grpc.CallOption) (*ListTablesResponse, error)
	// GetTable
	GetTable(ctx context.Context, in *GetTableRequest, opts ...grpc.CallOption) (*Table, error)
}

type sybilClient struct {
	cc *grpc.ClientConn
}

func NewSybilClient(cc *grpc.ClientConn) SybilClient {
	return &sybilClient{cc}
}

func (c *sybilClient) Ingest(ctx context.Context, in *IngestRequest, opts ...grpc.CallOption) (*IngestResponse, error) {
	out := new(IngestResponse)
	err := grpc.Invoke(ctx, "/pb.Sybil/Ingest", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sybilClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := grpc.Invoke(ctx, "/pb.Sybil/Query", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sybilClient) ListTables(ctx context.Context, in *ListTablesRequest, opts ...grpc.CallOption) (*ListTablesResponse, error) {
	out := new(ListTablesResponse)
	err := grpc.Invoke(ctx, "/pb.Sybil/ListTables", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sybilClient) GetTable(ctx context.Context, in *GetTableRequest, opts ...grpc.CallOption) (*Table, error) {
	out := new(Table)
	err := grpc.Invoke(ctx, "/pb.Sybil/GetTable", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Sybil service

type SybilServer interface {
	// Ingest inserts new data into a sybil dataset.
	Ingest(context.Context, *IngestRequest) (*IngestResponse, error)
	// Query
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
	// ListTables
	ListTables(context.Context, *ListTablesRequest) (*ListTablesResponse, error)
	// GetTable
	GetTable(context.Context, *GetTableRequest) (*Table, error)
}

func RegisterSybilServer(s *grpc.Server, srv SybilServer) {
	s.RegisterService(&_Sybil_serviceDesc, srv)
}

func _Sybil_Ingest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IngestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SybilServer).Ingest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Sybil/Ingest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SybilServer).Ingest(ctx, req.(*IngestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sybil_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SybilServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Sybil/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SybilServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sybil_ListTables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTablesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SybilServer).ListTables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Sybil/ListTables",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SybilServer).ListTables(ctx, req.(*ListTablesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sybil_GetTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SybilServer).GetTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Sybil/GetTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SybilServer).GetTable(ctx, req.(*GetTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Sybil_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Sybil",
	HandlerType: (*SybilServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ingest",
			Handler:    _Sybil_Ingest_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _Sybil_Query_Handler,
		},
		{
			MethodName: "ListTables",
			Handler:    _Sybil_ListTables_Handler,
		},
		{
			MethodName: "GetTable",
			Handler:    _Sybil_GetTable_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sybild.proto",
}

func init() { proto.RegisterFile("sybild.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1490 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x57, 0xdd, 0x72, 0xd3, 0xc6,
	0x17, 0xc7, 0x96, 0x3f, 0xa2, 0x63, 0x27, 0x56, 0x96, 0xfc, 0x83, 0x30, 0xff, 0x96, 0x44, 0x94,
	0x21, 0x93, 0x61, 0x9c, 0x92, 0x4e, 0x87, 0x52, 0x86, 0x76, 0x12, 0x10, 0xc1, 0x25, 0x71, 0xc8,
	0xda, 0xd0, 0x61, 0xa6, 0x83, 0x47, 0xb6, 0x17, 0xa3, 0x22, 0x4b, 0xaa, 0x76, 0x9d, 0xa9, 0x79,
	0x84, 0x5e, 0xf4, 0x45, 0xfa, 0x40, 0xbd, 0xea, 0x35, 0xd7, 0x7d, 0x83, 0xce, 0x9e, 0x5d, 0xc9,
	0x8a, 0x31, 0x65, 0x7a, 0xa5, 0xdd, 0xdf, 0x9e, 0x73, 0xf6, 0xec, 0x6f, 0xcf, 0xc7, 0x0a, 0xea,
	0x7c, 0x36, 0xf0, 0x83, 0x51, 0x2b, 0x4e, 0x22, 0x11, 0x91, 0x62, 0x3c, 0x68, 0x7e, 0x3d, 0xf6,
	0xc5, 0x9b, 0xe9, 0xa0, 0x35, 0x8c, 0x26, 0x7b, 0xe3, 0x28, 0xf0, 0xc2, 0xf1, 0x1e, 0x2e, 0x0e,
	0xa6, 0xaf, 0xf7, 0x62, 0x31, 0x8b, 0x19, 0xdf, 0xe3, 0x22, 0x99, 0x0e, 0x85, 0xfe, 0x28, 0x55,
	0xe7, 0x27, 0x58, 0x6d, 0x87, 0x63, 0xc6, 0x05, 0x65, 0xbf, 0x4c, 0x19, 0x17, 0xc4, 0x86, 0xea,
	0xc8, 0x13, 0x1e, 0x67, 0xc2, 0x2e, 0x6c, 0x15, 0x76, 0x4c, 0x9a, 0x4e, 0xc9, 0x1d, 0xa8, 0x26,
	0x6c, 0x18, 0x25, 0x23, 0x6e, 0x17, 0xb7, 0x8c, 0x9d, 0xda, 0xfe, 0x95, 0xd6, 0x38, 0x8a, 0xc6,
	0x01, 0x6b, 0xa5, 0x1b, 0xb5, 0xba, 0x68, 0x9a, 0xa6, 0x72, 0xce, 0x3d, 0x58, 0x4b, 0xad, 0xf3,
	0x38, 0x0a, 0x39, 0x23, 0xb7, 0xa0, 0x11, 0x4e, 0x27, 0x03, 0x96, 0xf4, 0xfd, 0x90, 0xb3, 0x44,
	0xb0, 0x11, 0x6e, 0x63, 0xd0, 0x35, 0x05, 0xb7, 0x35, 0xea, 0xbc, 0x82, 0xda, 0xd9, 0x94, 0x25,
	0xb3, 0xc7, 0x7e, 0x20, 0x58, 0x42, 0x36, 0xa1, 0x32, 0x8c, 0x82, 0xe9, 0x24, 0xd4, 0x5e, 0xe9,
	0x19, 0xd9, 0x86, 0x62, 0x14, 0xdb, 0xc5, 0xad, 0xc2, 0xce, 0xda, 0xfe, 0x7a, 0x2b, 0x1e, 0xb4,
	0x72, 0x4a, 0xa7, 0x31, 0x2d, 0x46, 0x31, 0xd9, 0x80, 0xf2, 0xb9, 0x17, 0x4c, 0x99, 0x6d, 0xa0,
	0xa6, 0x9a, 0x38, 0x7f, 0x94, 0xa0, 0x8e, 0xb2, 0x9f, 0x3e, 0xf8, 0x36, 0x94, 0x24, 0x81, 0x7a,
	0x97, 0xd5, 0x6c, 0x97, 0xde, 0x2c, 0x66, 0x14, 0x97, 0xe4, 0x1e, 0x81, 0x3f, 0xf1, 0x05, 0xee,
	0x61, 0x50, 0x35, 0x21, 0x04, 0x4a, 0x7e, 0x28, 0xb8, 0x5d, 0xda, 0x32, 0x76, 0x4c, 0x8a, 0x63,
	0x89, 0x71, 0x91, 0x70, 0xbb, 0xac, 0x30, 0x39, 0x96, 0x87, 0x1b, 0x27, 0xd1, 0x34, 0xe6, 0x76,
	0x05, 0x51, 0x3d, 0x23, 0xd7, 0xa1, 0x36, 0xf2, 0xb9, 0xf0, 0xc3, 0xa1, 0xe8, 0x0f, 0x66, 0x76,
	0x15, 0x17, 0x21, 0x85, 0x0e, 0x67, 0x68, 0x2c, 0x4a, 0x84, 0xbd, 0x82, 0x0e, 0xe3, 0x58, 0x2a,
	0x09, 0x7f, 0xc2, 0xfa, 0x9a, 0x2e, 0x13, 0x97, 0x40, 0x42, 0x0f, 0x15, 0x65, 0xa9, 0xc0, 0x60,
	0x3a, 0x7c, 0xcb, 0x84, 0x0d, 0xe8, 0x31, 0x0a, 0x1c, 0x22, 0x42, 0xae, 0x21, 0xa7, 0x35, 0x3c,
	0x6d, 0x2d, 0x3b, 0xad, 0x66, 0xf3, 0x06, 0xac, 0x06, 0xd1, 0xb8, 0xff, 0xc6, 0xe7, 0x22, 0x1a,
	0x27, 0xde, 0xc4, 0xae, 0x6f, 0x15, 0x76, 0x56, 0x68, 0x3d, 0x88, 0xc6, 0x4f, 0x52, 0x8c, 0x7c,
	0x09, 0x35, 0x3f, 0x14, 0xfd, 0xd7, 0x78, 0x0d, 0xdc, 0x5e, 0xc5, 0x70, 0x69, 0x2c, 0x5c, 0x0f,
	0x05, 0x3f, 0x14, 0x6a, 0xc8, 0xa5, 0x06, 0x17, 0x49, 0xa6, 0xb1, 0xf6, 0x11, 0x0d, 0x2e, 0x92,
	0xbc, 0x06, 0x9b, 0xef, 0xd1, 0xf8, 0x98, 0x06, 0xcb, 0xf6, 0xb8, 0x07, 0x57, 0xa5, 0x57, 0x99,
	0xeb, 0x9a, 0x81, 0x3e, 0xf7, 0xdf, 0x31, 0xdb, 0x42, 0x1a, 0x36, 0xfd, 0x50, 0x64, 0xc7, 0x50,
	0x74, 0x74, 0xfd, 0x77, 0xcc, 0x79, 0x6f, 0x80, 0x39, 0x3f, 0x1e, 0x81, 0xd2, 0x84, 0x79, 0x2a,
	0x14, 0x0b, 0x14, 0xc7, 0xc4, 0x02, 0x63, 0xe2, 0xfd, 0x8a, 0x31, 0x62, 0x50, 0x39, 0x44, 0xc4,
	0x0f, 0x75, 0x44, 0xc8, 0x21, 0x32, 0x1f, 0x09, 0x2f, 0xe8, 0x0f, 0xa3, 0x69, 0x28, 0xec, 0x92,
	0x66, 0x5e, 0x42, 0x0f, 0x25, 0x42, 0xb6, 0xa0, 0x16, 0xb3, 0x64, 0xc8, 0x42, 0xe1, 0x07, 0x4c,
	0xc5, 0x88, 0x41, 0xf3, 0x10, 0xf9, 0x4e, 0xf1, 0xa4, 0x3c, 0x57, 0xf1, 0x52, 0xdb, 0xff, 0x4c,
	0x9e, 0x3a, 0x73, 0x4f, 0x66, 0xa1, 0x72, 0x9d, 0xbb, 0xa1, 0x48, 0x66, 0xc8, 0x9a, 0x06, 0xa4,
	0xbe, 0xe4, 0x20, 0xd5, 0xaf, 0x2e, 0xd3, 0x6f, 0x87, 0xe2, 0xa2, 0xbe, 0x9f, 0x01, 0xf2, 0x08,
	0x89, 0x17, 0x8e, 0x59, 0x9f, 0x0b, 0x4f, 0x07, 0x9e, 0x41, 0x01, 0xa1, 0xae, 0x44, 0xc8, 0x35,
	0x30, 0x95, 0x00, 0x0b, 0x47, 0x18, 0x7c, 0x06, 0x5d, 0x41, 0xc0, 0x0d, 0x47, 0x32, 0x78, 0xb8,
	0x18, 0xf5, 0x47, 0xec, 0xdc, 0xf7, 0x84, 0x1f, 0x85, 0x18, 0x7c, 0x05, 0x5a, 0xe7, 0x62, 0xf4,
	0x28, 0xc5, 0x9a, 0x0f, 0xa0, 0xb1, 0x70, 0x02, 0x49, 0xe5, 0x5b, 0x36, 0xd3, 0x79, 0x29, 0x87,
	0xf3, 0xa4, 0x56, 0x84, 0xab, 0xc9, 0xb7, 0xc5, 0x6f, 0x0a, 0x52, 0x7d, 0xe1, 0x00, 0x79, 0x75,
	0xe3, 0x13, 0xea, 0xce, 0xef, 0x45, 0x5d, 0x78, 0x28, 0xe3, 0xd3, 0x40, 0x90, 0xef, 0x01, 0xb2,
	0x80, 0xe1, 0x76, 0x01, 0xf9, 0xba, 0x9e, 0x45, 0x99, 0x12, 0x9a, 0x73, 0x97, 0x32, 0x36, 0x57,
	0x21, 0x9f, 0x03, 0x1c, 0xc9, 0x74, 0x3e, 0x9c, 0x3d, 0x65, 0x33, 0x5d, 0x83, 0x72, 0x88, 0xbc,
	0xf3, 0x43, 0x3f, 0xf4, 0x92, 0x99, 0x12, 0x28, 0xa1, 0x40, 0x1e, 0x92, 0xce, 0x62, 0x78, 0xd8,
	0x65, 0xe5, 0xac, 0x8a, 0x15, 0x1b, 0xaa, 0x5d, 0x6f, 0x12, 0xcb, 0x38, 0xa9, 0x20, 0x9e, 0x4e,
	0x9b, 0xc7, 0xd0, 0x58, 0x70, 0x68, 0x09, 0x81, 0x37, 0xf2, 0x0c, 0xd4, 0x54, 0x55, 0xcb, 0xb4,
	0xf2, 0x84, 0xfc, 0x56, 0x00, 0x53, 0x1d, 0xf3, 0xc4, 0x8b, 0xc9, 0x1d, 0xa8, 0xe0, 0x52, 0x4a,
	0xc5, 0x55, 0xa9, 0x97, 0x2d, 0xb7, 0x5e, 0xe0, 0x9a, 0x22, 0x41, 0x0b, 0x36, 0x7f, 0x80, 0x5a,
	0x0e, 0x5e, 0xe2, 0xca, 0xcd, 0x8b, 0xae, 0x34, 0x16, 0xd8, 0xcd, 0x3b, 0xe3, 0xc0, 0x4a, 0x57,
	0x26, 0x34, 0x0b, 0x46, 0xb2, 0x6a, 0xe6, 0x5c, 0x31, 0xd3, 0xfd, 0x9c, 0xbf, 0x8a, 0x50, 0xa1,
	0xd8, 0x80, 0xc8, 0xff, 0xc1, 0x94, 0x75, 0x8d, 0x0b, 0x6f, 0x12, 0xeb, 0xeb, 0x9f, 0x03, 0x64,
	0x47, 0x97, 0x62, 0xd5, 0xcd, 0x36, 0xd4, 0x49, 0xa4, 0x9e, 0xcc, 0x20, 0x7d, 0x08, 0x55, 0xa0,
	0x77, 0x74, 0x21, 0x37, 0x3e, 0x90, 0x6c, 0x87, 0x69, 0x96, 0xa8, 0xf2, 0x2e, 0x6d, 0x32, 0x5d,
	0xf2, 0x17, 0x6c, 0x66, 0xf9, 0x84, 0x12, 0xcd, 0xbb, 0x60, 0x66, 0xdb, 0x7c, 0x2a, 0xc0, 0xcd,
	0x7c, 0x80, 0xdf, 0x05, 0x33, 0xdb, 0xf5, 0x3f, 0x65, 0x86, 0x0b, 0x66, 0xf7, 0x5f, 0x52, 0xca,
	0xb9, 0x78, 0x0d, 0x75, 0xe9, 0x7b, 0x4a, 0x76, 0xfe, 0x0e, 0xfe, 0x2e, 0x66, 0x9d, 0x53, 0x5e,
	0x0f, 0x27, 0x7b, 0x00, 0xc3, 0xe9, 0x64, 0x1a, 0x78, 0xc2, 0x3f, 0x67, 0x68, 0x71, 0xc9, 0x25,
	0xe6, 0x44, 0xc8, 0x2d, 0xf9, 0x92, 0x40, 0xdd, 0x7c, 0xf4, 0x65, 0x51, 0x44, 0xd3, 0x55, 0xf2,
	0x08, 0xea, 0xd8, 0xaa, 0x52, 0x69, 0xc5, 0xff, 0xf6, 0x82, 0x6d, 0xde, 0xea, 0xf9, 0x13, 0xa6,
	0xc7, 0x8a, 0x62, 0xec, 0x70, 0xa9, 0x7f, 0x37, 0x60, 0x75, 0xe2, 0x89, 0xe1, 0x1b, 0x36, 0xba,
	0x50, 0x78, 0xeb, 0x1a, 0x54, 0xe9, 0x74, 0x0b, 0x2a, 0xb2, 0x7d, 0xb2, 0x11, 0x56, 0xdd, 0x25,
	0x07, 0xd0, 0xcb, 0xe4, 0x0b, 0xa8, 0x6a, 0x45, 0x5d, 0x7d, 0x61, 0x7e, 0xc9, 0x34, 0x5d, 0x6a,
	0x9e, 0x80, 0xb5, 0xe8, 0xd4, 0x92, 0x32, 0xb4, 0x2c, 0x09, 0xe7, 0x34, 0xe4, 0x38, 0xbf, 0x0f,
	0xab, 0xa9, 0x2f, 0xea, 0x1d, 0xb5, 0x3b, 0xa7, 0x50, 0x11, 0x6e, 0x2d, 0x92, 0x92, 0xb1, 0xe8,
	0x5c, 0x86, 0xf5, 0x63, 0x9f, 0x8b, 0x9e, 0x37, 0x08, 0x18, 0xd7, 0xcf, 0x1d, 0xe7, 0x36, 0x90,
	0x3c, 0xa8, 0xcd, 0x6e, 0x42, 0x45, 0x20, 0x92, 0xe6, 0x94, 0x9a, 0x39, 0x37, 0xa1, 0x71, 0xc4,
	0x94, 0x70, 0xfa, 0x5e, 0x22, 0x50, 0x0a, 0xbd, 0x09, 0xd3, 0x11, 0x84, 0x63, 0xe7, 0x7d, 0x01,
	0xca, 0x28, 0xb4, 0x6c, 0x55, 0xf6, 0x0e, 0xd9, 0xbb, 0xd4, 0xc3, 0x44, 0xa5, 0x9d, 0x89, 0xcd,
	0x49, 0x3d, 0x4c, 0xb0, 0xb9, 0xc8, 0xe6, 0x94, 0x0a, 0x18, 0x4a, 0xc0, 0x0f, 0x45, 0x4e, 0x40,
	0xf6, 0xfc, 0x54, 0xa0, 0xa4, 0x2d, 0xb0, 0x4c, 0x60, 0x03, 0xca, 0xc3, 0x7c, 0xa9, 0xc4, 0x09,
	0xd9, 0x86, 0x3a, 0x17, 0x51, 0xe2, 0xc9, 0xb6, 0x25, 0x7b, 0xbd, 0xaa, 0x97, 0x35, 0x8d, 0xc9,
	0x06, 0x4f, 0x5a, 0x70, 0xd9, 0x3b, 0x67, 0x28, 0x12, 0x0d, 0x7e, 0x66, 0x43, 0xfd, 0x2a, 0xa8,
	0xa2, 0xe4, 0xba, 0x5e, 0x3a, 0xc5, 0x15, 0x29, 0xbf, 0xfb, 0x0a, 0xcc, 0xec, 0x0d, 0x48, 0x36,
	0x81, 0x9c, 0x3d, 0x77, 0xe9, 0xcb, 0x7e, 0xef, 0xe5, 0x33, 0xb7, 0xff, 0xbc, 0xf3, 0xb4, 0x73,
	0xfa, 0x63, 0xc7, 0xba, 0x44, 0x4c, 0x28, 0xf7, 0x0e, 0x0e, 0x8f, 0x5d, 0xab, 0x40, 0x1a, 0x50,
	0xeb, 0xb5, 0x4f, 0xdc, 0x7e, 0xd7, 0xa5, 0x6d, 0xb7, 0x6b, 0x15, 0x89, 0x05, 0xf5, 0x47, 0xed,
	0x6e, 0x8f, 0xb6, 0x0f, 0x9f, 0xf7, 0xda, 0xa7, 0x1d, 0xcb, 0x20, 0x35, 0xa8, 0x76, 0x0f, 0x4e,
	0x9e, 0x1d, 0xbb, 0x5d, 0xab, 0xb4, 0x7b, 0x1f, 0xaa, 0xfa, 0xd5, 0x45, 0x36, 0xc0, 0x52, 0xd6,
	0x4f, 0x9f, 0xe5, 0x6c, 0xd7, 0xa0, 0x7a, 0xf0, 0xc2, 0xa5, 0x07, 0x47, 0xd2, 0xfa, 0x2a, 0x98,
	0x4f, 0xda, 0xdd, 0xde, 0xe9, 0x11, 0x3d, 0x38, 0xb1, 0x8a, 0xbb, 0x91, 0x8e, 0x96, 0xf4, 0x19,
	0x4c, 0xae, 0xc1, 0x15, 0x65, 0xe2, 0x71, 0xfb, 0xb8, 0xe7, 0xd2, 0x8b, 0x96, 0x2a, 0x50, 0xa4,
	0xd2, 0x48, 0x15, 0x8c, 0x0e, 0x75, 0xad, 0xa2, 0x04, 0xdc, 0x33, 0xcb, 0x40, 0xc0, 0x3d, 0xb3,
	0x4a, 0x12, 0x38, 0xea, 0x59, 0x65, 0xf9, 0x3d, 0xee, 0x59, 0x15, 0xf9, 0x6d, 0x77, 0xac, 0x2a,
	0x0a, 0xb4, 0x3b, 0xd6, 0xca, 0xfe, 0x9f, 0x05, 0x28, 0x77, 0xe5, 0x1f, 0x89, 0xec, 0x0f, 0xea,
	0xc5, 0x4f, 0xf0, 0x35, 0x7e, 0xe1, 0xdf, 0xa2, 0x49, 0xf2, 0x90, 0x8a, 0x38, 0xe7, 0x12, 0x69,
	0x41, 0x19, 0xbd, 0x25, 0xf9, 0x10, 0x56, 0x0a, 0xeb, 0xf9, 0xa0, 0x4e, 0xe5, 0x1f, 0x00, 0xcc,
	0x23, 0x97, 0xfc, 0x4f, 0x8a, 0x7c, 0x10, 0xde, 0xcd, 0xcd, 0x45, 0x38, 0x53, 0xbf, 0x0d, 0x2b,
	0x69, 0x28, 0x93, 0xcb, 0x52, 0x6a, 0x21, 0xb0, 0x9b, 0xa6, 0x04, 0x11, 0x71, 0x2e, 0x0d, 0x2a,
	0xf8, 0x6f, 0xf3, 0xd5, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5e, 0xac, 0xf2, 0x2e, 0x71, 0x0d,
	0x00, 0x00,
}

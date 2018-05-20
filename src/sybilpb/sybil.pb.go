// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sybil.proto

/*
Package sybilpb is a generated protocol buffer package.

It is generated from these files:
	sybil.proto

It has these top-level messages:
	QueryFilter
	QueryRequest
	HistogramOptions
	Histogram
	ResultMap
	SetField
	FieldValue
	QueryResult
	QueryResponse
	Table
*/
package sybilpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/struct"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// QueryType defines the types of query that can be performed.
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

// QueryOp is the type of operation to perform in the query.
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

// QueryFilterOp is the operation to apply to a column as part of a QueryFilter.
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

// HistogramType describes the type of histogram.
type HistogramType int32

const (
	HistogramType_HISTOGRAM_TYPE_UNKNOWN HistogramType = 0
	HistogramType_NORMAL_HISTOGRAM       HistogramType = 1
	HistogramType_LOG_HISTOGRAM          HistogramType = 2
)

var HistogramType_name = map[int32]string{
	0: "HISTOGRAM_TYPE_UNKNOWN",
	1: "NORMAL_HISTOGRAM",
	2: "LOG_HISTOGRAM",
}
var HistogramType_value = map[string]int32{
	"HISTOGRAM_TYPE_UNKNOWN": 0,
	"NORMAL_HISTOGRAM":       1,
	"LOG_HISTOGRAM":          2,
}

func (x HistogramType) String() string {
	return proto.EnumName(HistogramType_name, int32(x))
}
func (HistogramType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// QueryFilter is a filter on a column.
type QueryFilter struct {
	Column string        `protobuf:"bytes,1,opt,name=column" json:"column,omitempty"`
	Op     QueryFilterOp `protobuf:"varint,2,opt,name=op,enum=sybilpb.QueryFilterOp" json:"op,omitempty"`
	Value  string        `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
}

func (m *QueryFilter) Reset()                    { *m = QueryFilter{} }
func (m *QueryFilter) String() string            { return proto.CompactTextString(m) }
func (*QueryFilter) ProtoMessage()               {}
func (*QueryFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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

// QueryRequest describes a query.
type QueryRequest struct {
	// Dataset is the name of the dataset.
	Dataset string `protobuf:"bytes,1,opt,name=dataset" json:"dataset,omitempty"`
	// The type of the query.
	Type QueryType `protobuf:"varint,2,opt,name=type,enum=sybilpb.QueryType" json:"type,omitempty"`
	// Limit number of results.
	Limit int64 `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
	// The integer fields to aggregate.
	Ints []string `protobuf:"bytes,4,rep,name=ints" json:"ints,omitempty"`
	// The string fields to aggregate.
	Strs []string `protobuf:"bytes,5,rep,name=strs" json:"strs,omitempty"`
	// The fields to group by.
	GroupBy         []string `protobuf:"bytes,6,rep,name=group_by,json=groupBy" json:"group_by,omitempty"`
	DistinctGroupBy []string `protobuf:"bytes,7,rep,name=distinct_group_by,json=distinctGroupBy" json:"distinct_group_by,omitempty"`
	// Field to sort by.
	SortBy string `protobuf:"bytes,8,opt,name=sort_by,json=sortBy" json:"sort_by,omitempty"`
	// Column to consider as the time column.
	TimeColumn string `protobuf:"bytes,9,opt,name=time_column,json=timeColumn" json:"time_column,omitempty"`
	// Time bucket size in seconds.
	TimeBucket int64 `protobuf:"varint,10,opt,name=time_bucket,json=timeBucket" json:"time_bucket,omitempty"`
	// The column to interpret as the weight.
	WeightColumn string `protobuf:"bytes,11,opt,name=weight_column,json=weightColumn" json:"weight_column,omitempty"`
	// The operation to run.
	Op QueryOp `protobuf:"varint,12,opt,name=op,enum=sybilpb.QueryOp" json:"op,omitempty"`
	// Filters on int columns.
	IntFilters []*QueryFilter `protobuf:"bytes,13,rep,name=int_filters,json=intFilters" json:"int_filters,omitempty"`
	// Filters on string columns.
	StrFilters []*QueryFilter `protobuf:"bytes,14,rep,name=str_filters,json=strFilters" json:"str_filters,omitempty"`
	// Filters on set columns.
	SetFilters []*QueryFilter `protobuf:"bytes,15,rep,name=set_filters,json=setFilters" json:"set_filters,omitempty"`
	// If type is DISTRIBUTION then this field controls hisogram options.
	HistogramOptions *HistogramOptions `protobuf:"bytes,16,opt,name=histogram_options,json=histogramOptions" json:"histogram_options,omitempty"`
	// If true, the ingestion log is also read to produce results.
	ReadIngestionLog bool `protobuf:"varint,17,opt,name=read_ingestion_log,json=readIngestionLog" json:"read_ingestion_log,omitempty"`
}

func (m *QueryRequest) Reset()                    { *m = QueryRequest{} }
func (m *QueryRequest) String() string            { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()               {}
func (*QueryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

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

func (m *QueryRequest) GetGroupBy() []string {
	if m != nil {
		return m.GroupBy
	}
	return nil
}

func (m *QueryRequest) GetDistinctGroupBy() []string {
	if m != nil {
		return m.DistinctGroupBy
	}
	return nil
}

func (m *QueryRequest) GetSortBy() string {
	if m != nil {
		return m.SortBy
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

func (m *QueryRequest) GetWeightColumn() string {
	if m != nil {
		return m.WeightColumn
	}
	return ""
}

func (m *QueryRequest) GetOp() QueryOp {
	if m != nil {
		return m.Op
	}
	return QueryOp_QUERY_OP_UNKNOWN
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

func (m *QueryRequest) GetHistogramOptions() *HistogramOptions {
	if m != nil {
		return m.HistogramOptions
	}
	return nil
}

func (m *QueryRequest) GetReadIngestionLog() bool {
	if m != nil {
		return m.ReadIngestionLog
	}
	return false
}

// HistogramOptions
type HistogramOptions struct {
	Type       HistogramType `protobuf:"varint,1,opt,name=type,enum=sybilpb.HistogramType" json:"type,omitempty"`
	BucketSize int64         `protobuf:"varint,2,opt,name=bucket_size,json=bucketSize" json:"bucket_size,omitempty"`
}

func (m *HistogramOptions) Reset()                    { *m = HistogramOptions{} }
func (m *HistogramOptions) String() string            { return proto.CompactTextString(m) }
func (*HistogramOptions) ProtoMessage()               {}
func (*HistogramOptions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *HistogramOptions) GetType() HistogramType {
	if m != nil {
		return m.Type
	}
	return HistogramType_HISTOGRAM_TYPE_UNKNOWN
}

func (m *HistogramOptions) GetBucketSize() int64 {
	if m != nil {
		return m.BucketSize
	}
	return 0
}

// Histogram describes a distribution of values.
type Histogram struct {
	Mean         float64 `protobuf:"fixed64,1,opt,name=mean" json:"mean,omitempty"`
	Percentiles  []int64 `protobuf:"varint,5,rep,packed,name=percentiles" json:"percentiles,omitempty"`
	Buckets      []int64 `protobuf:"varint,6,rep,packed,name=buckets" json:"buckets,omitempty"`
	StdDeviation float64 `protobuf:"fixed64,9,opt,name=std_deviation,json=stdDeviation" json:"std_deviation,omitempty"`
}

func (m *Histogram) Reset()                    { *m = Histogram{} }
func (m *Histogram) String() string            { return proto.CompactTextString(m) }
func (*Histogram) ProtoMessage()               {}
func (*Histogram) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Histogram) GetMean() float64 {
	if m != nil {
		return m.Mean
	}
	return 0
}

func (m *Histogram) GetPercentiles() []int64 {
	if m != nil {
		return m.Percentiles
	}
	return nil
}

func (m *Histogram) GetBuckets() []int64 {
	if m != nil {
		return m.Buckets
	}
	return nil
}

func (m *Histogram) GetStdDeviation() float64 {
	if m != nil {
		return m.StdDeviation
	}
	return 0
}

// ResultMap
type ResultMap struct {
	Values map[string]*QueryResult `protobuf:"bytes,1,rep,name=values" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ResultMap) Reset()                    { *m = ResultMap{} }
func (m *ResultMap) String() string            { return proto.CompactTextString(m) }
func (*ResultMap) ProtoMessage()               {}
func (*ResultMap) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ResultMap) GetValues() map[string]*QueryResult {
	if m != nil {
		return m.Values
	}
	return nil
}

// SetField
type SetField struct {
	Values []string `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
}

func (m *SetField) Reset()                    { *m = SetField{} }
func (m *SetField) String() string            { return proto.CompactTextString(m) }
func (*SetField) ProtoMessage()               {}
func (*SetField) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SetField) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type FieldValue struct {
	// Types that are valid to be assigned to Value:
	//	*FieldValue_Avg
	//	*FieldValue_Hist
	//	*FieldValue_Str
	Value isFieldValue_Value `protobuf_oneof:"value"`
}

func (m *FieldValue) Reset()                    { *m = FieldValue{} }
func (m *FieldValue) String() string            { return proto.CompactTextString(m) }
func (*FieldValue) ProtoMessage()               {}
func (*FieldValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type isFieldValue_Value interface {
	isFieldValue_Value()
}

type FieldValue_Avg struct {
	Avg float64 `protobuf:"fixed64,1,opt,name=avg,oneof"`
}
type FieldValue_Hist struct {
	Hist *Histogram `protobuf:"bytes,2,opt,name=hist,oneof"`
}
type FieldValue_Str struct {
	Str string `protobuf:"bytes,3,opt,name=str,oneof"`
}

func (*FieldValue_Avg) isFieldValue_Value()  {}
func (*FieldValue_Hist) isFieldValue_Value() {}
func (*FieldValue_Str) isFieldValue_Value()  {}

func (m *FieldValue) GetValue() isFieldValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *FieldValue) GetAvg() float64 {
	if x, ok := m.GetValue().(*FieldValue_Avg); ok {
		return x.Avg
	}
	return 0
}

func (m *FieldValue) GetHist() *Histogram {
	if x, ok := m.GetValue().(*FieldValue_Hist); ok {
		return x.Hist
	}
	return nil
}

func (m *FieldValue) GetStr() string {
	if x, ok := m.GetValue().(*FieldValue_Str); ok {
		return x.Str
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*FieldValue) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _FieldValue_OneofMarshaler, _FieldValue_OneofUnmarshaler, _FieldValue_OneofSizer, []interface{}{
		(*FieldValue_Avg)(nil),
		(*FieldValue_Hist)(nil),
		(*FieldValue_Str)(nil),
	}
}

func _FieldValue_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*FieldValue)
	// value
	switch x := m.Value.(type) {
	case *FieldValue_Avg:
		b.EncodeVarint(1<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.Avg))
	case *FieldValue_Hist:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Hist); err != nil {
			return err
		}
	case *FieldValue_Str:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Str)
	case nil:
	default:
		return fmt.Errorf("FieldValue.Value has unexpected type %T", x)
	}
	return nil
}

func _FieldValue_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*FieldValue)
	switch tag {
	case 1: // value.avg
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Value = &FieldValue_Avg{math.Float64frombits(x)}
		return true, err
	case 2: // value.hist
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Histogram)
		err := b.DecodeMessage(msg)
		m.Value = &FieldValue_Hist{msg}
		return true, err
	case 3: // value.str
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &FieldValue_Str{x}
		return true, err
	default:
		return false, nil
	}
}

func _FieldValue_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*FieldValue)
	// value
	switch x := m.Value.(type) {
	case *FieldValue_Avg:
		n += proto.SizeVarint(1<<3 | proto.WireFixed64)
		n += 8
	case *FieldValue_Hist:
		s := proto.Size(x.Hist)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *FieldValue_Str:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Str)))
		n += len(x.Str)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type QueryResult struct {
	Values   map[string]*FieldValue `protobuf:"bytes,1,rep,name=values" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Count    int64                  `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
	Samples  int64                  `protobuf:"varint,3,opt,name=samples" json:"samples,omitempty"`
	Distinct int64                  `protobuf:"varint,4,opt,name=distinct" json:"distinct,omitempty"`
	// if a time-series query, the time bucket for the result.
	Time int64 `protobuf:"varint,5,opt,name=time" json:"time,omitempty"`
}

func (m *QueryResult) Reset()                    { *m = QueryResult{} }
func (m *QueryResult) String() string            { return proto.CompactTextString(m) }
func (*QueryResult) ProtoMessage()               {}
func (*QueryResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *QueryResult) GetValues() map[string]*FieldValue {
	if m != nil {
		return m.Values
	}
	return nil
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

func (m *QueryResult) GetDistinct() int64 {
	if m != nil {
		return m.Distinct
	}
	return 0
}

func (m *QueryResult) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

// QueryResponse is the response type for a query.
type QueryResponse struct {
	Results    []*QueryResult            `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
	Cumulative *QueryResult              `protobuf:"bytes,2,opt,name=cumulative" json:"cumulative,omitempty"`
	Samples    []*google_protobuf.Struct `protobuf:"bytes,3,rep,name=samples" json:"samples,omitempty"`
}

func (m *QueryResponse) Reset()                    { *m = QueryResponse{} }
func (m *QueryResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()               {}
func (*QueryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *QueryResponse) GetResults() []*QueryResult {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *QueryResponse) GetCumulative() *QueryResult {
	if m != nil {
		return m.Cumulative
	}
	return nil
}

func (m *QueryResponse) GetSamples() []*google_protobuf.Struct {
	if m != nil {
		return m.Samples
	}
	return nil
}

// Table (aka dataset) is a collection of records in a sybil database.
type Table struct {
	// the name of the dataset.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// the set of string column names.
	StrColumns []string `protobuf:"bytes,2,rep,name=str_columns,json=strColumns" json:"str_columns,omitempty"`
	// the set of integer column names.
	IntColumns []string `protobuf:"bytes,3,rep,name=int_columns,json=intColumns" json:"int_columns,omitempty"`
	// the set of set column names.
	SetColumns []string `protobuf:"bytes,4,rep,name=set_columns,json=setColumns" json:"set_columns,omitempty"`
	// the approximate count of samples.
	Count int64 `protobuf:"varint,5,opt,name=count" json:"count,omitempty"`
	// the approximate size in terms of disk utilization.
	StorageSize int64 `protobuf:"varint,6,opt,name=storage_size,json=storageSize" json:"storage_size,omitempty"`
	// the mean sample storage size.
	AverageObjectSize float64 `protobuf:"fixed64,7,opt,name=average_object_size,json=averageObjectSize" json:"average_object_size,omitempty"`
}

func (m *Table) Reset()                    { *m = Table{} }
func (m *Table) String() string            { return proto.CompactTextString(m) }
func (*Table) ProtoMessage()               {}
func (*Table) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

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

func (m *Table) GetAverageObjectSize() float64 {
	if m != nil {
		return m.AverageObjectSize
	}
	return 0
}

func init() {
	proto.RegisterType((*QueryFilter)(nil), "sybilpb.QueryFilter")
	proto.RegisterType((*QueryRequest)(nil), "sybilpb.QueryRequest")
	proto.RegisterType((*HistogramOptions)(nil), "sybilpb.HistogramOptions")
	proto.RegisterType((*Histogram)(nil), "sybilpb.Histogram")
	proto.RegisterType((*ResultMap)(nil), "sybilpb.ResultMap")
	proto.RegisterType((*SetField)(nil), "sybilpb.SetField")
	proto.RegisterType((*FieldValue)(nil), "sybilpb.FieldValue")
	proto.RegisterType((*QueryResult)(nil), "sybilpb.QueryResult")
	proto.RegisterType((*QueryResponse)(nil), "sybilpb.QueryResponse")
	proto.RegisterType((*Table)(nil), "sybilpb.Table")
	proto.RegisterEnum("sybilpb.QueryType", QueryType_name, QueryType_value)
	proto.RegisterEnum("sybilpb.QueryOp", QueryOp_name, QueryOp_value)
	proto.RegisterEnum("sybilpb.QueryFilterOp", QueryFilterOp_name, QueryFilterOp_value)
	proto.RegisterEnum("sybilpb.HistogramType", HistogramType_name, HistogramType_value)
}

func init() { proto.RegisterFile("sybil.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1124 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x0e, 0x45, 0x49, 0x94, 0x86, 0x52, 0x42, 0x6f, 0x0c, 0x9b, 0x71, 0x8b, 0x46, 0x55, 0x81,
	0x40, 0x15, 0x0a, 0x05, 0x75, 0x7f, 0x10, 0xb4, 0x27, 0xa9, 0xa1, 0x6d, 0xa1, 0xb2, 0x64, 0xaf,
	0xe8, 0x14, 0xb9, 0x94, 0xa0, 0xa4, 0x8d, 0xcc, 0x86, 0x22, 0x59, 0xee, 0xd2, 0x85, 0x72, 0xea,
	0x33, 0x14, 0x7d, 0x88, 0xbe, 0x59, 0x1f, 0xa1, 0xd7, 0x62, 0x7f, 0x48, 0xd3, 0xaa, 0xe1, 0x9e,
	0xb8, 0xf3, 0xed, 0x37, 0x33, 0xbb, 0xf3, 0xcd, 0x0e, 0xc1, 0xa4, 0xdb, 0x45, 0x10, 0x0e, 0x92,
	0x34, 0x66, 0x31, 0x32, 0x84, 0x91, 0x2c, 0x8e, 0x3e, 0x5e, 0xc7, 0xf1, 0x3a, 0x24, 0x2f, 0x05,
	0xbc, 0xc8, 0xde, 0xbd, 0xa4, 0x2c, 0xcd, 0x96, 0x4c, 0xd2, 0xba, 0x4b, 0x30, 0x2f, 0x33, 0x92,
	0x6e, 0x4f, 0x82, 0x90, 0x91, 0x14, 0x1d, 0x40, 0x7d, 0x19, 0x87, 0xd9, 0x26, 0xb2, 0xb5, 0x8e,
	0xd6, 0x6b, 0x62, 0x65, 0xa1, 0x17, 0x50, 0x89, 0x13, 0xbb, 0xd2, 0xd1, 0x7a, 0x8f, 0x8f, 0x0f,
	0x06, 0x2a, 0xf4, 0xa0, 0xe4, 0x39, 0x4b, 0x70, 0x25, 0x4e, 0xd0, 0x3e, 0xd4, 0x6e, 0xfc, 0x30,
	0x23, 0xb6, 0x2e, 0xdc, 0xa5, 0xd1, 0xfd, 0xa3, 0x06, 0x2d, 0xc1, 0xc5, 0xe4, 0xd7, 0x8c, 0x50,
	0x86, 0x6c, 0x30, 0x56, 0x3e, 0xf3, 0x29, 0x61, 0x2a, 0x4f, 0x6e, 0xa2, 0x17, 0x50, 0x65, 0xdb,
	0x84, 0xa8, 0x54, 0xe8, 0x6e, 0x2a, 0x77, 0x9b, 0x10, 0x2c, 0xf6, 0x79, 0xa2, 0x30, 0xd8, 0x04,
	0x4c, 0x24, 0xd2, 0xb1, 0x34, 0x10, 0x82, 0x6a, 0x10, 0x31, 0x6a, 0x57, 0x3b, 0x7a, 0xaf, 0x89,
	0xc5, 0x9a, 0x63, 0x94, 0xa5, 0xd4, 0xae, 0x49, 0x8c, 0xaf, 0xd1, 0x33, 0x68, 0xac, 0xd3, 0x38,
	0x4b, 0xbc, 0xc5, 0xd6, 0xae, 0x0b, 0xdc, 0x10, 0xf6, 0x68, 0x8b, 0xfa, 0xb0, 0xb7, 0x0a, 0x28,
	0x0b, 0xa2, 0x25, 0xf3, 0x0a, 0x8e, 0x21, 0x38, 0x4f, 0xf2, 0x8d, 0x53, 0xc5, 0x3d, 0x04, 0x83,
	0xc6, 0x29, 0xe3, 0x8c, 0x86, 0x2c, 0x17, 0x37, 0x47, 0x5b, 0xf4, 0x1c, 0x4c, 0x16, 0x6c, 0x88,
	0xa7, 0x6a, 0xd9, 0x14, 0x9b, 0xc0, 0xa1, 0x1f, 0x64, 0x3d, 0x73, 0xc2, 0x22, 0x5b, 0xbe, 0x27,
	0xcc, 0x06, 0x71, 0x09, 0x41, 0x18, 0x09, 0x04, 0x7d, 0x06, 0xed, 0xdf, 0x48, 0xb0, 0xbe, 0x66,
	0x79, 0x0c, 0x53, 0xc4, 0x68, 0x49, 0x50, 0x45, 0xe9, 0x08, 0x55, 0x5a, 0xa2, 0x54, 0xd6, 0xdd,
	0x52, 0x29, 0x3d, 0xbe, 0x01, 0x33, 0x88, 0x98, 0xf7, 0x4e, 0x68, 0x44, 0xed, 0x76, 0x47, 0xef,
	0x99, 0xc7, 0xfb, 0xf7, 0x09, 0x88, 0x21, 0x88, 0x98, 0x5c, 0x52, 0xee, 0x46, 0x59, 0x5a, 0xb8,
	0x3d, 0x7e, 0xc8, 0x8d, 0xb2, 0xb4, 0xec, 0x46, 0x6e, 0xb3, 0x3d, 0x79, 0xd0, 0x8d, 0x14, 0xd9,
	0x4e, 0x60, 0xef, 0x3a, 0xa0, 0x2c, 0x5e, 0xa7, 0xfe, 0xc6, 0x8b, 0x13, 0x16, 0xc4, 0x11, 0xb5,
	0xad, 0x8e, 0xd6, 0x33, 0x8f, 0x9f, 0x15, 0xce, 0x67, 0x39, 0x63, 0x26, 0x09, 0xd8, 0xba, 0xde,
	0x41, 0xd0, 0x17, 0x80, 0x52, 0xe2, 0xaf, 0xbc, 0x20, 0x5a, 0x13, 0xca, 0x21, 0x2f, 0x8c, 0xd7,
	0xf6, 0x5e, 0x47, 0xeb, 0x35, 0xb0, 0xc5, 0x77, 0xc6, 0xf9, 0xc6, 0x24, 0x5e, 0x77, 0x3d, 0xb0,
	0x76, 0x63, 0xa2, 0xbe, 0xea, 0x3e, 0x6d, 0xa7, 0xd1, 0x0b, 0x62, 0xa9, 0x03, 0x9f, 0x83, 0x29,
	0xd5, 0xf3, 0x68, 0xf0, 0x41, 0x36, 0xac, 0x8e, 0x41, 0x42, 0xf3, 0xe0, 0x03, 0xe9, 0xfe, 0xae,
	0x41, 0xb3, 0x70, 0xe4, 0x6d, 0xb8, 0x21, 0xbe, 0x7c, 0x57, 0x1a, 0x16, 0x6b, 0xd4, 0x01, 0x33,
	0x21, 0xe9, 0x92, 0x44, 0x2c, 0x08, 0x89, 0xec, 0x50, 0x1d, 0x97, 0x21, 0xfe, 0x50, 0x64, 0x44,
	0x2a, 0xfa, 0x54, 0xc7, 0xb9, 0xc9, 0x1b, 0x84, 0xb2, 0x95, 0xb7, 0x22, 0x37, 0x81, 0xcf, 0x0f,
	0x2f, 0x9a, 0x4c, 0xc3, 0x2d, 0xca, 0x56, 0xaf, 0x73, 0xac, 0xfb, 0xa7, 0x06, 0x4d, 0x4c, 0x68,
	0x16, 0xb2, 0x73, 0x3f, 0x41, 0xdf, 0x42, 0x5d, 0xbc, 0x47, 0x6a, 0x6b, 0x42, 0x99, 0x4f, 0x8a,
	0xfb, 0x15, 0x9c, 0xc1, 0x1b, 0x41, 0x70, 0x22, 0x96, 0x6e, 0xb1, 0x62, 0x1f, 0xcd, 0xc0, 0x2c,
	0xc1, 0xc8, 0x02, 0xfd, 0x3d, 0xd9, 0xaa, 0x87, 0xcb, 0x97, 0xa8, 0x9f, 0xbf, 0xfa, 0x8a, 0x10,
	0x6d, 0x47, 0x71, 0x19, 0x5c, 0xcd, 0x82, 0xef, 0x2a, 0xaf, 0xb4, 0x6e, 0x17, 0x1a, 0x73, 0x2e,
	0x3f, 0x09, 0x57, 0x7c, 0xe2, 0x94, 0x0e, 0xd5, 0xcc, 0x93, 0x76, 0xd7, 0x00, 0x82, 0x20, 0x32,
	0x23, 0x04, 0xba, 0x7f, 0xb3, 0x96, 0xc5, 0x3b, 0x7b, 0x84, 0xb9, 0x81, 0x7a, 0x50, 0xe5, 0x2d,
	0xa0, 0x92, 0xa2, 0xff, 0x8a, 0x75, 0xf6, 0x08, 0x0b, 0x06, 0xf7, 0xa6, 0x2c, 0x95, 0x33, 0x89,
	0x7b, 0x53, 0x96, 0x8e, 0x0c, 0x75, 0xe6, 0xee, 0x3f, 0x9a, 0x1a, 0x81, 0xf2, 0x9c, 0xe8, 0xd5,
	0x4e, 0x95, 0x3a, 0xf7, 0xdd, 0xe6, 0xbe, 0x3a, 0xf1, 0x99, 0xb4, 0x8c, 0xb3, 0x88, 0xa9, 0x5e,
	0x90, 0x06, 0x97, 0x90, 0xfa, 0x9b, 0x84, 0x0b, 0x2c, 0x67, 0x55, 0x6e, 0xa2, 0x23, 0x68, 0xe4,
	0x13, 0xc5, 0xae, 0x8a, 0xad, 0xc2, 0xe6, 0xed, 0xc2, 0xa7, 0x81, 0x5d, 0x13, 0xb8, 0x58, 0x1f,
	0x4d, 0xff, 0x4f, 0x87, 0xcf, 0xef, 0xea, 0xf0, 0xb4, 0x38, 0xf9, 0x6d, 0x25, 0xcb, 0x32, 0xfc,
	0xa5, 0x41, 0x3b, 0xbf, 0x53, 0x12, 0x47, 0x94, 0xa0, 0x01, 0x18, 0xa9, 0xb8, 0x5f, 0x7e, 0xf9,
	0xfb, 0xa5, 0xcc, 0x49, 0xe8, 0x6b, 0x80, 0x65, 0xb6, 0xc9, 0x42, 0x9f, 0x05, 0x37, 0x0f, 0xab,
	0x5f, 0xe2, 0xa1, 0x2f, 0xcb, 0x15, 0xe1, 0x59, 0x0e, 0x07, 0xf2, 0x1f, 0x35, 0xc8, 0xff, 0x51,
	0x83, 0xb9, 0xf8, 0x47, 0x15, 0xa5, 0xea, 0xfe, 0xad, 0x41, 0xcd, 0xf5, 0x17, 0x21, 0xef, 0x84,
	0x6a, 0xe4, 0x6f, 0x88, 0xba, 0xb6, 0x58, 0xf3, 0xa7, 0xc8, 0xc7, 0x95, 0x9c, 0x94, 0xd4, 0xae,
	0x88, 0x46, 0xe2, 0x83, 0x49, 0xce, 0x49, 0xca, 0x09, 0x7c, 0x0c, 0xe6, 0x04, 0x5d, 0x12, 0x82,
	0x88, 0x95, 0x08, 0x7c, 0x72, 0xe5, 0x84, 0xaa, 0x8a, 0x40, 0x0a, 0x42, 0xa1, 0x6d, 0xad, 0xac,
	0xed, 0xa7, 0xd0, 0xa2, 0x2c, 0x4e, 0xfd, 0x35, 0x91, 0x43, 0xa0, 0x2e, 0x36, 0x4d, 0x85, 0xf1,
	0x29, 0x80, 0x06, 0xf0, 0xd4, 0xbf, 0x21, 0x82, 0x12, 0x2f, 0x7e, 0x21, 0x4b, 0x35, 0x2e, 0x0c,
	0xf1, 0x5a, 0xf7, 0xd4, 0xd6, 0x4c, 0xec, 0x70, 0x7e, 0xff, 0x67, 0x68, 0x16, 0xff, 0x3a, 0x74,
	0x00, 0xe8, 0xf2, 0xca, 0xc1, 0x6f, 0x3d, 0xf7, 0xed, 0x85, 0xe3, 0x5d, 0x4d, 0x7f, 0x9c, 0xce,
	0x7e, 0x9a, 0x5a, 0x8f, 0x50, 0x13, 0x6a, 0xee, 0x70, 0x34, 0x71, 0x2c, 0x0d, 0x3d, 0x01, 0xd3,
	0x1d, 0x9f, 0x3b, 0xde, 0xdc, 0xc1, 0x63, 0x67, 0x6e, 0x55, 0x90, 0x05, 0xad, 0xd7, 0xe3, 0xb9,
	0x8b, 0xc7, 0xa3, 0x2b, 0x77, 0x3c, 0x9b, 0x5a, 0x3a, 0x32, 0xc1, 0x98, 0x0f, 0xcf, 0x2f, 0x26,
	0xce, 0xdc, 0xaa, 0xf6, 0xbf, 0x07, 0x43, 0xfd, 0x20, 0xd0, 0x3e, 0x58, 0x32, 0xfa, 0xec, 0xa2,
	0x14, 0xdb, 0x04, 0x63, 0xf8, 0xc6, 0xc1, 0xc3, 0x53, 0x1e, 0xbd, 0x0d, 0xcd, 0xb3, 0xf1, 0xdc,
	0x9d, 0x9d, 0xe2, 0xe1, 0xb9, 0x55, 0xe9, 0xc7, 0xaa, 0x61, 0xf2, 0x7f, 0x3e, 0xfa, 0x08, 0x0e,
	0x65, 0x88, 0x93, 0xf1, 0xc4, 0x75, 0xf0, 0xdd, 0x48, 0x75, 0xa8, 0x60, 0x1e, 0xc4, 0x00, 0x7d,
	0x8a, 0x1d, 0xab, 0xc2, 0x01, 0xe7, 0xd2, 0xd2, 0x05, 0xe0, 0x5c, 0x5a, 0x55, 0x0e, 0x9c, 0xba,
	0x56, 0x8d, 0x7f, 0x27, 0xae, 0x55, 0xe7, 0xdf, 0xf1, 0xd4, 0x32, 0x04, 0x61, 0x3c, 0xb5, 0x1a,
	0x7d, 0x17, 0xda, 0x77, 0x66, 0x2f, 0x3a, 0x82, 0x83, 0xe2, 0x40, 0xbb, 0x55, 0xd9, 0x07, 0x6b,
	0x3a, 0xc3, 0xe7, 0xc3, 0x89, 0x77, 0x7b, 0x66, 0x0d, 0xed, 0x41, 0x7b, 0x32, 0x3b, 0x2d, 0x41,
	0x95, 0x45, 0x5d, 0xf4, 0xd9, 0x57, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x7b, 0xea, 0x20, 0xa5,
	0x31, 0x09, 0x00, 0x00,
}

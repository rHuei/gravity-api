// Code generated by protoc-gen-go. DO NOT EDIT.
// source: querykit.proto

package gravity_api_querykit

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type DataType int32

const (
	DataType_BOOLEAN DataType = 0
	DataType_BINARY  DataType = 1
	DataType_STRING  DataType = 2
	DataType_UINT64  DataType = 3
	DataType_INT64   DataType = 4
	DataType_FLOAT64 DataType = 5
	DataType_ARRAY   DataType = 6
	DataType_MAP     DataType = 7
)

var DataType_name = map[int32]string{
	0: "BOOLEAN",
	1: "BINARY",
	2: "STRING",
	3: "UINT64",
	4: "INT64",
	5: "FLOAT64",
	6: "ARRAY",
	7: "MAP",
}

var DataType_value = map[string]int32{
	"BOOLEAN": 0,
	"BINARY":  1,
	"STRING":  2,
	"UINT64":  3,
	"INT64":   4,
	"FLOAT64": 5,
	"ARRAY":   6,
	"MAP":     7,
}

func (x DataType) String() string {
	return proto.EnumName(DataType_name, int32(x))
}

func (DataType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eda43376942b3faf, []int{0}
}

type Operator int32

const (
	Operator_EQUAL         Operator = 0
	Operator_GREATER_THAN  Operator = 1
	Operator_GREATER_EQUAL Operator = 2
	Operator_LESS_THAN     Operator = 3
	Operator_LESS_EQUAL    Operator = 4
	Operator_AND           Operator = 5
	Operator_OR            Operator = 6
	Operator_NOT_EQUAL     Operator = 7
	Operator_IS_EXIST      Operator = 8
)

var Operator_name = map[int32]string{
	0: "EQUAL",
	1: "GREATER_THAN",
	2: "GREATER_EQUAL",
	3: "LESS_THAN",
	4: "LESS_EQUAL",
	5: "AND",
	6: "OR",
	7: "NOT_EQUAL",
	8: "IS_EXIST",
}

var Operator_value = map[string]int32{
	"EQUAL":         0,
	"GREATER_THAN":  1,
	"GREATER_EQUAL": 2,
	"LESS_THAN":     3,
	"LESS_EQUAL":    4,
	"AND":           5,
	"OR":            6,
	"NOT_EQUAL":     7,
	"IS_EXIST":      8,
}

func (x Operator) String() string {
	return proto.EnumName(Operator_name, int32(x))
}

func (Operator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eda43376942b3faf, []int{1}
}

type Record struct {
	Fields               []*Field `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_eda43376942b3faf, []int{0}
}

func (m *Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Record.Unmarshal(m, b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Record.Marshal(b, m, deterministic)
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return xxx_messageInfo_Record.Size(m)
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetFields() []*Field {
	if m != nil {
		return m.Fields
	}
	return nil
}

type Field struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                *Value   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Field) Reset()         { *m = Field{} }
func (m *Field) String() string { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()    {}
func (*Field) Descriptor() ([]byte, []int) {
	return fileDescriptor_eda43376942b3faf, []int{1}
}

func (m *Field) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Field.Unmarshal(m, b)
}
func (m *Field) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Field.Marshal(b, m, deterministic)
}
func (m *Field) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Field.Merge(m, src)
}
func (m *Field) XXX_Size() int {
	return xxx_messageInfo_Field.Size(m)
}
func (m *Field) XXX_DiscardUnknown() {
	xxx_messageInfo_Field.DiscardUnknown(m)
}

var xxx_messageInfo_Field proto.InternalMessageInfo

func (m *Field) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Field) GetValue() *Value {
	if m != nil {
		return m.Value
	}
	return nil
}

type Condition struct {
	Name                 string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                *Value       `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Operator             Operator     `protobuf:"varint,3,opt,name=operator,proto3,enum=gravity.api.querykit.Operator" json:"operator,omitempty"`
	Conditions           []*Condition `protobuf:"bytes,4,rep,name=conditions,proto3" json:"conditions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Condition) Reset()         { *m = Condition{} }
func (m *Condition) String() string { return proto.CompactTextString(m) }
func (*Condition) ProtoMessage()    {}
func (*Condition) Descriptor() ([]byte, []int) {
	return fileDescriptor_eda43376942b3faf, []int{2}
}

func (m *Condition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Condition.Unmarshal(m, b)
}
func (m *Condition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Condition.Marshal(b, m, deterministic)
}
func (m *Condition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Condition.Merge(m, src)
}
func (m *Condition) XXX_Size() int {
	return xxx_messageInfo_Condition.Size(m)
}
func (m *Condition) XXX_DiscardUnknown() {
	xxx_messageInfo_Condition.DiscardUnknown(m)
}

var xxx_messageInfo_Condition proto.InternalMessageInfo

func (m *Condition) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Condition) GetValue() *Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Condition) GetOperator() Operator {
	if m != nil {
		return m.Operator
	}
	return Operator_EQUAL
}

func (m *Condition) GetConditions() []*Condition {
	if m != nil {
		return m.Conditions
	}
	return nil
}

type Value struct {
	Type                 DataType    `protobuf:"varint,1,opt,name=type,proto3,enum=gravity.api.querykit.DataType" json:"type,omitempty"`
	Value                []byte      `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Map                  *MapValue   `protobuf:"bytes,3,opt,name=map,proto3" json:"map,omitempty"`
	Array                *ArrayValue `protobuf:"bytes,4,opt,name=array,proto3" json:"array,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Value) Reset()         { *m = Value{} }
func (m *Value) String() string { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()    {}
func (*Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_eda43376942b3faf, []int{3}
}

func (m *Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value.Unmarshal(m, b)
}
func (m *Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value.Marshal(b, m, deterministic)
}
func (m *Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value.Merge(m, src)
}
func (m *Value) XXX_Size() int {
	return xxx_messageInfo_Value.Size(m)
}
func (m *Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Value.DiscardUnknown(m)
}

var xxx_messageInfo_Value proto.InternalMessageInfo

func (m *Value) GetType() DataType {
	if m != nil {
		return m.Type
	}
	return DataType_BOOLEAN
}

func (m *Value) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Value) GetMap() *MapValue {
	if m != nil {
		return m.Map
	}
	return nil
}

func (m *Value) GetArray() *ArrayValue {
	if m != nil {
		return m.Array
	}
	return nil
}

type MapValue struct {
	Fields               []*Field `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MapValue) Reset()         { *m = MapValue{} }
func (m *MapValue) String() string { return proto.CompactTextString(m) }
func (*MapValue) ProtoMessage()    {}
func (*MapValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_eda43376942b3faf, []int{4}
}

func (m *MapValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MapValue.Unmarshal(m, b)
}
func (m *MapValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MapValue.Marshal(b, m, deterministic)
}
func (m *MapValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MapValue.Merge(m, src)
}
func (m *MapValue) XXX_Size() int {
	return xxx_messageInfo_MapValue.Size(m)
}
func (m *MapValue) XXX_DiscardUnknown() {
	xxx_messageInfo_MapValue.DiscardUnknown(m)
}

var xxx_messageInfo_MapValue proto.InternalMessageInfo

func (m *MapValue) GetFields() []*Field {
	if m != nil {
		return m.Fields
	}
	return nil
}

type ArrayValue struct {
	Elements             []*Value `protobuf:"bytes,1,rep,name=elements,proto3" json:"elements,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArrayValue) Reset()         { *m = ArrayValue{} }
func (m *ArrayValue) String() string { return proto.CompactTextString(m) }
func (*ArrayValue) ProtoMessage()    {}
func (*ArrayValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_eda43376942b3faf, []int{5}
}

func (m *ArrayValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArrayValue.Unmarshal(m, b)
}
func (m *ArrayValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArrayValue.Marshal(b, m, deterministic)
}
func (m *ArrayValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArrayValue.Merge(m, src)
}
func (m *ArrayValue) XXX_Size() int {
	return xxx_messageInfo_ArrayValue.Size(m)
}
func (m *ArrayValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ArrayValue.DiscardUnknown(m)
}

var xxx_messageInfo_ArrayValue proto.InternalMessageInfo

func (m *ArrayValue) GetElements() []*Value {
	if m != nil {
		return m.Elements
	}
	return nil
}

type QueryRequest struct {
	Table                string     `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	Condition            *Condition `protobuf:"bytes,2,opt,name=condition,proto3" json:"condition,omitempty"`
	Limit                int64      `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int64      `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	OrderBy              string     `protobuf:"bytes,5,opt,name=orderBy,proto3" json:"orderBy,omitempty"`
	Descending           bool       `protobuf:"varint,6,opt,name=descending,proto3" json:"descending,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eda43376942b3faf, []int{6}
}

func (m *QueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryRequest.Unmarshal(m, b)
}
func (m *QueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryRequest.Marshal(b, m, deterministic)
}
func (m *QueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequest.Merge(m, src)
}
func (m *QueryRequest) XXX_Size() int {
	return xxx_messageInfo_QueryRequest.Size(m)
}
func (m *QueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequest proto.InternalMessageInfo

func (m *QueryRequest) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *QueryRequest) GetCondition() *Condition {
	if m != nil {
		return m.Condition
	}
	return nil
}

func (m *QueryRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *QueryRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *QueryRequest) GetOrderBy() string {
	if m != nil {
		return m.OrderBy
	}
	return ""
}

func (m *QueryRequest) GetDescending() bool {
	if m != nil {
		return m.Descending
	}
	return false
}

type QueryReply struct {
	Success              bool      `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Reason               string    `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	Records              []*Record `protobuf:"bytes,3,rep,name=records,proto3" json:"records,omitempty"`
	TotalCount           int64     `protobuf:"varint,4,opt,name=totalCount,proto3" json:"totalCount,omitempty"`
	TotalPage            int64     `protobuf:"varint,5,opt,name=totalPage,proto3" json:"totalPage,omitempty"`
	CurrentPage          int64     `protobuf:"varint,6,opt,name=currentPage,proto3" json:"currentPage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *QueryReply) Reset()         { *m = QueryReply{} }
func (m *QueryReply) String() string { return proto.CompactTextString(m) }
func (*QueryReply) ProtoMessage()    {}
func (*QueryReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_eda43376942b3faf, []int{7}
}

func (m *QueryReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryReply.Unmarshal(m, b)
}
func (m *QueryReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryReply.Marshal(b, m, deterministic)
}
func (m *QueryReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryReply.Merge(m, src)
}
func (m *QueryReply) XXX_Size() int {
	return xxx_messageInfo_QueryReply.Size(m)
}
func (m *QueryReply) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryReply.DiscardUnknown(m)
}

var xxx_messageInfo_QueryReply proto.InternalMessageInfo

func (m *QueryReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *QueryReply) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *QueryReply) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

func (m *QueryReply) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *QueryReply) GetTotalPage() int64 {
	if m != nil {
		return m.TotalPage
	}
	return 0
}

func (m *QueryReply) GetCurrentPage() int64 {
	if m != nil {
		return m.CurrentPage
	}
	return 0
}

func init() {
	proto.RegisterEnum("gravity.api.querykit.DataType", DataType_name, DataType_value)
	proto.RegisterEnum("gravity.api.querykit.Operator", Operator_name, Operator_value)
	proto.RegisterType((*Record)(nil), "gravity.api.querykit.Record")
	proto.RegisterType((*Field)(nil), "gravity.api.querykit.Field")
	proto.RegisterType((*Condition)(nil), "gravity.api.querykit.Condition")
	proto.RegisterType((*Value)(nil), "gravity.api.querykit.Value")
	proto.RegisterType((*MapValue)(nil), "gravity.api.querykit.MapValue")
	proto.RegisterType((*ArrayValue)(nil), "gravity.api.querykit.ArrayValue")
	proto.RegisterType((*QueryRequest)(nil), "gravity.api.querykit.QueryRequest")
	proto.RegisterType((*QueryReply)(nil), "gravity.api.querykit.QueryReply")
}

func init() { proto.RegisterFile("querykit.proto", fileDescriptor_eda43376942b3faf) }

var fileDescriptor_eda43376942b3faf = []byte{
	// 682 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x41, 0x4f, 0xdb, 0x4a,
	0x10, 0xc7, 0x71, 0x1c, 0x3b, 0xce, 0x24, 0x20, 0xbf, 0x11, 0x7a, 0xb2, 0xde, 0x43, 0x3c, 0xcb,
	0xa7, 0x88, 0x43, 0xf4, 0x1a, 0x2a, 0x2a, 0x55, 0x42, 0xc8, 0x40, 0xa0, 0x51, 0x83, 0x03, 0x1b,
	0x53, 0x15, 0xf5, 0x80, 0x96, 0x64, 0x89, 0xac, 0x3a, 0xb6, 0xb1, 0x37, 0x48, 0xfe, 0x04, 0xfd,
	0x50, 0xfd, 0x08, 0x3d, 0xf4, 0xd6, 0xcf, 0x53, 0xed, 0xda, 0x0e, 0x39, 0x84, 0x22, 0x55, 0xbd,
	0xed, 0xcc, 0xfc, 0x66, 0x3c, 0xff, 0xd9, 0x59, 0xc3, 0xd6, 0xc3, 0x82, 0xa5, 0xf9, 0xe7, 0x80,
	0x77, 0x93, 0x34, 0xe6, 0x31, 0x6e, 0xcf, 0x52, 0xfa, 0x18, 0xf0, 0xbc, 0x4b, 0x93, 0xa0, 0x5b,
	0xc5, 0x9c, 0x43, 0xd0, 0x09, 0x9b, 0xc4, 0xe9, 0x14, 0xf7, 0x41, 0xbf, 0x0f, 0x58, 0x38, 0xcd,
	0x2c, 0xc5, 0x56, 0x3b, 0xad, 0xde, 0xbf, 0xdd, 0x75, 0x09, 0xdd, 0x33, 0xc1, 0x90, 0x12, 0x75,
	0x3c, 0xd0, 0xa4, 0x03, 0x11, 0xea, 0x11, 0x9d, 0x33, 0x4b, 0xb1, 0x95, 0x4e, 0x93, 0xc8, 0x33,
	0xbe, 0x02, 0xed, 0x91, 0x86, 0x0b, 0x66, 0xd5, 0x6c, 0xe5, 0xf9, 0x82, 0x1f, 0x04, 0x42, 0x0a,
	0xd2, 0xf9, 0xae, 0x40, 0xf3, 0x24, 0x8e, 0xa6, 0x01, 0x0f, 0xe2, 0xe8, 0x0f, 0x15, 0xc5, 0xb7,
	0x60, 0xc4, 0x09, 0x4b, 0x29, 0x8f, 0x53, 0x4b, 0xb5, 0x95, 0xce, 0x56, 0x6f, 0x77, 0x7d, 0xd6,
	0xa8, 0xa4, 0xc8, 0x92, 0xc7, 0x23, 0x80, 0x49, 0xd5, 0x4f, 0x66, 0xd5, 0xe5, 0x64, 0xfe, 0x5b,
	0x9f, 0xbd, 0xec, 0x9b, 0xac, 0xa4, 0x38, 0x5f, 0x15, 0xd0, 0x64, 0x37, 0xd8, 0x83, 0x3a, 0xcf,
	0x93, 0x42, 0xcd, 0xb3, 0x2d, 0x9c, 0x52, 0x4e, 0xfd, 0x3c, 0x61, 0x44, 0xb2, 0xb8, 0xbd, 0xaa,
	0xb6, 0x5d, 0x09, 0xfa, 0x1f, 0xd4, 0x39, 0x4d, 0xa4, 0x96, 0xd6, 0x73, 0x85, 0x2e, 0x68, 0x52,
	0x0c, 0x41, 0xa0, 0x78, 0x00, 0x1a, 0x4d, 0x53, 0x9a, 0x5b, 0x75, 0x99, 0x63, 0xaf, 0xcf, 0x71,
	0x05, 0x52, 0x8e, 0x4e, 0xe2, 0xce, 0x11, 0x18, 0x55, 0xa1, 0xdf, 0x5b, 0x90, 0x3e, 0xc0, 0x53,
	0x55, 0x7c, 0x03, 0x06, 0x0b, 0xd9, 0x9c, 0x45, 0xfc, 0x85, 0x22, 0x45, 0x13, 0x4b, 0xd8, 0xf9,
	0xa6, 0x40, 0xfb, 0x4a, 0x04, 0x09, 0x7b, 0x58, 0xb0, 0x8c, 0x8b, 0xc1, 0x70, 0x7a, 0x17, 0x56,
	0xbb, 0x51, 0x18, 0x78, 0x08, 0xcd, 0xe5, 0xe8, 0xcb, 0x05, 0x79, 0xf1, 0xb2, 0x9e, 0x32, 0x44,
	0xd1, 0x30, 0x98, 0x07, 0x5c, 0x4e, 0x56, 0x25, 0x85, 0x81, 0x7f, 0x83, 0x1e, 0xdf, 0xdf, 0x67,
	0x8c, 0xcb, 0xe1, 0xa9, 0xa4, 0xb4, 0xd0, 0x82, 0x46, 0x9c, 0x4e, 0x59, 0x7a, 0x9c, 0x5b, 0x9a,
	0x6c, 0xa2, 0x32, 0x71, 0x17, 0x60, 0xca, 0xb2, 0x09, 0x8b, 0xa6, 0x41, 0x34, 0xb3, 0x74, 0x5b,
	0xe9, 0x18, 0x64, 0xc5, 0xe3, 0xfc, 0x50, 0x00, 0x4a, 0x35, 0x49, 0x98, 0x8b, 0x42, 0xd9, 0x62,
	0x32, 0x61, 0x59, 0x26, 0xd5, 0x18, 0xa4, 0x32, 0xc5, 0xa7, 0x53, 0x46, 0xb3, 0x52, 0x4c, 0x93,
	0x94, 0x16, 0x1e, 0x40, 0x23, 0x95, 0xaf, 0x36, 0xb3, 0x54, 0x39, 0xc6, 0x9d, 0xf5, 0x2a, 0x8b,
	0xa7, 0x4d, 0x2a, 0x58, 0x34, 0xc6, 0x63, 0x4e, 0xc3, 0x93, 0x78, 0x11, 0x55, 0x72, 0x56, 0x3c,
	0xb8, 0x03, 0x4d, 0x69, 0x5d, 0xd2, 0x19, 0x93, 0xa2, 0x54, 0xf2, 0xe4, 0x40, 0x1b, 0x5a, 0x93,
	0x45, 0x9a, 0xb2, 0x88, 0xcb, 0xb8, 0x2e, 0xe3, 0xab, 0xae, 0xbd, 0x19, 0x18, 0xd5, 0x02, 0x63,
	0x0b, 0x1a, 0xc7, 0xa3, 0xd1, 0xb0, 0xef, 0x7a, 0xe6, 0x06, 0x02, 0xe8, 0xc7, 0x03, 0xcf, 0x25,
	0x37, 0xa6, 0x22, 0xce, 0x63, 0x9f, 0x0c, 0xbc, 0x73, 0xb3, 0x26, 0xce, 0xd7, 0x03, 0xcf, 0x3f,
	0x78, 0x6d, 0xaa, 0xd8, 0x04, 0xad, 0x38, 0xd6, 0x45, 0xee, 0xd9, 0x70, 0xe4, 0x0a, 0x43, 0x13,
	0x7e, 0x97, 0x10, 0xf7, 0xc6, 0xd4, 0xb1, 0x01, 0xea, 0x85, 0x7b, 0x69, 0x36, 0xf6, 0xbe, 0x28,
	0x60, 0x54, 0xaf, 0x55, 0x00, 0xfd, 0xab, 0x6b, 0x77, 0x68, 0x6e, 0xa0, 0x09, 0xed, 0x73, 0xd2,
	0x77, 0xfd, 0x3e, 0xb9, 0xf5, 0xdf, 0xb9, 0x9e, 0xa9, 0xe0, 0x5f, 0xb0, 0x59, 0x79, 0x0a, 0xa8,
	0x86, 0x9b, 0xd0, 0x1c, 0xf6, 0xc7, 0xe3, 0x82, 0x50, 0x71, 0x0b, 0x40, 0x9a, 0x45, 0xb8, 0x2e,
	0x3e, 0xe2, 0x7a, 0xa7, 0xa6, 0x86, 0x3a, 0xd4, 0x46, 0xc4, 0xd4, 0x05, 0xef, 0x8d, 0xfc, 0x32,
	0xde, 0xc0, 0x36, 0x18, 0x83, 0xf1, 0x6d, 0xff, 0xe3, 0x60, 0xec, 0x9b, 0x46, 0xef, 0x13, 0x18,
	0xf2, 0x2a, 0xdf, 0x07, 0x1c, 0x47, 0xa0, 0xc9, 0x33, 0x3a, 0xeb, 0xaf, 0x63, 0x75, 0x83, 0xff,
	0xb1, 0x7f, 0xc9, 0x24, 0x61, 0xee, 0x6c, 0xdc, 0xe9, 0xf2, 0xd7, 0xbd, 0xff, 0x33, 0x00, 0x00,
	0xff, 0xff, 0x4a, 0x52, 0xbd, 0xcb, 0xcc, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryKitClient is the client API for QueryKit service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryKitClient interface {
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryReply, error)
}

type queryKitClient struct {
	cc *grpc.ClientConn
}

func NewQueryKitClient(cc *grpc.ClientConn) QueryKitClient {
	return &queryKitClient{cc}
}

func (c *queryKitClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryReply, error) {
	out := new(QueryReply)
	err := c.cc.Invoke(ctx, "/gravity.api.querykit.QueryKit/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryKitServer is the server API for QueryKit service.
type QueryKitServer interface {
	Query(context.Context, *QueryRequest) (*QueryReply, error)
}

// UnimplementedQueryKitServer can be embedded to have forward compatible implementations.
type UnimplementedQueryKitServer struct {
}

func (*UnimplementedQueryKitServer) Query(ctx context.Context, req *QueryRequest) (*QueryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}

func RegisterQueryKitServer(s *grpc.Server, srv QueryKitServer) {
	s.RegisterService(&_QueryKit_serviceDesc, srv)
}

func _QueryKit_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryKitServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gravity.api.querykit.QueryKit/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryKitServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QueryKit_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gravity.api.querykit.QueryKit",
	HandlerType: (*QueryKitServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _QueryKit_Query_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "querykit.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/interservice/cfgmgmt/request/parameters.proto

package request // import "github.com/chef/automate/api/interservice/cfgmgmt/request"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Order int32

const (
	Order_asc  Order = 0
	Order_desc Order = 1
)

var Order_name = map[int32]string{
	0: "asc",
	1: "desc",
}
var Order_value = map[string]int32{
	"asc":  0,
	"desc": 1,
}

func (x Order) String() string {
	return proto.EnumName(Order_name, int32(x))
}
func (Order) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_parameters_acefa4f27c4cfe02, []int{0}
}

type Pagination struct {
	Page                 int32    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty" toml:"page,omitempty" mapstructure:"page,omitempty"`
	Size                 int32    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty" toml:"size,omitempty" mapstructure:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *Pagination) Reset()         { *m = Pagination{} }
func (m *Pagination) String() string { return proto.CompactTextString(m) }
func (*Pagination) ProtoMessage()    {}
func (*Pagination) Descriptor() ([]byte, []int) {
	return fileDescriptor_parameters_acefa4f27c4cfe02, []int{0}
}
func (m *Pagination) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pagination.Unmarshal(m, b)
}
func (m *Pagination) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pagination.Marshal(b, m, deterministic)
}
func (dst *Pagination) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pagination.Merge(dst, src)
}
func (m *Pagination) XXX_Size() int {
	return xxx_messageInfo_Pagination.Size(m)
}
func (m *Pagination) XXX_DiscardUnknown() {
	xxx_messageInfo_Pagination.DiscardUnknown(m)
}

var xxx_messageInfo_Pagination proto.InternalMessageInfo

func (m *Pagination) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *Pagination) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type Sorting struct {
	Field                string   `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty" toml:"field,omitempty" mapstructure:"field,omitempty"`
	Order                Order    `protobuf:"varint,2,opt,name=order,proto3,enum=chef.automate.domain.cfgmgmt.request.Order" json:"order,omitempty" toml:"order,omitempty" mapstructure:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *Sorting) Reset()         { *m = Sorting{} }
func (m *Sorting) String() string { return proto.CompactTextString(m) }
func (*Sorting) ProtoMessage()    {}
func (*Sorting) Descriptor() ([]byte, []int) {
	return fileDescriptor_parameters_acefa4f27c4cfe02, []int{1}
}
func (m *Sorting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sorting.Unmarshal(m, b)
}
func (m *Sorting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sorting.Marshal(b, m, deterministic)
}
func (dst *Sorting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sorting.Merge(dst, src)
}
func (m *Sorting) XXX_Size() int {
	return xxx_messageInfo_Sorting.Size(m)
}
func (m *Sorting) XXX_DiscardUnknown() {
	xxx_messageInfo_Sorting.DiscardUnknown(m)
}

var xxx_messageInfo_Sorting proto.InternalMessageInfo

func (m *Sorting) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *Sorting) GetOrder() Order {
	if m != nil {
		return m.Order
	}
	return Order_asc
}

type Suggestion struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty" toml:"type,omitempty" mapstructure:"type,omitempty"`
	Text                 string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty" toml:"text,omitempty" mapstructure:"text,omitempty"`
	Filter               []string `protobuf:"bytes,3,rep,name=filter,proto3" json:"filter,omitempty" toml:"filter,omitempty" mapstructure:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *Suggestion) Reset()         { *m = Suggestion{} }
func (m *Suggestion) String() string { return proto.CompactTextString(m) }
func (*Suggestion) ProtoMessage()    {}
func (*Suggestion) Descriptor() ([]byte, []int) {
	return fileDescriptor_parameters_acefa4f27c4cfe02, []int{2}
}
func (m *Suggestion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Suggestion.Unmarshal(m, b)
}
func (m *Suggestion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Suggestion.Marshal(b, m, deterministic)
}
func (dst *Suggestion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Suggestion.Merge(dst, src)
}
func (m *Suggestion) XXX_Size() int {
	return xxx_messageInfo_Suggestion.Size(m)
}
func (m *Suggestion) XXX_DiscardUnknown() {
	xxx_messageInfo_Suggestion.DiscardUnknown(m)
}

var xxx_messageInfo_Suggestion proto.InternalMessageInfo

func (m *Suggestion) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Suggestion) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Suggestion) GetFilter() []string {
	if m != nil {
		return m.Filter
	}
	return nil
}

func init() {
	proto.RegisterType((*Pagination)(nil), "chef.automate.domain.cfgmgmt.request.Pagination")
	proto.RegisterType((*Sorting)(nil), "chef.automate.domain.cfgmgmt.request.Sorting")
	proto.RegisterType((*Suggestion)(nil), "chef.automate.domain.cfgmgmt.request.Suggestion")
	proto.RegisterEnum("chef.automate.domain.cfgmgmt.request.Order", Order_name, Order_value)
}

func init() {
	proto.RegisterFile("api/interservice/cfgmgmt/request/parameters.proto", fileDescriptor_parameters_acefa4f27c4cfe02)
}

var fileDescriptor_parameters_acefa4f27c4cfe02 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xcf, 0x4b, 0x33, 0x31,
	0x10, 0x86, 0xbf, 0x7e, 0xed, 0xb6, 0x76, 0x0e, 0x52, 0x82, 0xc8, 0xe2, 0xa9, 0x2c, 0x1e, 0x8a,
	0x42, 0x82, 0x3f, 0x2e, 0xe2, 0x49, 0xcf, 0x82, 0xb2, 0xbd, 0x79, 0xcb, 0x66, 0x67, 0xd3, 0x40,
	0xb3, 0x59, 0x93, 0x59, 0xa9, 0xfe, 0xf5, 0x92, 0xec, 0xf6, 0xac, 0xb7, 0x77, 0x5e, 0xf2, 0x64,
	0x1e, 0x06, 0x6e, 0x64, 0x67, 0x84, 0x69, 0x09, 0x7d, 0x40, 0xff, 0x69, 0x14, 0x0a, 0xd5, 0x68,
	0xab, 0x2d, 0x09, 0x8f, 0x1f, 0x3d, 0x06, 0x12, 0x9d, 0xf4, 0xd2, 0x62, 0x7c, 0xc0, 0x3b, 0xef,
	0xc8, 0xb1, 0x4b, 0xb5, 0xc3, 0x86, 0xcb, 0x9e, 0x9c, 0x95, 0x84, 0xbc, 0x76, 0x56, 0x9a, 0x96,
	0x8f, 0x18, 0x1f, 0xb1, 0xe2, 0x1e, 0xe0, 0x4d, 0x6a, 0xd3, 0x4a, 0x32, 0xae, 0x65, 0x0c, 0x66,
	0x9d, 0xd4, 0x98, 0x4f, 0xd6, 0x93, 0x4d, 0x56, 0xa6, 0x1c, 0xbb, 0x60, 0xbe, 0x31, 0xff, 0x3f,
	0x74, 0x31, 0x17, 0x15, 0x2c, 0xb6, 0xce, 0x93, 0x69, 0x35, 0x3b, 0x83, 0xac, 0x31, 0xb8, 0xaf,
	0x13, 0xb3, 0x2c, 0x87, 0x81, 0x3d, 0x41, 0xe6, 0x7c, 0x8d, 0x3e, 0x51, 0xa7, 0xb7, 0xd7, 0xfc,
	0x2f, 0x32, 0xfc, 0x35, 0x22, 0xe5, 0x40, 0x16, 0x2f, 0x00, 0xdb, 0x5e, 0x6b, 0x0c, 0x47, 0x33,
	0xfa, 0xea, 0x70, 0xdc, 0x92, 0x72, 0xea, 0xf0, 0x40, 0x69, 0x47, 0xec, 0xf0, 0x40, 0xec, 0x1c,
	0xe6, 0x8d, 0xd9, 0x13, 0xfa, 0x7c, 0xba, 0x9e, 0x6e, 0x96, 0xe5, 0x38, 0x5d, 0x5d, 0x40, 0x96,
	0x7e, 0x67, 0x0b, 0x98, 0xca, 0xa0, 0x56, 0xff, 0xd8, 0x09, 0xcc, 0x6a, 0x0c, 0x6a, 0x35, 0x79,
	0x7e, 0x7c, 0x7f, 0xd0, 0x86, 0x76, 0x7d, 0xc5, 0x95, 0xb3, 0x22, 0x9a, 0x8a, 0xa3, 0xa9, 0xf8,
	0xed, 0xee, 0xd5, 0x3c, 0x5d, 0xfb, 0xee, 0x27, 0x00, 0x00, 0xff, 0xff, 0x69, 0xa6, 0x47, 0x71,
	0xa2, 0x01, 0x00, 0x00,
}

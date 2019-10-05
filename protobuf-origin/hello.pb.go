// Code generated by protoc-gen-go. DO NOT EDIT.
// source: src/protobuf-origin/hello.proto

package awesomepackage

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type AwesomeMessage struct {
	AwesomeField         string   `protobuf:"bytes,1,opt,name=awesomeField,proto3" json:"awesomeField,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AwesomeMessage) Reset()         { *m = AwesomeMessage{} }
func (m *AwesomeMessage) String() string { return proto.CompactTextString(m) }
func (*AwesomeMessage) ProtoMessage()    {}
func (*AwesomeMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_f11e4b3b8bab0012, []int{0}
}

func (m *AwesomeMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AwesomeMessage.Unmarshal(m, b)
}
func (m *AwesomeMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AwesomeMessage.Marshal(b, m, deterministic)
}
func (m *AwesomeMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AwesomeMessage.Merge(m, src)
}
func (m *AwesomeMessage) XXX_Size() int {
	return xxx_messageInfo_AwesomeMessage.Size(m)
}
func (m *AwesomeMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_AwesomeMessage.DiscardUnknown(m)
}

var xxx_messageInfo_AwesomeMessage proto.InternalMessageInfo

func (m *AwesomeMessage) GetAwesomeField() string {
	if m != nil {
		return m.AwesomeField
	}
	return ""
}

func init() {
	proto.RegisterType((*AwesomeMessage)(nil), "awesomepackage.AwesomeMessage")
}

func init() { proto.RegisterFile("src/protobuf-origin/hello.proto", fileDescriptor_f11e4b3b8bab0012) }

var fileDescriptor_f11e4b3b8bab0012 = []byte{
	// 109 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2f, 0x2e, 0x4a, 0xd6,
	0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0xcd, 0x2f, 0xca, 0x4c, 0xcf, 0xcc, 0xd3,
	0xcf, 0x48, 0xcd, 0xc9, 0xc9, 0xd7, 0x03, 0x8b, 0x0a, 0xf1, 0x25, 0x96, 0xa7, 0x16, 0xe7, 0xe7,
	0xa6, 0x16, 0x24, 0x26, 0x67, 0x27, 0xa6, 0xa7, 0x2a, 0x99, 0x70, 0xf1, 0x39, 0x42, 0x44, 0x7c,
	0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x85, 0x94, 0xb8, 0x78, 0xa0, 0x6a, 0xdc, 0x32, 0x53, 0x73,
	0x52, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x50, 0xc4, 0x92, 0xd8, 0xc0, 0x86, 0x19, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x2e, 0xd5, 0x7e, 0x81, 0x6f, 0x00, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/hello.proto

package hello

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

type Req struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Req) Reset()         { *m = Req{} }
func (m *Req) String() string { return proto.CompactTextString(m) }
func (*Req) ProtoMessage()    {}
func (*Req) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9418a89e4940c19, []int{0}
}

func (m *Req) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Req.Unmarshal(m, b)
}
func (m *Req) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Req.Marshal(b, m, deterministic)
}
func (m *Req) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Req.Merge(m, src)
}
func (m *Req) XXX_Size() int {
	return xxx_messageInfo_Req.Size(m)
}
func (m *Req) XXX_DiscardUnknown() {
	xxx_messageInfo_Req.DiscardUnknown(m)
}

var xxx_messageInfo_Req proto.InternalMessageInfo

func (m *Req) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Res struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Res) Reset()         { *m = Res{} }
func (m *Res) String() string { return proto.CompactTextString(m) }
func (*Res) ProtoMessage()    {}
func (*Res) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9418a89e4940c19, []int{1}
}

func (m *Res) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Res.Unmarshal(m, b)
}
func (m *Res) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Res.Marshal(b, m, deterministic)
}
func (m *Res) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Res.Merge(m, src)
}
func (m *Res) XXX_Size() int {
	return xxx_messageInfo_Res.Size(m)
}
func (m *Res) XXX_DiscardUnknown() {
	xxx_messageInfo_Res.DiscardUnknown(m)
}

var xxx_messageInfo_Res proto.InternalMessageInfo

func (m *Res) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Req)(nil), "hello_grpc.Req")
	proto.RegisterType((*Res)(nil), "hello_grpc.Res")
}

func init() { proto.RegisterFile("pb/hello.proto", fileDescriptor_b9418a89e4940c19) }

var fileDescriptor_b9418a89e4940c19 = []byte{
	// 128 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x48, 0xd2, 0xcf,
	0x48, 0xcd, 0xc9, 0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x02, 0x73, 0xe2, 0xd3,
	0x8b, 0x0a, 0x92, 0x95, 0xe4, 0xb9, 0x98, 0x83, 0x52, 0x0b, 0x85, 0x24, 0xb8, 0xd8, 0x73, 0x53,
	0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0x88, 0x82,
	0x62, 0xdc, 0x0a, 0x8c, 0xcc, 0xb8, 0x38, 0x3d, 0x40, 0xe6, 0xb9, 0x17, 0x15, 0x24, 0x0b, 0x69,
	0x72, 0xb1, 0x06, 0x27, 0x56, 0x7a, 0x64, 0x0a, 0xf1, 0xeb, 0x21, 0x2c, 0xd1, 0x0b, 0x4a, 0x2d,
	0x94, 0x42, 0x13, 0x28, 0x76, 0xe2, 0x8a, 0xe2, 0xd0, 0xd3, 0xb7, 0x06, 0x0b, 0x26, 0xb1, 0x81,
	0x1d, 0x66, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x29, 0x30, 0x90, 0x91, 0xaa, 0x00, 0x00, 0x00,
}
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/common.proto

package common

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

type BaseReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BaseReq) Reset()         { *m = BaseReq{} }
func (m *BaseReq) String() string { return proto.CompactTextString(m) }
func (*BaseReq) ProtoMessage()    {}
func (*BaseReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0}
}

func (m *BaseReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseReq.Unmarshal(m, b)
}
func (m *BaseReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseReq.Marshal(b, m, deterministic)
}
func (m *BaseReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseReq.Merge(m, src)
}
func (m *BaseReq) XXX_Size() int {
	return xxx_messageInfo_BaseReq.Size(m)
}
func (m *BaseReq) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseReq.DiscardUnknown(m)
}

var xxx_messageInfo_BaseReq proto.InternalMessageInfo

func (m *BaseReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*BaseReq)(nil), "pb.BaseReq")
}

func init() {
	proto.RegisterFile("common/common.proto", fileDescriptor_8f954d82c0b891f6)
}

var fileDescriptor_8f954d82c0b891f6 = []byte{
	// 102 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x87, 0x50, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x4c, 0x05, 0x49, 0x4a,
	0xf2, 0x5c, 0xec, 0x4e, 0x89, 0xc5, 0xa9, 0x41, 0xa9, 0x85, 0x42, 0x22, 0x5c, 0xac, 0x25, 0xf9,
	0xd9, 0xa9, 0x79, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x8e, 0x93, 0x7c, 0x94, 0x6c,
	0x46, 0x6a, 0x4e, 0x4e, 0xbe, 0x7e, 0x41, 0x92, 0x7e, 0x7a, 0x6a, 0x9e, 0x6e, 0x7a, 0x3e, 0xd4,
	0x10, 0x6b, 0x08, 0x95, 0xc4, 0x06, 0x36, 0xcc, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x4f,
	0x36, 0x14, 0x63, 0x00, 0x00, 0x00,
}

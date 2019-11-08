// Code generated by protoc-gen-go. DO NOT EDIT.
// source: com/duan/session/session.proto

package com_duan_session

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

type SessionStatus int32

const (
	SessionStatus_SESSION_OPEN   SessionStatus = 0
	SessionStatus_SESSION_CLOSED SessionStatus = 1
)

var SessionStatus_name = map[int32]string{
	0: "SESSION_OPEN",
	1: "SESSION_CLOSED",
}

var SessionStatus_value = map[string]int32{
	"SESSION_OPEN":   0,
	"SESSION_CLOSED": 1,
}

func (x SessionStatus) String() string {
	return proto.EnumName(SessionStatus_name, int32(x))
}

func (SessionStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c02f53dda92c45ee, []int{0}
}

type Session struct {
	Id                   string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 int32         `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Status               SessionStatus `protobuf:"varint,3,opt,name=status,proto3,enum=com.duan.session.SessionStatus" json:"status,omitempty"`
	Topic                string        `protobuf:"bytes,4,opt,name=topic,proto3" json:"topic,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_c02f53dda92c45ee, []int{0}
}

func (m *Session) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Session.Unmarshal(m, b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Session.Marshal(b, m, deterministic)
}
func (m *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(m, src)
}
func (m *Session) XXX_Size() int {
	return xxx_messageInfo_Session.Size(m)
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Session) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Session) GetStatus() SessionStatus {
	if m != nil {
		return m.Status
	}
	return SessionStatus_SESSION_OPEN
}

func (m *Session) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func init() {
	proto.RegisterEnum("com.duan.session.SessionStatus", SessionStatus_name, SessionStatus_value)
	proto.RegisterType((*Session)(nil), "com.duan.session.Session")
}

func init() { proto.RegisterFile("com/duan/session/session.proto", fileDescriptor_c02f53dda92c45ee) }

var fileDescriptor_c02f53dda92c45ee = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xce, 0xcf, 0xd5,
	0x4f, 0x29, 0x4d, 0xcc, 0xd3, 0x2f, 0x4e, 0x2d, 0x2e, 0xce, 0xcc, 0x87, 0xd3, 0x7a, 0x05, 0x45,
	0xf9, 0x25, 0xf9, 0x42, 0x02, 0xc9, 0xf9, 0xb9, 0x7a, 0x20, 0x79, 0x3d, 0xa8, 0xb8, 0x52, 0x0d,
	0x17, 0x7b, 0x30, 0x84, 0x29, 0xc4, 0xc7, 0xc5, 0x94, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1,
	0x19, 0xc4, 0x94, 0x99, 0x22, 0x24, 0xc4, 0xc5, 0x52, 0x52, 0x59, 0x90, 0x2a, 0xc1, 0xa4, 0xc0,
	0xa8, 0xc1, 0x1a, 0x04, 0x66, 0x0b, 0x99, 0x73, 0xb1, 0x15, 0x97, 0x24, 0x96, 0x94, 0x16, 0x4b,
	0x30, 0x2b, 0x30, 0x6a, 0xf0, 0x19, 0xc9, 0xeb, 0xa1, 0x9b, 0xa8, 0x07, 0x35, 0x2e, 0x18, 0xac,
	0x2c, 0x08, 0xaa, 0x5c, 0x48, 0x84, 0x8b, 0xb5, 0x24, 0xbf, 0x20, 0x33, 0x59, 0x82, 0x05, 0x6c,
	0x3e, 0x84, 0xa3, 0x65, 0xca, 0xc5, 0x8b, 0xa2, 0x5c, 0x48, 0x80, 0x8b, 0x27, 0xd8, 0x35, 0x38,
	0xd8, 0xd3, 0xdf, 0x2f, 0xde, 0x3f, 0xc0, 0xd5, 0x4f, 0x80, 0x41, 0x48, 0x88, 0x8b, 0x0f, 0x26,
	0xe2, 0xec, 0xe3, 0x1f, 0xec, 0xea, 0x22, 0xc0, 0x98, 0xc4, 0x06, 0xf6, 0x8d, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0x1d, 0x8e, 0xb6, 0x24, 0xef, 0x00, 0x00, 0x00,
}

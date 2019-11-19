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
	SessionStatus_DONT_USE_0_IN_SESSIONSTATUS SessionStatus = 0
	SessionStatus_NEWBORN                     SessionStatus = 1
	SessionStatus_SESSION_CLOSED              SessionStatus = 2
	SessionStatus_SESSION_OPEN                SessionStatus = 3
)

var SessionStatus_name = map[int32]string{
	0: "DONT_USE_0_IN_SESSIONSTATUS",
	1: "NEWBORN",
	2: "SESSION_CLOSED",
	3: "SESSION_OPEN",
}

var SessionStatus_value = map[string]int32{
	"DONT_USE_0_IN_SESSIONSTATUS": 0,
	"NEWBORN":                     1,
	"SESSION_CLOSED":              2,
	"SESSION_OPEN":                3,
}

func (x SessionStatus) String() string {
	return proto.EnumName(SessionStatus_name, int32(x))
}

func (SessionStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c02f53dda92c45ee, []int{0}
}

type SessionType int32

const (
	SessionType_DONT_USE_0_IN_SESSIONTYPE SessionType = 0
	SessionType_LONG_TERM                 SessionType = 1
)

var SessionType_name = map[int32]string{
	0: "DONT_USE_0_IN_SESSIONTYPE",
	1: "LONG_TERM",
}

var SessionType_value = map[string]int32{
	"DONT_USE_0_IN_SESSIONTYPE": 0,
	"LONG_TERM":                 1,
}

func (x SessionType) String() string {
	return proto.EnumName(SessionType_name, int32(x))
}

func (SessionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c02f53dda92c45ee, []int{1}
}

type Session struct {
	Id                   int64         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 SessionType   `protobuf:"varint,2,opt,name=type,proto3,enum=com.duan.session.SessionType" json:"type,omitempty"`
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

func (m *Session) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Session) GetType() SessionType {
	if m != nil {
		return m.Type
	}
	return SessionType_DONT_USE_0_IN_SESSIONTYPE
}

func (m *Session) GetStatus() SessionStatus {
	if m != nil {
		return m.Status
	}
	return SessionStatus_DONT_USE_0_IN_SESSIONSTATUS
}

func (m *Session) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func init() {
	proto.RegisterEnum("com.duan.session.SessionStatus", SessionStatus_name, SessionStatus_value)
	proto.RegisterEnum("com.duan.session.SessionType", SessionType_name, SessionType_value)
	proto.RegisterType((*Session)(nil), "com.duan.session.Session")
}

func init() { proto.RegisterFile("com/duan/session/session.proto", fileDescriptor_c02f53dda92c45ee) }

var fileDescriptor_c02f53dda92c45ee = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcf, 0x4f, 0xc2, 0x30,
	0x14, 0xc7, 0xe9, 0x86, 0x10, 0x1e, 0xb2, 0x34, 0x2f, 0x1e, 0x66, 0x0c, 0xb2, 0x78, 0x5a, 0x38,
	0x0c, 0x7f, 0x1c, 0x3c, 0x78, 0x52, 0x69, 0x0c, 0x09, 0xb6, 0xa4, 0x2d, 0x31, 0x9e, 0x1a, 0xdc,
	0x76, 0xd8, 0x01, 0xbb, 0xb8, 0x72, 0xe0, 0x6f, 0xf1, 0x9f, 0x35, 0x8e, 0x69, 0xd4, 0xc8, 0xa9,
	0xed, 0xf7, 0xfb, 0xf9, 0xa4, 0x79, 0x0f, 0x4e, 0x53, 0xbb, 0x9e, 0x64, 0x9b, 0xd5, 0xeb, 0xa4,
	0xca, 0xab, 0xaa, 0xb0, 0xdf, 0x67, 0x52, 0xbe, 0x59, 0x67, 0x91, 0xa6, 0x76, 0x9d, 0x7c, 0xf6,
	0x49, 0x93, 0x9f, 0xbd, 0x13, 0xe8, 0xaa, 0xdd, 0x1d, 0x03, 0xf0, 0x8a, 0x2c, 0x24, 0x11, 0x89,
	0x7d, 0xe9, 0x15, 0x19, 0x5e, 0x40, 0xdb, 0x6d, 0xcb, 0x3c, 0xf4, 0x22, 0x12, 0x07, 0x97, 0xc3,
	0xe4, 0xaf, 0x9c, 0x34, 0xa2, 0xde, 0x96, 0xb9, 0xac, 0x51, 0xbc, 0x86, 0x4e, 0xe5, 0x56, 0x6e,
	0x53, 0x85, 0x7e, 0x2d, 0x8d, 0xf6, 0x4a, 0xaa, 0xc6, 0x64, 0x83, 0xe3, 0x11, 0x1c, 0x38, 0x5b,
	0x16, 0x69, 0xd8, 0x8e, 0x48, 0xdc, 0x93, 0xbb, 0xc7, 0x38, 0x85, 0xc1, 0x2f, 0x1c, 0x47, 0x70,
	0x32, 0x15, 0x5c, 0x9b, 0xa5, 0x62, 0xe6, 0xdc, 0xcc, 0xb8, 0x51, 0x4c, 0xa9, 0x99, 0xe0, 0x4a,
	0xdf, 0xea, 0xa5, 0xa2, 0x2d, 0xec, 0x43, 0x97, 0xb3, 0xa7, 0x3b, 0x21, 0x39, 0x25, 0x88, 0x10,
	0x34, 0xbd, 0xb9, 0x9f, 0x0b, 0xc5, 0xa6, 0xd4, 0x43, 0x0a, 0x87, 0x5f, 0x99, 0x58, 0x30, 0x4e,
	0xfd, 0xf1, 0x0d, 0xf4, 0x7f, 0x0c, 0x82, 0x43, 0x38, 0xfe, 0xf7, 0x0b, 0xfd, 0xbc, 0x60, 0xb4,
	0x85, 0x03, 0xe8, 0xcd, 0x05, 0x7f, 0x30, 0x9a, 0xc9, 0x47, 0x4a, 0x5e, 0x3a, 0xf5, 0x62, 0xaf,
	0x3e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x2b, 0xb0, 0xc5, 0xc1, 0x7a, 0x01, 0x00, 0x00,
}

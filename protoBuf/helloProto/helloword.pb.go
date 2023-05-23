// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloword.proto

package pb

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

// 请求消息
type HelloRequest struct {
	// 姓名
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// 密码
	Password             int32    `protobuf:"varint,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_408dc8e96988f433, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HelloRequest) GetPassword() int32 {
	if m != nil {
		return m.Password
	}
	return 0
}

// 响应消息
type HelloResponse struct {
	// 错误码
	Error int32 `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	// 消息
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_408dc8e96988f433, []int{1}
}

func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (m *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(m, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetError() int32 {
	if m != nil {
		return m.Error
	}
	return 0
}

func (m *HelloResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "helloword.helloRequest")
	proto.RegisterType((*HelloResponse)(nil), "helloword.helloResponse")
}

func init() { proto.RegisterFile("helloword.proto", fileDescriptor_408dc8e96988f433) }

var fileDescriptor_408dc8e96988f433 = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0x48, 0xcd, 0xc9,
	0xc9, 0x2f, 0xcf, 0x2f, 0x4a, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x0b, 0x28,
	0xd9, 0x71, 0xf1, 0x80, 0x39, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42, 0x5c, 0x2c,
	0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x90, 0x14, 0x17,
	0x47, 0x41, 0x62, 0x71, 0x31, 0x48, 0xbd, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x9c, 0xaf,
	0x64, 0xcf, 0xc5, 0x0b, 0xd5, 0x5f, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x24, 0xc2, 0xc5, 0x9a,
	0x5a, 0x54, 0x94, 0x5f, 0x04, 0x36, 0x81, 0x35, 0x08, 0xc2, 0x11, 0x92, 0xe0, 0x62, 0xcf, 0x4d,
	0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x05, 0x9b, 0xc0, 0x19, 0x04, 0xe3, 0x1a, 0x05, 0x72, 0x09, 0xc0,
	0x5d, 0x13, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0x64, 0xcb, 0xc5, 0x11, 0x9c, 0x58, 0xe9,
	0x01, 0x12, 0x16, 0x12, 0xd7, 0x43, 0xb8, 0x1e, 0xd9, 0xa5, 0x52, 0x12, 0x98, 0x12, 0x10, 0x27,
	0x38, 0xb1, 0x47, 0xb1, 0xea, 0xe9, 0x5b, 0x17, 0x24, 0x25, 0xb1, 0x81, 0xbd, 0x6b, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x24, 0x94, 0x0d, 0x83, 0x01, 0x01, 0x00, 0x00,
}

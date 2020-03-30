// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package go_micro_api

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

type Endpoint struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Host                 []string `protobuf:"bytes,2,rep,name=host,proto3" json:"host,omitempty"`
	Path                 []string `protobuf:"bytes,3,rep,name=path,proto3" json:"path,omitempty"`
	Method               []string `protobuf:"bytes,4,rep,name=method,proto3" json:"method,omitempty"`
	Stream               bool     `protobuf:"varint,5,opt,name=stream,proto3" json:"stream,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Endpoint) Reset()         { *m = Endpoint{} }
func (m *Endpoint) String() string { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()    {}
func (*Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *Endpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endpoint.Unmarshal(m, b)
}
func (m *Endpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endpoint.Marshal(b, m, deterministic)
}
func (m *Endpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint.Merge(m, src)
}
func (m *Endpoint) XXX_Size() int {
	return xxx_messageInfo_Endpoint.Size(m)
}
func (m *Endpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint proto.InternalMessageInfo

func (m *Endpoint) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Endpoint) GetHost() []string {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *Endpoint) GetPath() []string {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *Endpoint) GetMethod() []string {
	if m != nil {
		return m.Method
	}
	return nil
}

func (m *Endpoint) GetStream() bool {
	if m != nil {
		return m.Stream
	}
	return false
}

type EmptyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyResponse) Reset()         { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()    {}
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *EmptyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyResponse.Unmarshal(m, b)
}
func (m *EmptyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyResponse.Marshal(b, m, deterministic)
}
func (m *EmptyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyResponse.Merge(m, src)
}
func (m *EmptyResponse) XXX_Size() int {
	return xxx_messageInfo_EmptyResponse.Size(m)
}
func (m *EmptyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Endpoint)(nil), "go.micro.api.Endpoint")
	proto.RegisterType((*EmptyResponse)(nil), "go.micro.api.EmptyResponse")
}

func init() {
	proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c)
}

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0xd0, 0x41, 0x4a, 0xc5, 0x30,
	0x10, 0x06, 0x60, 0x63, 0x9f, 0x8f, 0xbe, 0x41, 0x11, 0xb2, 0x78, 0x04, 0xdd, 0x94, 0xae, 0xde,
	0x2a, 0x0b, 0x3d, 0x41, 0xd1, 0x5e, 0x20, 0x37, 0x88, 0x76, 0x68, 0xb3, 0x48, 0x66, 0x48, 0x06,
	0xc1, 0x43, 0x78, 0x67, 0x49, 0xad, 0x50, 0xdc, 0xba, 0xfb, 0xe7, 0x5b, 0xfc, 0xfc, 0x0c, 0x9c,
	0x3c, 0x07, 0xcb, 0x99, 0x84, 0xf4, 0xed, 0x4c, 0x36, 0x86, 0xf7, 0x4c, 0xd6, 0x73, 0xe8, 0x3f,
	0xa0, 0x1d, 0xd3, 0xc4, 0x14, 0x92, 0x68, 0x0d, 0x87, 0xe4, 0x23, 0x1a, 0xd5, 0xa9, 0xcb, 0xc9,
	0xad, 0xb9, 0xda, 0x42, 0x45, 0xcc, 0x75, 0xd7, 0x54, 0xab, 0xb9, 0x1a, 0x7b, 0x59, 0x4c, 0xf3,
	0x63, 0x35, 0xeb, 0x33, 0x1c, 0x23, 0xca, 0x42, 0x93, 0x39, 0xac, 0xba, 0x5d, 0xd5, 0x8b, 0x64,
	0xf4, 0xd1, 0xdc, 0x74, 0xea, 0xd2, 0xba, 0xed, 0xea, 0xef, 0xe1, 0x6e, 0x8c, 0x2c, 0x9f, 0x0e,
	0x0b, 0x53, 0x2a, 0xf8, 0xf4, 0xa5, 0xa0, 0x19, 0x38, 0xe8, 0x01, 0x5a, 0x87, 0x73, 0x28, 0x82,
	0x59, 0x9f, 0xed, 0x7e, 0xab, 0xfd, 0x1d, 0xfa, 0xf0, 0xf8, 0xc7, 0xf7, 0x45, 0xfd, 0x95, 0x7e,
	0x01, 0x78, 0xc5, 0xfc, 0xbf, 0x92, 0xb7, 0xe3, 0xfa, 0xad, 0xe7, 0xef, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x1f, 0xf0, 0xd9, 0x19, 0x3a, 0x01, 0x00, 0x00,
}
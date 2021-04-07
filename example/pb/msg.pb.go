// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg.proto

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

type Request struct {
	I                    string   `protobuf:"bytes,1,opt,name=I,proto3" json:"I,omitempty"`
	ID                   string   `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`
	Service              string   `protobuf:"bytes,3,opt,name=Service,proto3" json:"Service,omitempty"`
	Path                 string   `protobuf:"bytes,4,opt,name=Path,proto3" json:"Path,omitempty"`
	Version              uint32   `protobuf:"varint,5,opt,name=Version,proto3" json:"Version,omitempty"`
	ST                   int64    `protobuf:"varint,6,opt,name=ST,proto3" json:"ST,omitempty"`
	Data                 []byte   `protobuf:"bytes,7,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetI() string {
	if m != nil {
		return m.I
	}
	return ""
}

func (m *Request) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Request) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Request) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Request) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Request) GetST() int64 {
	if m != nil {
		return m.ST
	}
	return 0
}

func (m *Request) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type RequestAck struct {
	I                    string   `protobuf:"bytes,1,opt,name=I,proto3" json:"I,omitempty"`
	ID                   string   `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`
	ST                   int64    `protobuf:"varint,3,opt,name=ST,proto3" json:"ST,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestAck) Reset()         { *m = RequestAck{} }
func (m *RequestAck) String() string { return proto.CompactTextString(m) }
func (*RequestAck) ProtoMessage()    {}
func (*RequestAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{1}
}

func (m *RequestAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestAck.Unmarshal(m, b)
}
func (m *RequestAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestAck.Marshal(b, m, deterministic)
}
func (m *RequestAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestAck.Merge(m, src)
}
func (m *RequestAck) XXX_Size() int {
	return xxx_messageInfo_RequestAck.Size(m)
}
func (m *RequestAck) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestAck.DiscardUnknown(m)
}

var xxx_messageInfo_RequestAck proto.InternalMessageInfo

func (m *RequestAck) GetI() string {
	if m != nil {
		return m.I
	}
	return ""
}

func (m *RequestAck) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *RequestAck) GetST() int64 {
	if m != nil {
		return m.ST
	}
	return 0
}

type Response struct {
	I                    string   `protobuf:"bytes,1,opt,name=I,proto3" json:"I,omitempty"`
	ID                   string   `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`
	Service              string   `protobuf:"bytes,3,opt,name=Service,proto3" json:"Service,omitempty"`
	Path                 string   `protobuf:"bytes,4,opt,name=Path,proto3" json:"Path,omitempty"`
	Version              uint32   `protobuf:"varint,5,opt,name=Version,proto3" json:"Version,omitempty"`
	ST                   int64    `protobuf:"varint,6,opt,name=ST,proto3" json:"ST,omitempty"`
	Code                 int64    `protobuf:"varint,7,opt,name=Code,proto3" json:"Code,omitempty"`
	Message              string   `protobuf:"bytes,8,opt,name=Message,proto3" json:"Message,omitempty"`
	Data                 []byte   `protobuf:"bytes,9,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetI() string {
	if m != nil {
		return m.I
	}
	return ""
}

func (m *Response) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Response) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Response) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Response) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Response) GetST() int64 {
	if m != nil {
		return m.ST
	}
	return 0
}

func (m *Response) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Response) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ResponseAck struct {
	I                    string   `protobuf:"bytes,1,opt,name=I,proto3" json:"I,omitempty"`
	ID                   string   `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`
	ST                   int64    `protobuf:"varint,3,opt,name=ST,proto3" json:"ST,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseAck) Reset()         { *m = ResponseAck{} }
func (m *ResponseAck) String() string { return proto.CompactTextString(m) }
func (*ResponseAck) ProtoMessage()    {}
func (*ResponseAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{3}
}

func (m *ResponseAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseAck.Unmarshal(m, b)
}
func (m *ResponseAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseAck.Marshal(b, m, deterministic)
}
func (m *ResponseAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseAck.Merge(m, src)
}
func (m *ResponseAck) XXX_Size() int {
	return xxx_messageInfo_ResponseAck.Size(m)
}
func (m *ResponseAck) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseAck.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseAck proto.InternalMessageInfo

func (m *ResponseAck) GetI() string {
	if m != nil {
		return m.I
	}
	return ""
}

func (m *ResponseAck) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *ResponseAck) GetST() int64 {
	if m != nil {
		return m.ST
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "pb.Request")
	proto.RegisterType((*RequestAck)(nil), "pb.RequestAck")
	proto.RegisterType((*Response)(nil), "pb.Response")
	proto.RegisterType((*ResponseAck)(nil), "pb.ResponseAck")
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor_c06e4cca6c2cc899) }

var fileDescriptor_c06e4cca6c2cc899 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x51, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0xd5, 0x39, 0xa5, 0x69, 0x8e, 0xc2, 0xe0, 0xe9, 0xc6, 0x28, 0x93, 0x27, 0x16, 0x36, 0x98,
	0x10, 0x59, 0x3c, 0x20, 0x21, 0xa7, 0x62, 0x77, 0xcb, 0xa9, 0x54, 0x88, 0x3a, 0xc4, 0x86, 0x1f,
	0xe1, 0x8f, 0xf8, 0x32, 0xe4, 0x6b, 0x8d, 0xd8, 0x10, 0x13, 0xdb, 0x7b, 0x27, 0xdf, 0xf3, 0x7b,
	0xef, 0xb0, 0x79, 0x89, 0xdb, 0x8b, 0x71, 0x0a, 0x29, 0x68, 0x35, 0xae, 0xbb, 0x0f, 0xc0, 0xda,
	0xf1, 0xeb, 0x1b, 0xc7, 0xa4, 0x97, 0x08, 0x96, 0xa0, 0x05, 0xd3, 0x38, 0xb0, 0xfa, 0x1c, 0x95,
	0xed, 0x49, 0x09, 0x55, 0xb6, 0xd7, 0x84, 0xf5, 0xc0, 0xd3, 0xfb, 0x6e, 0xc3, 0x54, 0xc9, 0xb0,
	0x50, 0xad, 0x71, 0x76, 0xef, 0xd3, 0x13, 0xcd, 0x64, 0x2c, 0x38, 0xbf, 0x7e, 0xe0, 0x29, 0xee,
	0xc2, 0x9e, 0x4e, 0x5a, 0x30, 0x67, 0xae, 0xd0, 0xac, 0x3b, 0xac, 0x68, 0xde, 0x82, 0xa9, 0x9c,
	0x1a, 0x56, 0x79, 0xbb, 0xf7, 0xc9, 0x53, 0xdd, 0x82, 0x59, 0x3a, 0xc1, 0xdd, 0x15, 0xe2, 0xd1,
	0xd4, 0xcd, 0xe6, 0xf9, 0x17, 0x5f, 0x07, 0xbd, 0xaa, 0xe8, 0x75, 0x9f, 0x80, 0x0b, 0xc7, 0x71,
	0x0c, 0xfb, 0xc8, 0xff, 0x17, 0xe9, 0x36, 0x3c, 0xb2, 0x44, 0xaa, 0x9c, 0xe0, 0xbc, 0x7d, 0xc7,
	0x31, 0xfa, 0x2d, 0xd3, 0xe2, 0xf0, 0xd7, 0x91, 0x7e, 0x17, 0xd0, 0xfc, 0x28, 0xe0, 0x1a, 0x4f,
	0x4b, 0x86, 0x3f, 0x37, 0xb0, 0x9e, 0xcb, 0x79, 0x2f, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x74,
	0x46, 0xb6, 0xba, 0xeb, 0x01, 0x00, 0x00,
}
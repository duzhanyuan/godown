// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	Request
	Response
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Response_Type int32

const (
	Response_NIL    Response_Type = 0
	Response_OK     Response_Type = 1
	Response_STRING Response_Type = 2
	Response_INT    Response_Type = 3
	Response_SLICE  Response_Type = 4
	Response_HELP   Response_Type = 5
	Response_ERR    Response_Type = 6
)

var Response_Type_name = map[int32]string{
	0: "NIL",
	1: "OK",
	2: "STRING",
	3: "INT",
	4: "SLICE",
	5: "HELP",
	6: "ERR",
}
var Response_Type_value = map[string]int32{
	"NIL":    0,
	"OK":     1,
	"STRING": 2,
	"INT":    3,
	"SLICE":  4,
	"HELP":   5,
	"ERR":    6,
}

func (x Response_Type) String() string {
	return proto.EnumName(Response_Type_name, int32(x))
}
func (Response_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

// Request is a request to a server.
type Request struct {
	Command string `protobuf:"bytes,1,opt,name=command" json:"command,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

// Response is a response that contains command's execution result.
type Response struct {
	Result *Response_Result `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetResult() *Response_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

type Response_Result struct {
	Type  Response_Type `protobuf:"varint,1,opt,name=type,enum=Response_Type" json:"type,omitempty"`
	Item  string        `protobuf:"bytes,2,opt,name=item" json:"item,omitempty"`
	Items []string      `protobuf:"bytes,3,rep,name=items" json:"items,omitempty"`
}

func (m *Response_Result) Reset()                    { *m = Response_Result{} }
func (m *Response_Result) String() string            { return proto.CompactTextString(m) }
func (*Response_Result) ProtoMessage()               {}
func (*Response_Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Result) GetType() Response_Type {
	if m != nil {
		return m.Type
	}
	return Response_NIL
}

func (m *Response_Result) GetItem() string {
	if m != nil {
		return m.Item
	}
	return ""
}

func (m *Response_Result) GetItems() []string {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterType((*Response)(nil), "Response")
	proto.RegisterType((*Response_Result)(nil), "Response.Result")
	proto.RegisterEnum("Response_Type", Response_Type_name, Response_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Godown service

type GodownClient interface {
	ExecuteCommand(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type godownClient struct {
	cc *grpc.ClientConn
}

func NewGodownClient(cc *grpc.ClientConn) GodownClient {
	return &godownClient{cc}
}

func (c *godownClient) ExecuteCommand(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/Godown/ExecuteCommand", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Godown service

type GodownServer interface {
	ExecuteCommand(context.Context, *Request) (*Response, error)
}

func RegisterGodownServer(s *grpc.Server, srv GodownServer) {
	s.RegisterService(&_Godown_serviceDesc, srv)
}

func _Godown_ExecuteCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GodownServer).ExecuteCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Godown/ExecuteCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GodownServer).ExecuteCommand(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Godown_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Godown",
	HandlerType: (*GodownServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecuteCommand",
			Handler:    _Godown_ExecuteCommand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0xcf, 0x6a, 0xf3, 0x30,
	0x10, 0xc4, 0x3f, 0xff, 0x93, 0xed, 0xfd, 0xc0, 0x88, 0xa5, 0x07, 0x93, 0x53, 0x70, 0x29, 0xf8,
	0xe4, 0x82, 0xfb, 0x08, 0xc1, 0xa4, 0x6e, 0x8d, 0x5b, 0x14, 0xd3, 0x7b, 0x9a, 0xec, 0x21, 0x10,
	0x5b, 0x6e, 0x24, 0xd3, 0xe6, 0x4d, 0xfb, 0x38, 0x45, 0x72, 0xd2, 0x9e, 0x34, 0xa3, 0x9f, 0x06,
	0xed, 0x0e, 0xc4, 0xdb, 0xf1, 0x50, 0x8c, 0x27, 0xa9, 0x65, 0x76, 0x0b, 0xa1, 0xa0, 0x8f, 0x89,
	0x94, 0xc6, 0x14, 0xc2, 0x9d, 0xec, 0xfb, 0xed, 0xb0, 0x4f, 0x9d, 0xa5, 0x93, 0xc7, 0xe2, 0x6a,
	0xb3, 0x6f, 0x07, 0x22, 0x41, 0x6a, 0x94, 0x83, 0x22, 0xcc, 0x81, 0x9d, 0x48, 0x4d, 0x47, 0x6d,
	0x5f, 0xfd, 0x2f, 0x79, 0x71, 0x45, 0x46, 0x4c, 0x47, 0x2d, 0x2e, 0x7c, 0xf1, 0x06, 0x6c, 0xbe,
	0xc1, 0x0c, 0x7c, 0x7d, 0x1e, 0xc9, 0x26, 0x92, 0x32, 0xf9, 0x4b, 0x74, 0xe7, 0x91, 0x84, 0x65,
	0x88, 0xe0, 0x1f, 0x34, 0xf5, 0xa9, 0x6b, 0xff, 0xb6, 0x1a, 0x6f, 0x20, 0x30, 0xa7, 0x4a, 0xbd,
	0xa5, 0x97, 0xc7, 0x62, 0x36, 0xd9, 0x13, 0xf8, 0x26, 0x87, 0x21, 0x78, 0x6d, 0xdd, 0xf0, 0x7f,
	0xc8, 0xc0, 0x7d, 0x79, 0xe6, 0x0e, 0x02, 0xb0, 0x4d, 0x27, 0xea, 0x76, 0xcd, 0x5d, 0x03, 0xeb,
	0xb6, 0xe3, 0x1e, 0xc6, 0x10, 0x6c, 0x9a, 0x7a, 0x55, 0x71, 0x1f, 0x23, 0xf0, 0x1f, 0xab, 0xe6,
	0x95, 0x07, 0x86, 0x56, 0x42, 0x70, 0x56, 0xde, 0x03, 0x5b, 0xcb, 0xbd, 0xfc, 0x1c, 0xf0, 0x0e,
	0x92, 0xea, 0x8b, 0x76, 0x93, 0xa6, 0xd5, 0xbc, 0x36, 0x46, 0xc5, 0xa5, 0x9a, 0x45, 0xfc, 0x3b,
	0xf1, 0x3b, 0xb3, 0xbd, 0x3d, 0xfc, 0x04, 0x00, 0x00, 0xff, 0xff, 0xbe, 0xf4, 0xb7, 0xee, 0x44,
	0x01, 0x00, 0x00,
}
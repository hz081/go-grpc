// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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
	A                    int64    `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    int64    `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	C                    int64    `protobuf:"varint,3,opt,name=c,proto3" json:"c,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
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

func (m *Request) GetA() int64 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *Request) GetB() int64 {
	if m != nil {
		return m.B
	}
	return 0
}

func (m *Request) GetC() int64 {
	if m != nil {
		return m.C
	}
	return 0
}

type Response struct {
	Result               int64    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
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

func (m *Response) GetResult() int64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "proto.Request")
	proto.RegisterType((*Response)(nil), "proto.Response")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xc6, 0x5c,
	0xec, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x3c, 0x5c, 0x8c, 0x89, 0x12, 0x8c, 0x0a,
	0x8c, 0x1a, 0xcc, 0x41, 0x8c, 0x89, 0x20, 0x5e, 0x92, 0x04, 0x13, 0x84, 0x97, 0x04, 0xe2, 0x25,
	0x4b, 0x30, 0x43, 0x78, 0xc9, 0x4a, 0x4a, 0x5c, 0x1c, 0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5,
	0xa9, 0x42, 0x62, 0x5c, 0x6c, 0x45, 0xa9, 0xc5, 0xa5, 0x39, 0x25, 0x50, 0xad, 0x50, 0x9e, 0x51,
	0x37, 0x23, 0x17, 0x97, 0x63, 0x4a, 0x4a, 0x30, 0xc4, 0x52, 0x21, 0x35, 0x2e, 0x66, 0xc7, 0x94,
	0x14, 0x21, 0x3e, 0x88, 0xed, 0x7a, 0x50, 0x3b, 0xa5, 0xf8, 0xe1, 0x7c, 0xa8, 0x71, 0xda, 0x5c,
	0x1c, 0xbe, 0xa5, 0x39, 0x25, 0x99, 0x05, 0x39, 0x95, 0x84, 0x15, 0x6b, 0x70, 0xb1, 0xfa, 0x66,
	0xe6, 0x95, 0x16, 0x13, 0x54, 0x99, 0xc4, 0x06, 0xe6, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x78, 0x3e, 0xe5, 0x5b, 0x05, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AddServiceClient is the client API for AddService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AddServiceClient interface {
	Add(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Multiply(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Minus(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type addServiceClient struct {
	cc *grpc.ClientConn
}

func NewAddServiceClient(cc *grpc.ClientConn) AddServiceClient {
	return &addServiceClient{cc}
}

func (c *addServiceClient) Add(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.AddService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addServiceClient) Multiply(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.AddService/Multiply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addServiceClient) Minus(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.AddService/Minus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddServiceServer is the server API for AddService service.
type AddServiceServer interface {
	Add(context.Context, *Request) (*Response, error)
	Multiply(context.Context, *Request) (*Response, error)
	Minus(context.Context, *Request) (*Response, error)
}

func RegisterAddServiceServer(s *grpc.Server, srv AddServiceServer) {
	s.RegisterService(&_AddService_serviceDesc, srv)
}

func _AddService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AddService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServiceServer).Add(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddService_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServiceServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AddService/Multiply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServiceServer).Multiply(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddService_Minus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServiceServer).Minus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AddService/Minus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServiceServer).Minus(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _AddService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AddService",
	HandlerType: (*AddServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _AddService_Add_Handler,
		},
		{
			MethodName: "Multiply",
			Handler:    _AddService_Multiply_Handler,
		},
		{
			MethodName: "Minus",
			Handler:    _AddService_Minus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

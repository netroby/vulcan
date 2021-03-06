// Code generated by protoc-gen-go.
// source: cacher.proto
// DO NOT EDIT!

/*
Package cacher is a generated protocol buffer package.

It is generated from these files:
	cacher.proto

It has these top-level messages:
	ChunksRequest
	ChunksResponse
*/
package cacher

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

type ChunksRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// after is a timestamp provided in milliseconds since the unix epoch.
	After int64 `protobuf:"varint,2,opt,name=after" json:"after,omitempty"`
}

func (m *ChunksRequest) Reset()                    { *m = ChunksRequest{} }
func (m *ChunksRequest) String() string            { return proto.CompactTextString(m) }
func (*ChunksRequest) ProtoMessage()               {}
func (*ChunksRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ChunksRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ChunksRequest) GetAfter() int64 {
	if m != nil {
		return m.After
	}
	return 0
}

type ChunksResponse struct {
	Chunks [][]byte `protobuf:"bytes,1,rep,name=chunks,proto3" json:"chunks,omitempty"`
}

func (m *ChunksResponse) Reset()                    { *m = ChunksResponse{} }
func (m *ChunksResponse) String() string            { return proto.CompactTextString(m) }
func (*ChunksResponse) ProtoMessage()               {}
func (*ChunksResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ChunksResponse) GetChunks() [][]byte {
	if m != nil {
		return m.Chunks
	}
	return nil
}

func init() {
	proto.RegisterType((*ChunksRequest)(nil), "cacher.ChunksRequest")
	proto.RegisterType((*ChunksResponse)(nil), "cacher.ChunksResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Cacher service

type CacherClient interface {
	Chunks(ctx context.Context, in *ChunksRequest, opts ...grpc.CallOption) (*ChunksResponse, error)
}

type cacherClient struct {
	cc *grpc.ClientConn
}

func NewCacherClient(cc *grpc.ClientConn) CacherClient {
	return &cacherClient{cc}
}

func (c *cacherClient) Chunks(ctx context.Context, in *ChunksRequest, opts ...grpc.CallOption) (*ChunksResponse, error) {
	out := new(ChunksResponse)
	err := grpc.Invoke(ctx, "/cacher.Cacher/Chunks", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Cacher service

type CacherServer interface {
	Chunks(context.Context, *ChunksRequest) (*ChunksResponse, error)
}

func RegisterCacherServer(s *grpc.Server, srv CacherServer) {
	s.RegisterService(&_Cacher_serviceDesc, srv)
}

func _Cacher_Chunks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChunksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacherServer).Chunks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cacher.Cacher/Chunks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacherServer).Chunks(ctx, req.(*ChunksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cacher_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cacher.Cacher",
	HandlerType: (*CacherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Chunks",
			Handler:    _Cacher_Chunks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cacher.proto",
}

func init() { proto.RegisterFile("cacher.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 154 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x4e, 0x4c, 0xce,
	0x48, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0x4c, 0xb9, 0x78,
	0x9d, 0x33, 0x4a, 0xf3, 0xb2, 0x8b, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xf8, 0xb8,
	0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x98, 0x32, 0x53, 0x84, 0x44, 0xb8,
	0x58, 0x13, 0xd3, 0x4a, 0x52, 0x8b, 0x24, 0x98, 0x14, 0x18, 0x35, 0x98, 0x83, 0x20, 0x1c, 0x25,
	0x0d, 0x2e, 0x3e, 0x98, 0xb6, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x31, 0x2e, 0xb6, 0x64,
	0xb0, 0x88, 0x04, 0xa3, 0x02, 0xb3, 0x06, 0x4f, 0x10, 0x94, 0x67, 0xe4, 0xcc, 0xc5, 0xe6, 0x0c,
	0xb6, 0x4a, 0xc8, 0x92, 0x8b, 0x0d, 0xa2, 0x47, 0x48, 0x54, 0x0f, 0xea, 0x16, 0x14, 0xab, 0xa5,
	0xc4, 0xd0, 0x85, 0x21, 0x46, 0x2b, 0x31, 0x24, 0xb1, 0x81, 0x1d, 0x6d, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0x8a, 0x07, 0x7a, 0x34, 0xc4, 0x00, 0x00, 0x00,
}

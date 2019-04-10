// Code generated by protoc-gen-go. DO NOT EDIT.
// source: recom-engine.proto

package recomengine

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RecomEngineClient is the client API for RecomEngine service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RecomEngineClient interface {
	GetMultiProductRecomByUserID(ctx context.Context, in *GetMultiProductRecomRequest, opts ...grpc.CallOption) (*GetMultiProductRecomResponse, error)
}

type recomEngineClient struct {
	cc *grpc.ClientConn
}

func NewRecomEngineClient(cc *grpc.ClientConn) RecomEngineClient {
	return &recomEngineClient{cc}
}

func (c *recomEngineClient) GetMultiProductRecomByUserID(ctx context.Context, in *GetMultiProductRecomRequest, opts ...grpc.CallOption) (*GetMultiProductRecomResponse, error) {
	out := new(GetMultiProductRecomResponse)
	err := c.cc.Invoke(ctx, "/recomengine.RecomEngine/GetMultiProductRecomByUserID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecomEngineServer is the server API for RecomEngine service.
type RecomEngineServer interface {
	GetMultiProductRecomByUserID(context.Context, *GetMultiProductRecomRequest) (*GetMultiProductRecomResponse, error)
}

func RegisterRecomEngineServer(s *grpc.Server, srv RecomEngineServer) {
	s.RegisterService(&_RecomEngine_serviceDesc, srv)
}

func _RecomEngine_GetMultiProductRecomByUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMultiProductRecomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecomEngineServer).GetMultiProductRecomByUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/recomengine.RecomEngine/GetMultiProductRecomByUserID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecomEngineServer).GetMultiProductRecomByUserID(ctx, req.(*GetMultiProductRecomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RecomEngine_serviceDesc = grpc.ServiceDesc{
	ServiceName: "recomengine.RecomEngine",
	HandlerType: (*RecomEngineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMultiProductRecomByUserID",
			Handler:    _RecomEngine_GetMultiProductRecomByUserID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "recom-engine.proto",
}

func init() { proto.RegisterFile("recom-engine.proto", fileDescriptor_recom_engine_5b5a6791e12e18d6) }

var fileDescriptor_recom_engine_5b5a6791e12e18d6 = []byte{
	// 130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x4a, 0x4d, 0xce,
	0xcf, 0xd5, 0x4d, 0xcd, 0x4b, 0xcf, 0xcc, 0x4b, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x06, 0x8b, 0x41, 0x84, 0xa4, 0x84, 0x0b, 0x8a, 0xf2, 0x53, 0x4a, 0x93, 0x4b, 0x74, 0xc1, 0x82,
	0x10, 0x15, 0x46, 0x4d, 0x8c, 0x5c, 0xdc, 0x41, 0x20, 0xbe, 0x2b, 0x58, 0x91, 0x50, 0x31, 0x97,
	0x8c, 0x7b, 0x6a, 0x89, 0x6f, 0x69, 0x4e, 0x49, 0x66, 0x00, 0x44, 0x39, 0x58, 0xd6, 0xa9, 0x32,
	0xb4, 0x38, 0xb5, 0xc8, 0xd3, 0x45, 0x48, 0x43, 0x0f, 0xc9, 0x48, 0x3d, 0x6c, 0x4a, 0x83, 0x52,
	0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0xa4, 0x34, 0x89, 0x50, 0x59, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x9a,
	0xc4, 0x06, 0x76, 0x8b, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x28, 0x40, 0xf8, 0xc3, 0x00,
	0x00, 0x00,
}
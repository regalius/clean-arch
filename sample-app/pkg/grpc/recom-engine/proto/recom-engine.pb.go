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
	GetMultiAffinityProductRecom(ctx context.Context, in *GetProductRecomRequest, opts ...grpc.CallOption) (*GetProductRecomResponse, error)
	GetGenderAffinityProductRecom(ctx context.Context, in *GetProductRecomRequest, opts ...grpc.CallOption) (*GetProductRecomResponse, error)
	GetUserAffinityProductRecom(ctx context.Context, in *GetProductRecomRequest, opts ...grpc.CallOption) (*GetProductRecomResponse, error)
}

type recomEngineClient struct {
	cc *grpc.ClientConn
}

func NewRecomEngineClient(cc *grpc.ClientConn) RecomEngineClient {
	return &recomEngineClient{cc}
}

func (c *recomEngineClient) GetMultiAffinityProductRecom(ctx context.Context, in *GetProductRecomRequest, opts ...grpc.CallOption) (*GetProductRecomResponse, error) {
	out := new(GetProductRecomResponse)
	err := c.cc.Invoke(ctx, "/recomengine.RecomEngine/GetMultiAffinityProductRecom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recomEngineClient) GetGenderAffinityProductRecom(ctx context.Context, in *GetProductRecomRequest, opts ...grpc.CallOption) (*GetProductRecomResponse, error) {
	out := new(GetProductRecomResponse)
	err := c.cc.Invoke(ctx, "/recomengine.RecomEngine/GetGenderAffinityProductRecom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recomEngineClient) GetUserAffinityProductRecom(ctx context.Context, in *GetProductRecomRequest, opts ...grpc.CallOption) (*GetProductRecomResponse, error) {
	out := new(GetProductRecomResponse)
	err := c.cc.Invoke(ctx, "/recomengine.RecomEngine/GetUserAffinityProductRecom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecomEngineServer is the server API for RecomEngine service.
type RecomEngineServer interface {
	GetMultiAffinityProductRecom(context.Context, *GetProductRecomRequest) (*GetProductRecomResponse, error)
	GetGenderAffinityProductRecom(context.Context, *GetProductRecomRequest) (*GetProductRecomResponse, error)
	GetUserAffinityProductRecom(context.Context, *GetProductRecomRequest) (*GetProductRecomResponse, error)
}

func RegisterRecomEngineServer(s *grpc.Server, srv RecomEngineServer) {
	s.RegisterService(&_RecomEngine_serviceDesc, srv)
}

func _RecomEngine_GetMultiAffinityProductRecom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductRecomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecomEngineServer).GetMultiAffinityProductRecom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/recomengine.RecomEngine/GetMultiAffinityProductRecom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecomEngineServer).GetMultiAffinityProductRecom(ctx, req.(*GetProductRecomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecomEngine_GetGenderAffinityProductRecom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductRecomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecomEngineServer).GetGenderAffinityProductRecom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/recomengine.RecomEngine/GetGenderAffinityProductRecom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecomEngineServer).GetGenderAffinityProductRecom(ctx, req.(*GetProductRecomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecomEngine_GetUserAffinityProductRecom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductRecomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecomEngineServer).GetUserAffinityProductRecom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/recomengine.RecomEngine/GetUserAffinityProductRecom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecomEngineServer).GetUserAffinityProductRecom(ctx, req.(*GetProductRecomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RecomEngine_serviceDesc = grpc.ServiceDesc{
	ServiceName: "recomengine.RecomEngine",
	HandlerType: (*RecomEngineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMultiAffinityProductRecom",
			Handler:    _RecomEngine_GetMultiAffinityProductRecom_Handler,
		},
		{
			MethodName: "GetGenderAffinityProductRecom",
			Handler:    _RecomEngine_GetGenderAffinityProductRecom_Handler,
		},
		{
			MethodName: "GetUserAffinityProductRecom",
			Handler:    _RecomEngine_GetUserAffinityProductRecom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "recom-engine.proto",
}

func init() { proto.RegisterFile("recom-engine.proto", fileDescriptor_recom_engine_c70ce3d87bce2c52) }

var fileDescriptor_recom_engine_c70ce3d87bce2c52 = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x4a, 0x4d, 0xce,
	0xcf, 0xd5, 0x4d, 0xcd, 0x4b, 0xcf, 0xcc, 0x4b, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x06, 0x8b, 0x41, 0x84, 0xa4, 0x84, 0x0b, 0x8a, 0xf2, 0x53, 0x4a, 0x93, 0x4b, 0x74, 0xc1, 0x82,
	0x10, 0x15, 0x46, 0xe7, 0x98, 0xb8, 0xb8, 0x83, 0x40, 0x7c, 0x57, 0xb0, 0x22, 0xa1, 0x4c, 0x2e,
	0x19, 0xf7, 0xd4, 0x12, 0xdf, 0xd2, 0x9c, 0x92, 0x4c, 0xc7, 0xb4, 0xb4, 0xcc, 0xbc, 0xcc, 0x92,
	0xca, 0x00, 0x88, 0x36, 0xb0, 0x2a, 0x21, 0x65, 0x3d, 0x24, 0x23, 0xf5, 0xdc, 0x53, 0x4b, 0x90,
	0x65, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0xa4, 0x54, 0xf0, 0x2b, 0x2a, 0x2e, 0xc8, 0xcf,
	0x2b, 0x4e, 0x15, 0xca, 0xe2, 0x92, 0x75, 0x4f, 0x2d, 0x71, 0x4f, 0xcd, 0x4b, 0x49, 0x2d, 0xa2,
	0xb5, 0x5d, 0x19, 0x5c, 0xd2, 0xee, 0xa9, 0x25, 0xa1, 0xc5, 0x34, 0xb7, 0x29, 0x89, 0x0d, 0x1c,
	0xae, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x73, 0x6c, 0xb2, 0xaf, 0x8f, 0x01, 0x00, 0x00,
}

package client

import (
	"context"

	pb "github.com/regalius/clean-arch/sample-app/pkg/grpc/recom-engine/proto"
	"google.golang.org/grpc"
)

// GetMultiAffinityProductRecom get product recommendation by multi affinity
func (c GRPCClient) GetMultiAffinityProductRecom(ctx context.Context, in *pb.GetProductRecomRequest, opts ...grpc.CallOption) (*pb.GetProductRecomResponse, error) {
	return c.recomEngineClient.GetMultiAffinityProductRecom(ctx, in, opts...)
}

// GetGenderAffinityProductRecom get product recommendation by gender affinity
func (c GRPCClient) GetGenderAffinityProductRecom(ctx context.Context, in *pb.GetProductRecomRequest, opts ...grpc.CallOption) (*pb.GetProductRecomResponse, error) {
	return c.recomEngineClient.GetGenderAffinityProductRecom(ctx, in, opts...)
}

// GetUserAffinityProductRecom get product recommendation by user affinity
func (c GRPCClient) GetUserAffinityProductRecom(ctx context.Context, in *pb.GetProductRecomRequest, opts ...grpc.CallOption) (*pb.GetProductRecomResponse, error) {
	return c.recomEngineClient.GetUserAffinityProductRecom(ctx, in, opts...)
}

package client

import (
	"context"

	pb "github.com/regalius/clean-arch/sample-app/pkg/grpc/recom-engine/proto"
	"google.golang.org/grpc"
)

// GetMultiProductRecomByUserID get product recommendation by user id with multiple affinity
func (c GRPCClient) GetMultiProductRecomByUserID(ctx context.Context, in *pb.GetMultiProductRecomRequest, opts ...grpc.CallOption) (response *pb.GetMultiProductRecomResponse, err error) {
	return c.recomEngineClient.GetMultiProductRecomByUserID(ctx, in, opts...)
}

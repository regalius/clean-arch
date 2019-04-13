package client

import (
	"log"

	pb "github.com/regalius/clean-arch/sample-app/pkg/grpc/recom-engine/proto"
	"google.golang.org/grpc"
)

// GRPCClient defines client structure
type GRPCClient struct {
	recomEngineClient pb.RecomEngineClient
	recomEngineConn   *grpc.ClientConn
	options           Options
}

// NewGRPCClient create new grpc client instance
func NewGRPCClient(o Options) (client pb.RecomEngineClient) {
	conn, err := newGRPCConnection(o)
	if err != nil {
		log.Println("[GRPC/RecomEngine] Failed to create grpc connection", err)
		return
	}
	grpcClient := pb.NewRecomEngineClient(conn)
	client = GRPCClient{
		recomEngineClient: grpcClient,
		recomEngineConn:   conn,
		options:           o,
	}
	return
}

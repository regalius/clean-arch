package grpc

import (
	pRecomUcase "github.com/regalius/clean-arch/sample-app/internal/usecase/product-recom"
	pb "github.com/regalius/clean-arch/sample-app/pkg/grpc/recom-engine/proto"
	"google.golang.org/grpc"
)

type GrpcRecomEngineDelivery struct {
	server            *grpc.Server
	configName        string
	pRecomGenderUcase pRecomUcase.Usecase
	pRecomUserUcase   pRecomUcase.Usecase
	pRecomMultiUcase  pRecomUcase.Usecase
}

// NewGrpcRecomEngineDelivery Generate new GRPC Delivery
func NewGrpcRecomEngineDelivery(
	server *grpc.Server,
	configName string,
	pRecomGenderUcase pRecomUcase.Usecase,
	pRecomUserUcase pRecomUcase.Usecase,
	pRecomMultiUcase pRecomUcase.Usecase) pb.RecomEngineServer {

	d := GrpcRecomEngineDelivery{
		server,
		configName,
		pRecomGenderUcase,
		pRecomUserUcase,
		pRecomMultiUcase,
	}
	pb.RegisterRecomEngineServer(server, d)
	return d
}

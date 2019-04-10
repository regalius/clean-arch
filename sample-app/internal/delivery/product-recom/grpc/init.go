package grpc

import (
	pRecomUcase "github.com/regalius/clean-arch/sample-app/internal/usecase/product-recom"
	pb "github.com/regalius/clean-arch/sample-app/pkg/grpc/recom-engine/proto"
	"google.golang.org/grpc"
)

type grpcDelivery struct {
	server            *grpc.Server
	pRecomGenderUcase pRecomUcase.Usecase
	pRecomUserUcase   pRecomUcase.Usecase
	pRecomMultiUcase  pRecomUcase.Usecase
}

func NewGRPCDelivery(
	server *grpc.Server,
	pRecomGenderUcase pRecomUcase.Usecase,
	pRecomUserUcase pRecomUcase.Usecase,
	pRecomMultiUcase pRecomUcase.Usecase) pb.RecomEngineServer {

	return grpcDelivery{
		server,
		pRecomGenderUcase,
		pRecomUserUcase,
		pRecomMultiUcase,
	}
}

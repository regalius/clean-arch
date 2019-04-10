package grpc

import (
	"context"
	"errors"
	"log"

	pRecomUCase "github.com/regalius/clean-arch/sample-app/internal/usecase/product-recom"
	pb "github.com/regalius/clean-arch/sample-app/pkg/grpc/recom-engine/proto"
)

func (d grpcDelivery) GetMultiProductRecomByUserID(ctx context.Context, in *pb.GetMultiProductRecomRequest) (response *pb.GetMultiProductRecomResponse, err error) {
	if in == nil {
		log.Println("[RecomEngine/GRPC] GetMultiProductRecomByUserID input is nil")
		err = errors.New("GetMultiProductRecomByUserID input is nil")
		return
	}
	userID := in.UserID
	limit := int(in.Limit)
	recomm, err := d.pRecomMultiUcase.GetRecommendationByUserID(ctx, userID, pRecomUCase.GetRecommendationByUserIDOptions{
		Limit: limit,
	})
	response = mapProductRecomResultToResponse(recomm)
	return
}

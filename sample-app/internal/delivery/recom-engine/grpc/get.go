package grpc

import (
	"context"
	"errors"
	"log"

	"github.com/regalius/clean-arch/sample-app/common/config"
	pRecomUCase "github.com/regalius/clean-arch/sample-app/internal/usecase/product-recom"
	pb "github.com/regalius/clean-arch/sample-app/pkg/grpc/recom-engine/proto"
)

func (d GrpcRecomEngineDelivery) GetMultiAffinityProductRecom(ctx context.Context, in *pb.GetProductRecomRequest) (*pb.GetProductRecomResponse, error) {
	if in == nil {
		log.Println("[RecomEngine/GRPC] GetMultiAffinityProductRecom input is nil")
		err := errors.New("GetMultiAffinityProductRecom input is nil")
		return nil, err
	}

	userID := in.UserID
	limit := int(in.Limit)

	rootURLInterf, err := config.Get(d.configName, "ProductRecom.rootURL")
	if err != nil {
		return nil, err
	}

	rootURL, ok := rootURLInterf.(string)
	if !ok {
		log.Println("[RecomEngine/GRPC] Can't assert rootURL to string")
		return nil, nil
	}
	recomm, err := d.pRecomMultiUcase.GetRecommendationByUserID(ctx, userID, pRecomUCase.GetRecommendationByUserIDOptions{
		Limit: limit,
	})
	if err != nil {
		return nil, err
	}
	response := mapProductRecomResultToResponse(recomm, rootURL)
	return response, nil
}

func (d GrpcRecomEngineDelivery) GetGenderAffinityProductRecom(ctx context.Context, in *pb.GetProductRecomRequest) (*pb.GetProductRecomResponse, error) {
	if in == nil {
		log.Println("[RecomEngine/GRPC] GetGenderAffinityProductRecom input is nil")
		err := errors.New("GetGenderAffinityProductRecom input is nil")
		return nil, err
	}
	userID := in.UserID
	limit := int(in.Limit)

	rootURLInterf, err := config.Get(d.configName, "ProductRecom.rootURL")
	if err != nil {
		return nil, err
	}

	rootURL, ok := rootURLInterf.(string)
	if !ok {
		log.Println("[RecomEngine/GRPC] Can't assert rootURL to string")
		return nil, nil
	}
	recomm, err := d.pRecomGenderUcase.GetRecommendationByUserID(ctx, userID, pRecomUCase.GetRecommendationByUserIDOptions{
		Limit: limit,
	})
	if err != nil {
		return nil, err
	}
	response := mapProductRecomResultToResponse(recomm, rootURL)
	return response, nil
}

func (d GrpcRecomEngineDelivery) GetUserAffinityProductRecom(ctx context.Context, in *pb.GetProductRecomRequest) (*pb.GetProductRecomResponse, error) {
	if in == nil {
		log.Println("[RecomEngine/GRPC] GetGenderAffinityProductRecom input is nil")
		err := errors.New("GetGenderAffinityProductRecom input is nil")
		return nil, err
	}
	userID := in.UserID
	limit := int(in.Limit)

	rootURLInterf, err := config.Get(d.configName, "ProductRecom.rootURL")
	if err != nil {
		return nil, err
	}

	rootURL, ok := rootURLInterf.(string)
	if !ok {
		log.Println("[RecomEngine/GRPC] Can't assert rootURL to string")
		return nil, nil
	}
	recomm, err := d.pRecomUserUcase.GetRecommendationByUserID(ctx, userID, pRecomUCase.GetRecommendationByUserIDOptions{
		Limit: limit,
	})
	if err != nil {
		return nil, err
	}
	response := mapProductRecomResultToResponse(recomm, rootURL)
	return response, nil
}

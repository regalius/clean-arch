package grpc

import (
	pRecomUCase "github.com/regalius/clean-arch/sample-app/internal/usecase/product-recom"
	pb "github.com/regalius/clean-arch/sample-app/pkg/grpc/recom-engine/proto"
)

func mapProductRecomResultToResponse(recomm pRecomUCase.SingleUserResult, rootURL string) (response *pb.GetProductRecomResponse) {
	response.Meta.OverallAffinities = recomm.Meta.OverallAffinities
	response.Data = []*pb.ProductRecom{}
	for _, product := range recomm.Data {
		productItem := &pb.ProductRecom{
			Id:       product.ID,
			Price:    product.Price,
			ShopId:   product.ShopID,
			Created:  product.TimeCreated,
			Url:      product.GetURL(rootURL),
			Name:     product.Name,
			Affinity: product.Affinity,
		}

		for _, image := range product.Images {
			productItem.Images = append(productItem.Images, image.GetFullImagePath(rootURL))
		}

		response.Data = append(response.Data, productItem)
	}

	return
}

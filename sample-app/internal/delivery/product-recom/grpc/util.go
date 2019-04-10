package grpc

import (
	pRecomUCase "github.com/regalius/clean-arch/sample-app/internal/usecase/product-recom"
	pb "github.com/regalius/clean-arch/sample-app/pkg/grpc/recom-engine/proto"
)

func mapProductRecomResultToResponse(recomm pRecomUCase.SingleUserResult) (response *pb.GetMultiProductRecomResponse) {
	return
}

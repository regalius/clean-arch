package useraffinity

import (
	"context"
	"log"

	sliceUtils "github.com/regalius/clean-arch/sample-app/common/slice"
	uPARepo "github.com/regalius/clean-arch/sample-app/internal/repository/user-product-affinity"
	pRecomUCase "github.com/regalius/clean-arch/sample-app/internal/usecase/product-recom"
)

func (uc productRecomUserUsecase) GetRecommendationByUserID(ctx context.Context, userID int64, options pRecomUCase.GetRecommendationByUserIDOptions) (recommendation pRecomUCase.SingleUserResult, err error) {
	user, err := uc.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		log.Println("[Usecase/PRecom/UAffinity] Failed to fetch user", err)
		return
	}

	if options.Limit == 0 {
		options.Limit = pRecomUCase.DefaultGetRecommendationByUserIDOptions.Limit
	}

	uPAffinity, err := uc.userPAffinityRepo.GetUserProductAffinitiesByUserID(
		ctx,
		user.ID,
		uPARepo.GetUserProductAffinitiesByUserIDOptions{
			Limit: options.Limit,
		})
	if err != nil {
		log.Println("[Usecase/PRecom/UAffinity] Failed to fetch user product affinity", err)
		return
	}

	var selectedPIDs []int64
	for _, affinity := range uPAffinity {
		selectedPIDs = append(selectedPIDs, affinity.ProductID)
	}

	products, err := uc.productRepo.GetByIDList(ctx, selectedPIDs)

	var productRecoms []pRecomUCase.ProductRecom
	for _, product := range products {
		affinityArrIndex := sliceUtils.IndexOfInt64(selectedPIDs, product.ID)
		productRecom := pRecomUCase.ProductRecom{
			Product:  product,
			Affinity: uPAffinity[affinityArrIndex].AffinityScore,
		}
		productRecoms = append(productRecoms, productRecom)
	}

	recommendation.SetData(productRecoms)
	return
}

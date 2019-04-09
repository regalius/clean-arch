package useraffinity

import (
	"context"
	"log"

	pRecomUCase "github.com/regalius/clean-arch/sample-app/internal/usecase/product-recom"
	sliceUtils "github.com/regalius/clean-arch/sample-app/common/slice"
)

func (uc productRecomUserUsecase) GetRecommendationByUserID(userID int64, options pRecomUCase.GetRecommendationByUserIDoptions) (recommendation pRecomUCase.SingleUserResult, err error) {
	ctx := context.Background()
	user, err := uc.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		log.Println("[Usecase/PRecom/UAffinity] Failed to fetch user", err)
		return
	}
	uPAffinity, err := userPAffinityRepo.GetUserProductAffinityByUserID(ctx, user.ID)
	if err != nil {
		log.Println("[Usecase/PRecom/UAffinity] Failed to fetch user product affinity", err)
		return
	}

	processedAffinities := boostWithMagicFormula(uPAffinity, options.Limit)

	var selectedPIDs []int64
	for _, affinity := range processedAffinities {
		selectedPIDs = append(selectedPIDs, affinity.ProductID)
	}

	products, err := uc.productRepo.GetByIDList(ctx, selectedPIDs)
	
	var productRecoms []pRecomUCase.ProductRecom
	for index, product := products {
		affinityArrIndex := sliceUtils.IndexOfInt64(selectedPIDs, product.ID)

		productRecom := pRecomUCase.ProductRecom{
			Product: product,
			Affinity: processedAffinities[affinityArrIndex],
		}
	}

	recommendation.SetData(productRecoms)
	return
}

package genderaffinity

import (
	"context"

	gPAModel "github.com/regalius/clean-arch/sample-app/internal/model/gender-product-affinity"
	gPARepo "github.com/regalius/clean-arch/sample-app/internal/repository/gender-product-affinity"
	pRecomUCase "github.com/regalius/clean-arch/sample-app/internal/usecase/product-recom"
)

func (uc productRecomGenderUsecase) GetRecommendationByUserID(userID int64, options pRecomUCase.GetRecommendationByUserIDoptions) (recommendation pRecomUCase.SingleUserResult, err error) {
	ctx := context.Background()
	user, err := uc.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		log.Println("[Usecase/PRecom/GAffinity] Failed to fetch user", err)
		return
	}
	gPAffinity, err := genderPAffinityRepo.GetGenderProductAffinitiesByGender(
		ctx, user.Gender, genderThreshold,
		gPARepo.GetGenderProductAffinitiesByGenderOptions{
			Limit: options.Limit,
		})

	if err != nil {
		log.Println("[Usecase/PRecom/GAffinity] Failed to fetch gender product affinity", err)
		return
	}

	processedAffinities := boostWithMagicFormula(gPAffinity, options.Limit)

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

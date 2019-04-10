package genderaffinity

import (
	"context"
	"log"

	sliceUtils "github.com/regalius/clean-arch/sample-app/common/slice"
	gPARepo "github.com/regalius/clean-arch/sample-app/internal/repository/gender-product-affinity"
	pRecomUCase "github.com/regalius/clean-arch/sample-app/internal/usecase/product-recom"
)

func (uc productRecomGenderUsecase) GetRecommendationByUserID(ctx context.Context, userID int64, options pRecomUCase.GetRecommendationByUserIDOptions) (recommendation pRecomUCase.SingleUserResult, err error) {
	user, err := uc.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		log.Println("[Usecase/PRecom/GAffinity] Failed to fetch user", err)
		return
	}

	if options.Limit == 0 {
		options.Limit = pRecomUCase.DefaultGetRecommendationByUserIDOptions.Limit
	}

	gPAffinity, err := uc.genderPAffinityRepo.GetGenderProductAffinitiesByGender(
		ctx, user.Gender, genderThreshold,
		gPARepo.GetGenderProductAffinitiesByGenderOptions{
			Limit: options.Limit,
		})

	if err != nil {
		log.Println("[Usecase/PRecom/GAffinity] Failed to fetch gender product affinity", err)
		return
	}

	var selectedPIDs []int64
	for _, affinity := range gPAffinity {
		selectedPIDs = append(selectedPIDs, affinity.ProductID)
	}

	products, err := uc.productRepo.GetByIDList(ctx, selectedPIDs)

	var productRecoms []pRecomUCase.ProductRecom
	for _, product := range products {
		affinityArrIndex := sliceUtils.IndexOfInt64(selectedPIDs, product.ID)

		productRecom := pRecomUCase.ProductRecom{
			Product:  product,
			Affinity: gPAffinity[affinityArrIndex].AffinityScore,
		}
		productRecoms = append(productRecoms, productRecom)
	}

	recommendation.SetData(productRecoms)
	return
}

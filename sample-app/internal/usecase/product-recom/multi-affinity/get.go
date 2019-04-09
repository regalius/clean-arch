package multiaffinity

import (
	"context"

	gPARepo "github.com/regalius/clean-arch/sample-app/internal/repository/gender-product-affinity"
	gPAModel "github.com/regalius/clean-arch/sample-app/internal/model/gender-product-affinity"
	uPAModel "github.com/regalius/clean-arch/sample-app/internal/model/user-product-affinity"
	pRecomUCase "github.com/regalius/clean-arch/sample-app/internal/usecase/product-recom"
)

func (uc productRecomDefaultUsecase) GetRecommendationByUserID(userID int64, options pRecomUCase.GetRecommendationByUserIDoptions) (recommendation pRecomUCase.SingleUserResult, err error) {

	ctx := context.Background()
	user, err := uc.userRepo.GetUserByID(ctx, userID)
	
	var wg sync.WaitGroup
	var gPAffinities []gPAModel.GenderProductAffinityWithScore
	var uPAffinities []uPAModel.UserProductAffinity

	wg.Add(1)
	go func() {
		defer wg.Done()
		gPAffinities, err = genderPAffinityRepo.GetGenderProductAffinitiesByGender(
			ctx, user.Gender, genderThreshold,
			gPARepo.GetGenderProductAffinitiesByGenderOptions{

		})
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		uPAffinities, err = userPAffinityRepo.GetUserProductAffinityByUserID(ctx, user.ID)
	}
	wg.Wait()

	processedAffinities := boostWithMagicFormula(gPAffinities, uPAffinities)

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

package multiaffinity

import (
	"context"
	"sync"

	sliceUtils "github.com/regalius/clean-arch/sample-app/common/slice"
	gPAModel "github.com/regalius/clean-arch/sample-app/internal/model/gender-product-affinity"
	uPAModel "github.com/regalius/clean-arch/sample-app/internal/model/user-product-affinity"
	gPARepo "github.com/regalius/clean-arch/sample-app/internal/repository/gender-product-affinity"
	uPARepo "github.com/regalius/clean-arch/sample-app/internal/repository/user-product-affinity"
	pRecomUCase "github.com/regalius/clean-arch/sample-app/internal/usecase/product-recom"
)

func (uc productRecomMultiUsecase) GetRecommendationByUserID(ctx context.Context, userID int64, options pRecomUCase.GetRecommendationByUserIDOptions) (recommendation pRecomUCase.SingleUserResult, err error) {
	user, err := uc.userRepo.GetUserByID(ctx, userID)

	var wg sync.WaitGroup
	var gPAffinities []gPAModel.GenderProductAffinityWithScore
	var uPAffinities []uPAModel.UserProductAffinity

	wg.Add(1)
	go func() {
		defer wg.Done()
		gPAffinities, err = uc.genderPAffinityRepo.GetGenderProductAffinitiesByGender(
			ctx, user.Gender, genderThreshold,
			gPARepo.GetGenderProductAffinitiesByGenderOptions{
				Limit: options.Limit,
			})
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		uPAffinities, err = uc.userPAffinityRepo.GetUserProductAffinitiesByUserID(
			ctx,
			user.ID,
			uPARepo.GetUserProductAffinitiesByUserIDOptions{},
		)
	}()
	wg.Wait()

	processedAffinities := boostWithMagicFormula(gPAffinities, uPAffinities, options.Limit)

	var selectedPIDs []int64
	for _, affinity := range processedAffinities {
		selectedPIDs = append(selectedPIDs, affinity.ProductID)
	}

	products, err := uc.productRepo.GetByIDList(ctx, selectedPIDs)

	var productRecoms []pRecomUCase.ProductRecom
	for _, product := range products {
		affinityArrIndex := sliceUtils.IndexOfInt64(selectedPIDs, product.ID)

		productRecom := pRecomUCase.ProductRecom{
			Product:  product,
			Affinity: processedAffinities[affinityArrIndex].AffinityScore,
		}
		productRecoms = append(productRecoms, productRecom)
	}

	recommendation.SetData(productRecoms)
	return
}

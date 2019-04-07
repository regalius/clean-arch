package genderproductaffinity

import (
	"context"

	gPAModel "github.com/regalius/clean-arch/sample-app/internal/model/gender-product-affinity"
)

// Repository interface that represents all methods that can be used on gender product affinity ops
type Repository interface {
	GetGenderProductAffinitiesByPID(ctx context.Context, productID int64) (gPAffinity gPAModel.GenderProductAffinity, err error)
	GetGenderProductAffinitiesByGender(ctx context.Context, gender float32, threshold float32, options GetGenderProductAffinitiesByGenderOptions) (gPAffinities []gPAModel.GenderProductAffinity, err error)
}

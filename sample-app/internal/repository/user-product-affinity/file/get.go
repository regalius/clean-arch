package file

import (
	"context"

	uPAModel "github.com/regalius/clean-arch/sample-app/internal/model/user-product-affinity"
)

func (r fileUserProductAffinityRepository) GetUserProductAffinityByID(ctx context.Context, userID int64, productID int64) (uPAffinity uPAModel.UserProductAffinity, err error) {
	return
}

func (r fileUserProductAffinityRepository) GetUserProductAffinitiesByUserID(ctx context.Context, userID int64) (uPAffinities []uPAModel.UserProductAffinity, err error) {
	return
}

func (r fileUserProductAffinityRepository) GetUserProductAffinitiesByProductID(ctx context.Context, productID int64) (uPAffinities []uPAModel.UserProductAffinity, err error) {
	return
}

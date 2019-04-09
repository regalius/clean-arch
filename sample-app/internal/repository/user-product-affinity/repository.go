package userproductaffinity

import (
	"context"

	uPAModel "github.com/regalius/clean-arch/sample-app/internal/model/user-product-affinity"
)

// Repository interface that represents all methods that can be used on user product affinity ops
type Repository interface {
	GetUserProductAffinityByID(ctx context.Context, userID int64, productID int64) (uPAffinity uPAModel.UserProductAffinity, err error)
	GetUserProductAffinitiesByUserID(ctx context.Context, userID int64, options GetUserProductAffinitiesByUserIDOptions) (uPAffinities []uPAModel.UserProductAffinity, err error)
	GetUserProductAffinitiesByProductID(ctx context.Context, productID int64, options GetUserProductAffinitiesByProductIDOptions) (uPAffinities []uPAModel.UserProductAffinity, err error)
}

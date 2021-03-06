package file

import (
	"context"
	"sort"

	uPAModel "github.com/regalius/clean-arch/sample-app/internal/model/user-product-affinity"
	uPARepo "github.com/regalius/clean-arch/sample-app/internal/repository/user-product-affinity"
)

// GetUserProductAffinityByID get user product affinity by userID and productID
func (r fileUserProductAffinityRepository) GetUserProductAffinityByID(ctx context.Context, userID int64, productID int64) (uPAffinity uPAModel.UserProductAffinity, err error) {
	if r.buffer == nil || len(r.buffer) <= 0 {
		return
	}
	for _, affinity := range r.buffer {
		if affinity.UserID == userID && affinity.ProductID == productID {
			uPAffinity = affinity
			break
		}
	}

	return
}

// GetUserProductAffinitiesByUserID get user's all user product affinities by userID
// Options:
// Limit - default -1
func (r fileUserProductAffinityRepository) GetUserProductAffinitiesByUserID(ctx context.Context, userID int64, options uPARepo.GetUserProductAffinitiesByUserIDOptions) (uPAffinities []uPAModel.UserProductAffinity, err error) {
	if r.buffer == nil || len(r.buffer) <= 0 {
		return
	}

	if options.Limit == 0 {
		options.Limit = uPARepo.DefaultGetUserProductAffinitiesByUserIDOptions.Limit
	}

	for _, affinity := range r.buffer {
		if affinity.UserID == userID {
			uPAffinities = append(uPAffinities, affinity)
		}
	}

	if options.Limit > 0 {
		sort.Slice(uPAffinities, func(i, j int) bool {
			return uPAffinities[i].AffinityScore > uPAffinities[j].AffinityScore
		})
		var max int
		if len(uPAffinities) < options.Limit {
			max = len(uPAffinities)
		} else {
			max = options.Limit
		}
		uPAffinities = uPAffinities[0:max]
	}

	return
}

// GetUserProductAffinitiesByProductID get product's all user product affinities by productID
// Options:
// Limit - NOT YET IMPLEMENTED
func (r fileUserProductAffinityRepository) GetUserProductAffinitiesByProductID(ctx context.Context, productID int64, options uPARepo.GetUserProductAffinitiesByProductIDOptions) (uPAffinities []uPAModel.UserProductAffinity, err error) {
	if r.buffer == nil || len(r.buffer) <= 0 {
		return
	}

	for _, affinity := range r.buffer {
		if affinity.ProductID == productID {
			uPAffinities = append(uPAffinities, affinity)
		}
	}

	return
}

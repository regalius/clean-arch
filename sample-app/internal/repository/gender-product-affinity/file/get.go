package file

import (
	"context"
	"math"

	gPAModel "github.com/regalius/clean-arch/sample-app/internal/model/gender-product-affinity"
	gPARepo "github.com/regalius/clean-arch/sample-app/internal/repository/gender-product-affinity"
)

// GetGenderProductAffinitiesByPID get product's gender product affinity by ProductID
func (r fileGenderProductAffinityRepository) GetGenderProductAffinitiesByPID(ctx context.Context, productID int64) (gPAffinity gPAModel.GenderProductAffinity, err error) {
	if r.buffer == nil || len(r.buffer) <= 0 {
		return
	}

	for _, affinity := range r.buffer {
		if affinity.ProductID == productID {
			gPAffinity = affinity
			break
		}
	}

	return
}

// GetGenderProductAffinitiesByGender get closest gender product affinity given gender degree and threshold
func (r fileGenderProductAffinityRepository) GetGenderProductAffinitiesByGender(ctx context.Context, gender float32, threshold float32, options gPARepo.GetGenderProductAffinitiesByGenderOptions) (gPAffinities []gPAModel.GenderProductAffinity, err error) {
	if r.buffer == nil || len(r.buffer) <= 0 {
		return
	}
	if options.Limit == 0 {
		options.Limit = gPARepo.DefaultGetGenderProductAffinitiesByGenderOptions.Limit
	}
	var (
		maxDistance float64
		distances   []float64
	)
	threshold64 := float64(threshold)

	for i, affinity := range r.buffer {
		distance := math.Abs(float64(affinity.GenderAffinity - gender))
		if distance <= threshold64 {
			if len(gPAffinities) < options.Limit {
				if i == 0 || maxDistance <= distance {
					maxDistance = distance
				}
				gPAffinities = append(gPAffinities, affinity)
				distances = append(distances, distance)
			} else if distance < maxDistance {
				nextMaxDistanceCandidate := float64(0)

				for index := len(distances) - 1; index >= 0; index-- {
					selectedDistance := distances[index]
					if selectedDistance >= maxDistance {
						distances[index] = distance
						gPAffinities[index] = affinity
					} else if selectedDistance >= nextMaxDistanceCandidate {
						nextMaxDistanceCandidate = selectedDistance
					}
				}
				maxDistance = nextMaxDistanceCandidate
			}
		}
	}
	return
}

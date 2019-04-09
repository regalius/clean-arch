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
func (r fileGenderProductAffinityRepository) GetGenderProductAffinitiesByGender(ctx context.Context, gender float32, threshold float32, options gPARepo.GetGenderProductAffinitiesByGenderOptions) (gPAffinities []gPAModel.GenderProductAffinityWithScore, err error) {
	if r.buffer == nil || len(r.buffer) <= 0 {
		return
	}
	if options.Limit == 0 {
		options.Limit = gPARepo.DefaultGetGenderProductAffinitiesByGenderOptions.Limit
	}
	var (
		maxScore float64
		scores   []float64
	)
	threshold64 := float64(threshold)

	for i, affinity := range r.buffer {
		score := math.Abs(float64(affinity.GenderAffinity - gender))
		if score <= threshold64 {
			if len(gPAffinities) < options.Limit {
				if i == 0 || maxScore <= score {
					maxScore = score
				}
				gPAffinities = append(gPAffinities, gPAModel.GenderProductAffinityWithScore{
					GenderProductAffinity: affinity,
					AffinityScore:         float32(score),
				})
				scores = append(scores, score)
			} else if score < maxScore {
				nextMaxScoreCandidate := float64(0)

				for index := len(scores) - 1; index >= 0; index-- {
					selectedDistance := scores[index]
					if selectedDistance >= maxScore {
						scores[index] = score
						gPAffinities[index] = gPAModel.GenderProductAffinityWithScore{
							GenderProductAffinity: affinity,
							AffinityScore:         float32(score),
						}
					} else if selectedDistance >= nextMaxScoreCandidate {
						nextMaxScoreCandidate = selectedDistance
					}
				}
				maxScore = nextMaxScoreCandidate
			}
		}
	}
	return
}

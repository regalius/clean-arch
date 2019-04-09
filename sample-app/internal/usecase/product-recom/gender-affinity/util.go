package genderaffinity

import (
	"sort"

	gPAModel "github.com/regalius/clean-arch/sample-app/internal/model/gender-product-affinity"
)

// boostWithMagicFormula in gender affinity magic formula's we only care about gender product affinity order
func boostWithMagicFormula(affinities []gPAModel.GenderProductAffinityWithScore, limit int) (result []gPAModel.GenderProductAffinityWithScore) {
	return sort.Slice(affinities, func(i, j int) {
		return affinities[i].AffinityScore < affinities[j].AffinityScore
	})[0:limit]
}

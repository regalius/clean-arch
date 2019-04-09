package useraffinity

import (
	"sort"

	uPAModel "github.com/regalius/clean-arch/sample-app/internal/model/user-product-affinity"
)

// boostWithMagicFormula in user affinity magic formula's we only care about user product affinity order
func boostWithMagicFormula(affinities []uPAModel.UserProductAffinity, limit int) (result []uPAModel.UserProductAffinity) {
	return sort.Slice(affinities, func(i, j int) {
		return affinities[i].AffinityScore < affinities[j].AffinityScore
	})[0:limit]
}

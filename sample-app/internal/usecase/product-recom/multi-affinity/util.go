package multiaffinity

import (
	"sort"

	gPAModel "github.com/regalius/clean-arch/sample-app/internal/model/gender-product-affinity"
	uPAModel "github.com/regalius/clean-arch/sample-app/internal/model/user-product-affinity"
)

// boostWithMagicFormula in multi affinity magic formula's combine gender affinity and user affinity with defined multiplier
func boostWithMagicFormula(gPAffinities []gPAModel.GenderProductAffinityWithScore, uPAffinities []uPAModel.UserProductAffinity, limit int) []combinedAffinities {
	affinityMap := make(map[int64]float32)
	for _, gPAffinity := range gPAffinities {
		currAffinity := (gPAffinity.AffinityScore * gPAffinityMultiplier) / (gPAffinityMultiplier + uPAffinityMultiplier)
		affinityMap[gPAffinity.ProductID] = currAffinity
	}

	for _, uPAffinity := range uPAffinities {
		currAffinity := (uPAffinity.AffinityScore * uPAffinityMultiplier) / (gPAffinityMultiplier + uPAffinityMultiplier)
		if affinityMap[uPAffinity.ProductID] != 0 {
			affinityMap[uPAffinity.ProductID] += currAffinity
		}
	}

	var affinities []combinedAffinities
	for productID, affinity := range affinityMap {
		affinities = append(affinities, combinedAffinities{
			ProductID:     productID,
			AffinityScore: affinity,
		})
	}

	sort.Slice(affinities, func(i, j int) bool {
		return affinities[i].AffinityScore > affinities[j].AffinityScore
	})

	var max int
	if len(affinities) < limit {
		max = len(affinities)
	} else {
		max = limit
	}

	return affinities[0:max]
}

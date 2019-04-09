package genderproductaffinity

type (
	// GenderProductAffinity model that represents gender and product affinity
	GenderProductAffinity struct {
		ID             int64
		ProductID      int64
		GenderAffinity float32
	}
	// GenderProductAffinityWithScore model that represents gender and product affinity and its computed score
	GenderProductAffinityWithScore struct {
		GenderProductAffinity
		AffinityScore float32
	}
)

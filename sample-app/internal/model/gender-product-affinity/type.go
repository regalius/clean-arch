package genderproductaffinity

type (
	// GenderProductAffinity model that represents gender and product affinity
	GenderProductAffinity struct {
		ID             int64
		ProductID      int64
		GenderAffinity float32
	}
)

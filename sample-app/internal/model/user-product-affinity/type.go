package userproductaffinity

type (
	// UserProductAffinity model that describes affinity between a user and a product
	UserProductAffinity struct {
		ID            int64
		UserID        int64
		ProductID     int64
		AffinityScore float32
	}
)

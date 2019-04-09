package productrecom

import (
	pModel "github.com/regalius/clean-arch/sample-app/internal/model/product"
)

type (
	// SingleUserResult result for a user's product recommendation
	SingleUserResult struct {
		Meta SingleUserResultMeta
		Data []ProductRecom
	}
	// ProductRecom recommended product
	ProductRecom struct {
		pModel.Product
		Affinity float32
	}
	// SingleUserResultMeta meta data for single user's product recommendation result
	SingleUserResultMeta struct {
		OverallAffinities float32
		Threshold         float32
	}
	// GetRecommendationByUserIDOptions available options for GetRecommendationByUserID method
	GetRecommendationByUserIDOptions struct {
		Limit int
	}
)

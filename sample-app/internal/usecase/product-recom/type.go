package productrecom

import (
	pModel "github.com/regalius/clean-arch/sample-app/pkg/model/product"
)

type (
	// SingleUserResult result for a user's product recommendation
	SingleUserResult struct {
		Meta SingleUserResultMeta
		Data []ProductRecom
	}
	// ProductRecom recommended product
	ProductRecom struct {
		Product    pModel.Product
		Confidence float32
	}
	// SingleUserResultMeta meta data for single user's product recommendation result
	SingleUserResultMeta struct {
		OverallConfidence float32
		Threshold         float32
	}
	// GetRecommendationForUserIDOptions available options for GetRecommendationForUserID method
	GetRecommendationForUserIDOptions struct {
		Limit int
	}
)

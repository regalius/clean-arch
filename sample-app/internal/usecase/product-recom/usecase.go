package productrecom

import (
	"context"
)

// Usecase interface that represents product recom usecase methods
type Usecase interface {
	GetRecommendationByUserID(ctx context.Context, userID int64, options GetRecommendationByUserIDOptions) (recommendation SingleUserResult, err error)
}

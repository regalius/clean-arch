package productrecom

import (
	"context"
)

// Usecase interface that represents product recom usecase methods
type Usecase interface {
	GetRecommendationForUserID(ctx context.Context, userID int64, options GetRecommendationForUserIDOptions) (recommendation SingleUserResult, err error)
}

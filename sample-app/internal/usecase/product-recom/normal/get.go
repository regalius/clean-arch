package normal

import (
	pRecomUCase "github.com/regalius/clean-arch/sample-app/internal/usecase/product-recom"
)

func (uc productRecomUsecase) GetRecommendationForUserID(userID int64, options pRecomUCase.GetRecommendationForUserIDoptions)
	(recommendation pRecomUCase.SingleUserResult, err error) {

}

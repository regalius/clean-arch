package http

import (
	r "github.com/julienschmidt/httprouter"
	pRecomUcase "github.com/regalius/clean-arch/sample-app/internal/usecase/product-recom"
)

type httpProductRecomDelivery struct {
	router            *r.Router
	pRecomGenderUcase pRecomUcase.Usecase
	pRecomUserUcase   pRecomUcase.Usecase
	pRecomMultiUcase  pRecomUcase.Usecase
}

func NewHTTPProductRecomDelivery(router *r.Router,
	pRecomGenderUcase pRecomUcase.Usecase,
	pRecomUserUcase pRecomUcase.Usecase,
	pRecomMultiUcase pRecomUcase.Usecase) httpProductRecomDelivery {

	handler := &httpProductRecomDelivery{
		router,
		pRecomGenderUcase,
		pRecomUserUcase,
		pRecomMultiUcase,
	}

	router.GET("/recommendation/product/multi/:userid", handler.getProductRecommendationMulti)
	router.GET("/recommendation/product/gender/:userid", handler.getProductRecommendationGender)
	router.GET("/recommendation/product/user/:userid", handler.getProductRecommendationUser)

	return *handler
}

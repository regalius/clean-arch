package http

import (
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	httpUtils "github.com/regalius/clean-arch/sample-app/common/http"
	pRecomUCase "github.com/regalius/clean-arch/sample-app/internal/usecase/product-recom"
)

func (d httpProductRecomDelivery) getProductRecommendationMulti(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID, err := getUserIDFromParams(ps)
	if err != nil {
		log.Println(err)
		httpUtils.WriteStatusCode(w, http.StatusBadRequest, "")
		return
	}

	ctx := r.Context()
	limit := 0

	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			log.Println(err)
			limit = 0
		}
	}
	recomm, err := d.pRecomMultiUcase.GetRecommendationByUserID(ctx, userID, pRecomUCase.GetRecommendationByUserIDOptions{
		Limit: limit,
	})
	response := GetProductRecomResponse{
		Data: []ProductResponse{},
	}
	if err != nil {
		log.Println(err)
		httpUtils.WriteJSON(w, response, http.StatusInternalServerError)
		return
	}

	response = mapProductRecomResultToResponse(recomm, r.URL.Host)

	httpUtils.WriteJSON(w, response, http.StatusOK)
}

func (d httpProductRecomDelivery) getProductRecommendationGender(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID, err := getUserIDFromParams(ps)
	if err != nil {
		log.Println(err)
		httpUtils.WriteStatusCode(w, http.StatusBadRequest, "")
		return
	}

	ctx := r.Context()
	limit := 0

	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			log.Println(err)
			limit = 0
		}
	}

	recomm, err := d.pRecomGenderUcase.GetRecommendationByUserID(ctx, userID, pRecomUCase.GetRecommendationByUserIDOptions{
		Limit: limit,
	})
	var response GetProductRecomResponse
	if err != nil {
		log.Println(err)
		httpUtils.WriteJSON(w, response, http.StatusInternalServerError)
		return
	}

	response = mapProductRecomResultToResponse(recomm, r.URL.Host)

	httpUtils.WriteJSON(w, response, http.StatusOK)
}

func (d httpProductRecomDelivery) getProductRecommendationUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID, err := getUserIDFromParams(ps)
	if err != nil {
		log.Println(err)
		httpUtils.WriteStatusCode(w, http.StatusBadRequest, "")
		return
	}

	ctx := r.Context()
	limit := 0

	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			log.Println(err)
			limit = 0
		}
	}

	recomm, err := d.pRecomUserUcase.GetRecommendationByUserID(ctx, userID, pRecomUCase.GetRecommendationByUserIDOptions{
		Limit: limit,
	})
	var response GetProductRecomResponse
	if err != nil {
		log.Println(err)
		httpUtils.WriteJSON(w, response, http.StatusInternalServerError)
		return
	}

	response = mapProductRecomResultToResponse(recomm, r.URL.Host)

	httpUtils.WriteJSON(w, response, http.StatusOK)
}

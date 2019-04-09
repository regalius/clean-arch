package http

import (
	"errors"
	"strconv"

	"github.com/julienschmidt/httprouter"
	pRecomUCase "github.com/regalius/clean-arch/sample-app/internal/usecase/product-recom"
)

func getUserIDFromParams(ps httprouter.Params) (userID int64, err error) {
	userIDStr := ps.ByName("userid")
	if userIDStr == "" {
		err = errors.New("User ID is empty")
		return
	}
	userID, err = strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return
	}
	return
}

func mapProductRecomResultToResponse(recomm pRecomUCase.SingleUserResult, rootURL string) (response GetProductRecomResponse) {
	response.Meta.OverallAffinities = recomm.Meta.OverallAffinities
	response.Data = []ProductResponse{}
	for _, product := range recomm.Data {
		productItem := ProductResponse{
			ID:              product.ID,
			Price:           product.Price,
			ShopID:          product.ShopID,
			TimeCreated:     product.TimeCreated,
			TimeLastUpdated: product.TimeLastUpdated,
			CreatedByID:     product.CreatedByID,
			URL:             product.GetURL(rootURL),
			Name:            product.Name,
			AffinityScore:   product.Affinity,
		}

		for _, image := range product.Images {
			productItem.Images = append(productItem.Images, image.GetFullImagePath())
		}

		response.Data = append(response.Data, productItem)
	}

	return
}

package main

// Here, we're trying to import pkg from the sample app
// as other app's perspective to reuse sample-app's internal code
// eventhough using other app's internal code should be treated as anti-pattern

import (
	"context"
	"log"
	"os"

	pRecom "github.com/regalius/clean-arch/sample-app/pkg/repository/product"
)

func main() {
	// Trying struct method
	pd := pRecom.Product{
		ID:   1,
		Name: "product 1",
	}
	pd.GetURL("http://sample.com")

	// Trying repository
	productFile, err := os.Open("../sample-app/configs/var/data/product.json")
	if err != nil {
		log.Println("[PersoEngine] Failed to open product.json", err)
	}
	pRecomProduct := pRecom.NewProductFileRepository(productFile)
	ctx := context.Background()
	product, err := pRecomProduct.GetByIDRange(ctx, 3, 5)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(product)

	// Trying mock
	mock := pRecom.ProductRepositoryMock{
		GetByIDFunc: func(ctx context.Context, id int64) (p pRecom.Product, err error) {
			return
		},
	}
	mock.GetByIDFunc(ctx, 10)
}

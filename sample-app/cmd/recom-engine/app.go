package main

import (
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"

	gPAFileRepo "github.com/regalius/clean-arch/sample-app/internal/repository/gender-product-affinity/file"
	pFileRepo "github.com/regalius/clean-arch/sample-app/internal/repository/product/file"
	uPAFileRepo "github.com/regalius/clean-arch/sample-app/internal/repository/user-product-affinity/file"
	uFileRepo "github.com/regalius/clean-arch/sample-app/internal/repository/user/file"

	pRecomGenderUcase "github.com/regalius/clean-arch/sample-app/internal/usecase/product-recom/gender-affinity"
	pRecomMultiUcase "github.com/regalius/clean-arch/sample-app/internal/usecase/product-recom/multi-affinity"
	pRecomUserUcase "github.com/regalius/clean-arch/sample-app/internal/usecase/product-recom/user-affinity"

	pRecomHTTPDelivery "github.com/regalius/clean-arch/sample-app/internal/delivery/product-recom/http"
)

func main() {
	openFiles := make(map[string]*os.File)
	var err error
	for context, path := range fileList {
		openFiles[context], err = os.Open(path)
		if err != nil {
			log.Fatal("[RecomEngine] Failed to open", path, err)
		}
	}

	router := httprouter.New()

	pRepo := pFileRepo.NewFileProductRepository(openFiles["product"])
	uRepo := uFileRepo.NewFileUserRepository(openFiles["user"])
	uPARepo := uPAFileRepo.NewFileUserProductAffinityRepository(openFiles["userPA"])
	gPARepo := gPAFileRepo.NewFileGenderProductAffinityRepository(openFiles["genderPA"])

	pRGenderUcase := pRecomGenderUcase.NewProductRecomGenderUsecase(uRepo, pRepo, gPARepo)
	pRUserUcase := pRecomUserUcase.NewProductRecomUserUsecase(uRepo, pRepo, uPARepo)
	pRMultiUcase := pRecomMultiUcase.NewProductRecomMultiUsecase(uRepo, pRepo, uPARepo, gPARepo)

	pRecomHTTPDelivery.NewHTTPProductRecomDelivery(router, pRGenderUcase, pRUserUcase, pRMultiUcase)

	log.Fatal(http.ListenAndServe(":8080", router))
}

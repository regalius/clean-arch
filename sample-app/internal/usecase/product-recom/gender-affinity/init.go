package genderaffinity

import (
	gPARepo "github.com/regalius/clean-arch/sample-app/internal/repository/gender-product-affinity"
	pRepo "github.com/regalius/clean-arch/sample-app/internal/repository/product"
	uRepo "github.com/regalius/clean-arch/sample-app/internal/repository/user"
	pRecomUCase "github.com/regalius/clean-arch/sample-app/internal/usecase/product-recom"
)

type productRecomGenderUsecase struct {
	userRepo            uRepo.Repository
	productRepo         pRepo.Repository
	genderPAffinityRepo gPARepo.Repository
}

// NewProductRecomGenderUsecase create new product recom default usecase
func NewProductRecomGenderUsecase(userRepo uRepo.Repository, productRepo pRepo.Repository, genderPAffinityRepo gPARepo.Repository) pRecomUCase.Usecase {
	return productRecomGenderUsecase{
		userRepo,
		productRepo,
		genderPAffinityRepo,
	}
}

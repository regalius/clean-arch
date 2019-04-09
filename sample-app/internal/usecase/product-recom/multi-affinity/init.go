package multiaffinity

import (
	gPARepo "github.com/regalius/clean-arch/sample-app/internal/repository/gender-product-affinity"
	pRepo "github.com/regalius/clean-arch/sample-app/internal/repository/product"
	uRepo "github.com/regalius/clean-arch/sample-app/internal/repository/user"
	uPARepo "github.com/regalius/clean-arch/sample-app/internal/repository/user-product-affinity"
	pRecomUCase "github.com/regalius/clean-arch/sample-app/internal/usecase/product-recom"
)

type productRecomMultiUsecase struct {
	userRepo            uRepo.Repository
	productRepo         pRepo.Repository
	userPAffinityRepo   uPARepo.Repository
	genderPAffinityRepo gPARepo.Repository
}

// NewProductRecomMultiUsecase create new product recom default usecase
func NewProductRecomMultiUsecase(userRepo uRepo.Repository, productRepo pRepo.Repository, userPAffinityRepo uPARepo.Repository, genderPAffinityRepo gPARepo.Repository) pRecomUCase.Usecase {
	return productRecomMultiUsecase{
		userRepo,
		productRepo,
		userPAffinityRepo,
		genderPAffinityRepo,
	}
}

package multiaffinity

import (
	gPARepo "github.com/regalius/clean-arch/sample-app/internal/repository/gender-product-affinity"
	uPARepo "github.com/regalius/clean-arch/sample-app/internal/repository/user-product-affinity"
	pRecomUCase "github.com/regalius/clean-arch/sample-app/internal/usecase/product-recom"
	pRepo "github.com/regalius/clean-arch/sample-app/internal/repository/product"
	uRepo "github.com/regalius/clean-arch/sample-app/internal/repository/user"
)

type productRecomDefaultUsecase struct {
	userRepo            uRepo.Repository
	productRepo         pRepo.Repository
	userPAffinityRepo   uPARepo.Repository
	genderPAffinityRepo gPARepo.Repository
}

// NewProductRecomDefaultUsecase create new product recom default usecase
func NewProductRecomDefaultUsecase(userRepo uRepo.Repository, productRepo pRepo.Repository, userPAffinityRepo uPARepo.Repository, genderPAffinityRepo gPARepo.Repository) pRecomUCase.Usecase {
	return productRecomUsecase{
		userRepo,
		productRepo,
		userPAffinityRepo,
		genderPAffinityRepo,
	}
}

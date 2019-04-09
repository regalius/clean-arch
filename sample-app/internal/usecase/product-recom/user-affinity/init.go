package useraffinity

import (
	pRepo "github.com/regalius/clean-arch/sample-app/internal/repository/product"
	uRepo "github.com/regalius/clean-arch/sample-app/internal/repository/user"
	uPARepo "github.com/regalius/clean-arch/sample-app/internal/repository/user-product-affinity"
	pRecomUCase "github.com/regalius/clean-arch/sample-app/internal/usecase/product-recom"
)

type productRecomUserUsecase struct {
	userRepo          uRepo.Repository
	productRepo       pRepo.Repository
	userPAffinityRepo uPARepo.Repository
}

// NewProductRecomUserUsecase create new product recom default usecase
func NewProductRecomUserUsecase(userRepo uRepo.Repository, productRepo pRepo.Repository, userPAffinityRepo uPARepo.Repository) pRecomUCase.Usecase {
	return productRecomUserUsecase{
		userRepo,
		productRepo,
		userPAffinityRepo,
	}
}

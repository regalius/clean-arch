package normal

import (
	gPARepo "github.com/regalius/clean-arch/sample-app/internal/repository/gender-product-affinity"
	uPARepo "github.com/regalius/clean-arch/sample-app/internal/repository/user-product-affinity"
	pRecomUCase "github.com/regalius/clean-arch/sample-app/internal/usecase/product-recom"
	pRepo "github.com/regalius/clean-arch/sample-app/pkg/repository/product"
)

type productRecomUsecase struct {
	productRepo        pRepo.Repository
	userPAffinityRepo  uPARepo.Repository
	genderPAffiityRepo gPARepo.Repository
}

func NewProductRecomUsecase(productRepo pRepo.Repository, userPAffinityRepo uPARepo.Repository, genderPAffinityRepo gPARepo.Repository) pRecomUCase.Usecase {
	return productRecomUsecase{
		productRepo,
		userPAffinityRepo,
		genderPAffiityRepo,
	}
}

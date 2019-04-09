package product

import (
	"io"

	pModel "github.com/regalius/clean-arch/sample-app/internal/model/product"
	pRepo "github.com/regalius/clean-arch/sample-app/internal/repository/product"
	pFileRepo "github.com/regalius/clean-arch/sample-app/internal/repository/product/file"
)

type Product = pModel.Product
type ProductImage = pModel.Image
type ProductRepository = pRepo.Repository
type ProductRepositoryMock = pRepo.RepositoryMock

func NewProductFileRepository(file io.Reader) ProductRepository {
	return pFileRepo.NewFileProductRepository(file)
}

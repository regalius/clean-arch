package product

import (
	"io"

	pRepo "github.com/regalius/clean-arch/sample-app/internal/repository/product"
	pFileRepo "github.com/regalius/clean-arch/sample-app/internal/repository/product/file"
)

// Repository interface for product repository
type Repository = pRepo.Repository

// RepositoryMock interface mock for ProductRepository
type RepositoryMock = pRepo.RepositoryMock

// NewProductFileRepository create new file-based product repository
func NewProductFileRepository(file io.Reader) Repository {
	return pFileRepo.NewFileProductRepository(file)
}

package file

import (
	"io"
	"log"

	pModel "github.com/regalius/clean-arch/sample-app/internal/model/product"
	pRepo "github.com/regalius/clean-arch/sample-app/internal/repository/product"
)

type fileProductRepository struct {
	file   io.Reader
	buffer []pModel.Product
	lock   bool
}

// NewFileProductRepository Initializes New product repository with File adapter
func NewFileProductRepository(file io.Reader) pRepo.Repository {
	products, err := fileToProducts(file)
	if err != nil {
		log.Panic("[Repo/Product/File] Error initializing Repository", err)
		return nil
	}

	return fileProductRepository{
		file:   file,
		buffer: products,
		lock:   false,
	}
}

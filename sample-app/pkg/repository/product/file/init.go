package file

import (
	"io"
	"log"

	pModel "github.com/regalius/clean-arch/sample-app/pkg/model/product"
	pRepo "github.com/regalius/clean-arch/sample-app/pkg/repository/product"
)

type fileProductRepository struct {
	file   io.Reader
	buffer []pModel.Product
	lock   bool
}

// NewFileProductRepository Initializes New product repository with File adapter
func NewFileProductRepository(file io.Reader) pRepo.Repository {
	products, err := fileToProducts(r.file)
	if err != nil {
		log.Panic("[Repo/Product/File] Error initializing Repository", err)
		return nil
	}

	return fileProductRepository{
		file: file,
		buffer: products,
		lock:   false,
	}
}

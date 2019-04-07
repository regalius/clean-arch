package file

import (
	"io"

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
	var products []pModel.Product
	products, err = fileToProducts(r.file)
	if err != nil {
		log.Panic("[Repo/Product] Error initializing Repository", err)
		return
	}

	return fileProductRepository{
		file,
		buffer: products,
		lock: false,
	}
}

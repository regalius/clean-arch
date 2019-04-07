package file

import (
	"context"
	"log"

	"errors"

	pModel "github.com/regalius/clean-arch/sample-app/pkg/model/product"
	pRepo "github.com/regalius/clean-arch/sample-app/pkg/repository/product"
)

// GetByID get a product from file given its ID
func (r fileProductRepository) GetByID(ctx context.Context, id int64) (result pModel.Product, err error) {
	if r.buffer == nil || len(r.buffer) <= 0 {
		return
	}
	for _, product := range r.buffer {
		if product.ID == id {
			result = product
			break
		}
	}
	return
}

// GetByIDRange get products from file given its min and max ID
func (r fileProductRepository) GetByIDRange(ctx context.Context, minID int64, maxID int64) (result []pModel.Product, err error) {
	if maxID < minID {
		err = errors.New("[Repo/Product/File] MaxID is smaller than MinID")
		return
	}

	if r.buffer == nil || len(r.buffer) <= 0 {
		return
	}

	for _, product := range r.buffer {
		if product.ID > maxID {
			break
		}
		if product.ID >= minID && product.ID <= maxID {
			result = append(result, product)
		}
	}

	return
}

// GetByFilter NOT YET IMPLEMENTED
func (r fileProductRepository) GetByFilter(ctx context.Context, filter pRepo.Filter) (result []pModel.Product, err error) {
	log.Panic("[Repo/Product/File] GetByFilter is not yet implemented")
	return
}

package file

import (
	"context"
	"log"

	"errors"

	sliceUtils "github.com/regalius/clean-arch/sample-app/common/slice"
	pModel "github.com/regalius/clean-arch/sample-app/internal/model/product"
	pRepo "github.com/regalius/clean-arch/sample-app/internal/repository/product"
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

// GetByIDRange get products from file given its min and max ID
func (r fileProductRepository) GetByIDList(ctx context.Context, ids []int64) (result []pModel.Product, err error) {
	if len(ids) <= 0 {
		return
	}
	if r.buffer == nil || len(r.buffer) <= 0 {
		return
	}
	result = make([]pModel.Product, len(ids))

	for _, product := range r.buffer {
		if index := sliceUtils.IndexOfInt64(ids, product.ID); index != -1 {
			result[index] = product
		}
	}

	return
}

// GetByFilter NOT IMPLEMENTED
func (r fileProductRepository) GetByFilter(ctx context.Context, filter pRepo.Filter) (result []pModel.Product, err error) {
	log.Panic("[Repo/Product/File] GetByFilter is not implemented")
	return
}

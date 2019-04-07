package file

import (
	"context"
	"log"

	pModel "github.com/regalius/clean-arch/sample-app/pkg/model/product"
)

// UpdateProduct NOT YET IMPLEMENTED
func (r *fileProductRepository) UpdateProduct(ctx context.Context, newProduct pModel.Product) (id int, err error) {
	log.Panic("[Repo/Product/File] UpdateProduct is not yet implemented")
	return
}

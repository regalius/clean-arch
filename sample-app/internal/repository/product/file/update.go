package file

import (
	"context"
	"log"

	pModel "github.com/regalius/clean-arch/sample-app/internal/model/product"
)

// func (r *fileProductRepository) updateProduct(ctx context.Context, newProduct pModel.Product) (id int, err error) {
// 	log.Panic("[Repo/Product/File] UpdateProduct is not implemented")
// 	return
// }

// UpdateProduct NOT IMPLEMENTED
func (r fileProductRepository) UpdateProduct(ctx context.Context, newProduct pModel.Product) (id int, err error) {
	log.Panic("[Repo/Product/File] UpdateProduct is not implemented")
	// r.updateProduct(ctx, newProduct)
	return
}

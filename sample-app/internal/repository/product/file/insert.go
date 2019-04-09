package file

import (
	"context"
	"log"

	pModel "github.com/regalius/clean-arch/sample-app/internal/model/product"
)

// func (r *fileProductRepository) InsertProduct(ctx context.Context, p pModel.Product) (id int, err error) {
// 	log.Panic("[Repo/Product/File] InsertProduct is not implemented")
// 	return
// }

// InsertProduct NOT IMPLEMENTED
func (r fileProductRepository) InsertProduct(ctx context.Context, p pModel.Product) (id int, err error) {
	// r.insertProduct(ctx, p)
	log.Panic("[Repo/Product/File] InsertProduct is not implemented")
	return
}

package file

import (
	"context"
	"log"

	pModel "github.com/regalius/clean-arch/sample-app/pkg/model/product"
)

// InsertProduct NOT YET IMPLEMENTED
func (r *fileProductRepository) InsertProduct(ctx context.Context, product pModel.Product) (id int, err error) {
	log.Panic("[Repo/Product/File] InsertProduct is not yet implemented")
	return
}

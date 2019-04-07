package product

import (
	"context"

	pModel "github.com/regalius/clean-arch/sample-app/pkg/model/product"
)

// Repository interface that represents product repository methods
type Repository interface {
	GetByID(ctx context.Context, id int64) (products pModel.Product, err error)
	GetByIDRange(ctx context.Context, minID int64, maxID int64) (products []pModel.Product, err error)
	GetByFilter(ctx context.Context, filter Filter) (products []pModel.Product, err error)
	UpdateProduct(ctx context.Context, newProduct pModel.Product) (id int, err error)
	InsertProduct(ctx context.Context, product pModel.Product) (id int, err error)
}

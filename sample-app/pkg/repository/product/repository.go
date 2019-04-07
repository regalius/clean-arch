package product

import "context"

type ProductRepository interface {
	GetByID(ctx context.Context, id int64)
	GetByIDRange(ctx context.Context, minID int64, maxID int64)
	GetByFilter(ctx context.Context, filter Filter)
}

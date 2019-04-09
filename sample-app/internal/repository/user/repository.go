package user

import (
	"context"

	uModel "github.com/regalius/clean-arch/sample-app/internal/model/user"
)

// Repository interface that represents what can be done with user repository
type Repository interface {
	GetUserByID(ctx context.Context, id int64) (user uModel.User, err error)
}

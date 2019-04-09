package file

import (
	"context"

	uModel "github.com/regalius/clean-arch/sample-app/internal/model/user"
)

// GetUserByID get user given its userID
func (r fileUserRepository) GetUserByID(ctx context.Context, id int64) (result uModel.User, err error) {
	if r.buffer == nil || len(r.buffer) <= 0 {
		return
	}

	for _, user := range r.buffer {
		if user.ID == id {
			result = user
			break
		}
	}
	return
}

package file

import (
	"io"
	"log"

	uModel "github.com/regalius/clean-arch/sample-app/internal/model/user"
	uRepo "github.com/regalius/clean-arch/sample-app/internal/repository/user"
)

type fileUserRepository struct {
	file   io.Reader
	buffer []uModel.User
	lock   bool
}

// NewFileUserRepository create new file user repository instance
func NewFileUserRepository(file io.Reader) uRepo.Repository {
	users, err := fileToUsers(file)
	if err != nil {
		log.Panic("[Repo/User/File] Error initializing Repository", err)
		return nil
	}

	return fileUserRepository{
		file:   file,
		buffer: users,
		lock:   false,
	}
}

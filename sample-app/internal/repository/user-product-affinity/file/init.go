package file

import (
	"io"
	"log"

	uPAModel "github.com/regalius/clean-arch/sample-app/internal/model/user-product-affinity"
	uPARepo "github.com/regalius/clean-arch/sample-app/internal/repository/user-product-affinity"
)

type fileUserProductAffinityRepository struct {
	file   io.Reader
	buffer []uPAModel.UserProductAffinity
	lock   bool
}

// NewFileUserProductAffinityRepository initialize new file repository for user product affinity
func NewFileUserProductAffinityRepository(file io.Reader) uPARepo.Repository {
	uPAffinities, err := fileToUserPAffinity(file)
	if err != nil {
		log.Panic("[Repo/UPAffinity/File] Error initializing Repository", err)
		return nil
	}

	return fileUserProductAffinityRepository{
		file:   file,
		buffer: uPAffinities,
		lock:   false,
	}
}

package file

import (
	"io"

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
	return fileUserProductAffinityRepository{
		file,
	}
}

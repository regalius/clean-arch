package file

import (
	"io"
	"log"

	gPAModel "github.com/regalius/clean-arch/sample-app/internal/model/gender-product-affinity"
	gPARepo "github.com/regalius/clean-arch/sample-app/internal/repository/gender-product-affinity"
)

type fileGenderProductAffinityRepository struct {
	file   io.Reader
	buffer []gPAModel.GenderProductAffinity
	lock   bool
}

// NewFileGenderProductAffinityRepository initialize new file repository for user product affinity
func NewFileGenderProductAffinityRepository(file io.Reader) gPARepo.Repository {
	uPAffinities, err := fileToGenderPAffinity(file)
	if err != nil {
		log.Panic("[Repo/GPAffinity/File] Error initializing Repository", err)
		return nil
	}

	return fileGenderProductAffinityRepository{
		file:   file,
		buffer: uPAffinities,
		lock:   false,
	}
}

package service

import (
	"go-gcs/features/images"

	"github.com/go-playground/validator/v10"
)

type imageService struct {
	imageData images.ImageData
	validate  *validator.Validate
}

// Create implements images.ImageService
func (*imageService) Create(input images.Core) error {
	panic("unimplemented")
}

// Delete implements images.ImageService
func (*imageService) Delete(data images.Core) error {
	panic("unimplemented")
}

func New(repo images.ImageData) images.ImageService {
	return &imageService{
		imageData: repo,
		validate:  validator.New(),
	}
}

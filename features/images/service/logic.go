package service

import (
	"go-gcs/app/storage"
	"go-gcs/features/images"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
)

type imageService struct {
	imageData images.ImageData
	validate  *validator.Validate
}

// Create implements images.ImageService
func (service *imageService) Create(input images.Core) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}

	input.ImageName = strconv.FormatInt(time.Now().UTC().UnixNano(), 10) + "_" + input.ImageName

	urlImage, errUpload := storage.GetStorageClinet().UploadFile(input.ImageFile, input.ImageName)
	if errUpload != nil {
		return errUpload
	}

	input.Url = urlImage

	errInsert := service.imageData.Insert(input)
	if errInsert != nil {
		return errInsert
	}

	return nil
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

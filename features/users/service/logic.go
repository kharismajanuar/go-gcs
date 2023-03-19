package service

import (
	"go-gcs/app/storage"
	"go-gcs/features/users"

	"github.com/go-playground/validator/v10"
	"github.com/lithammer/shortuuid/v4"
)

type userService struct {
	userData users.UserData
	validate *validator.Validate
}

// Create implements users.UserService
func (service *userService) Create(input users.Core) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}

	input.ImageName = shortuuid.New()

	urlImage, errUpload := storage.GetStorageClinet().UploadFile(input.ImageFile, input.ImageName)
	if errUpload != nil {
		return errUpload
	}

	input.DisplayImage = urlImage

	errInsert := service.userData.Insert(input)
	if errInsert != nil {
		return errInsert
	}

	return nil
}

func New(repo users.UserData) users.UserService {
	return &userService{
		userData: repo,
		validate: validator.New(),
	}
}

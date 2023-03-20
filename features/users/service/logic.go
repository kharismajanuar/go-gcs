package service

import (
	"go-gcs/app/storage"
	"go-gcs/features/users"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
)

type userService struct {
	userData users.UserData
	validate *validator.Validate
}

// DeleteAvatar implements users.UserService
func (service *userService) DeleteAvatar(data users.Core) error {
	errDeleteImage := storage.GetStorageClinet().DeleteFile(data.Avatar)
	if errDeleteImage != nil {
		return errDeleteImage
	}

	errDeleteAvatar := service.userData.DeleteAvatar(data)
	if errDeleteAvatar != nil {
		return errDeleteAvatar
	}

	return nil
}

// Create implements users.UserService
func (service *userService) Create(input users.Core) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}

	input.ImageName = strconv.FormatInt(time.Now().UTC().UnixNano(), 10) + "_" + input.ImageName

	urlImage, errUpload := storage.GetStorageClinet().UploadFile(input.ImageFile, input.ImageName)
	if errUpload != nil {
		return errUpload
	}

	input.Avatar = urlImage

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

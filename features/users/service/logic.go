package service

import (
	"go-gcs/features/users"

	"github.com/go-playground/validator/v10"
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

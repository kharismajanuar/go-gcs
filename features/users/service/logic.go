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
func (*userService) Create(input users.Core) error {
	panic("unimplemented")
}

func New(repo users.UserData) users.UserService {
	return &userService{
		userData: repo,
		validate: validator.New(),
	}
}

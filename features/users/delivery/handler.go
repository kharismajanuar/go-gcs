package delivery

import (
	"go-gcs/features/users"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userService users.UserService
}

// Add implements users.UserDelivery
func (*userHandler) Add() echo.HandlerFunc {
	panic("unimplemented")
}

func New(service users.UserService) users.UserDelivery {
	return &userHandler{
		userService: service,
	}
}

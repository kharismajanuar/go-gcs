package users

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID           uint
	Name         string `validate:"required,max=50"`
	DisplayImage string `validate:"required,max=50"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Image        ImageCore
}

type ImageCore struct {
	ID     uint
	UserID uint
	Url    string `validate:"required,max=50"`
}

type UserDelivery interface {
	Add() echo.HandlerFunc
}

type UserService interface {
	Create(input Core) error
}

type UserData interface {
	Insert(input Core) error
}

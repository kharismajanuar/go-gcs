package users

import (
	"mime/multipart"
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID        uint
	Name      string `validate:"required,max=50"`
	Avatar    string
	ImageFile multipart.File `json:"image" form:"image"`
	ImageName string
	CreatedAt time.Time
	UpdatedAt time.Time
	Image     ImageCore
}

type ImageCore struct {
	ID     uint
	UserID uint
	Url    string
}

type UserDelivery interface {
	Add(c echo.Context) error
	DeleteAvatar(c echo.Context) error
}

type UserService interface {
	Create(input Core) error
	DeleteAvatar(data Core) error
}

type UserData interface {
	Insert(input Core) error
	DeleteAvatar(data Core) error
}

package images

import (
	"mime/multipart"
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID        uint
	UserID    uint `validate:"required"`
	Url       string
	ImageFile multipart.File `validate:"required"`
	ImageName string         `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ImageDelivery interface {
	Add(c echo.Context) error
	Delete(c echo.Context) error
}

type ImageService interface {
	Create(input Core) error
	Delete(data Core) error
}

type ImageData interface {
	Insert(input Core) error
	Delete(data Core) error
}

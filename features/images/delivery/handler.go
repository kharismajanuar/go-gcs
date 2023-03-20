package delivery

import (
	"go-gcs/features/images"

	"github.com/labstack/echo/v4"
)

type imageHandler struct {
	imageService images.ImageService
}

// Add implements images.ImageDelivery
func (*imageHandler) Add(c echo.Context) error {
	panic("unimplemented")
}

// Delete implements images.ImageDelivery
func (*imageHandler) Delete(c echo.Context) error {
	panic("unimplemented")
}

func New(service images.ImageService) images.ImageDelivery {
	return &imageHandler{
		imageService: service,
	}
}

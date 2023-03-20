package delivery

import (
	"go-gcs/features/images"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type imageHandler struct {
	imageService images.ImageService
}

// Add implements images.ImageDelivery
func (delivery *imageHandler) Add(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, "id param must number")
	}

	imageInput := ImageRequest{}

	file, errFile := c.FormFile("image")
	if errFile != nil {
		return c.JSON(http.StatusBadRequest, "error form file")
	}

	blobFile, errBlob := file.Open()
	if errBlob != nil {
		return c.JSON(http.StatusNotFound, "blob file not found")
	}

	imageInput.ImageFile = blobFile

	dataCore := requestToCore(imageInput)
	dataCore.UserID = uint(idConv)
	dataCore.ImageName = file.Filename

	err := delivery.imageService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed add image")
	}
	return c.JSON(http.StatusCreated, "success add image")
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

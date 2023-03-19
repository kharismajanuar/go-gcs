package delivery

import (
	"go-gcs/features/users"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userService users.UserService
}

// Add implements users.UserDelivery
func (delivery *userHandler) Add(c echo.Context) error {
	userInput := UserRequest{}
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, "error bind data")
	}

	file, errFile := c.FormFile("image")
	if errFile != nil {
		return c.JSON(http.StatusBadRequest, "error form file")
	}

	blobFile, errBlob := file.Open()
	if errBlob != nil {
		return c.JSON(http.StatusNotFound, "blob file not found")
	}

	userInput.ImageFile = blobFile
	userInput.ImageName = file.Filename

	dataCore := requestToCore(userInput)

	err := delivery.userService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed add user")
	}
	return c.JSON(http.StatusCreated, "success add user")
}

func New(service users.UserService) users.UserDelivery {
	return &userHandler{
		userService: service,
	}
}

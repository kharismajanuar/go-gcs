package delivery

import (
	"go-gcs/features/users"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userService users.UserService
}

// DeleteAvatar implements users.UserDelivery
func (delivery *userHandler) DeleteAvatar(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, "id param must number")
	}

	userInput := UserRequest{}

	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, "error bind data")
	}

	dataCore := requestToCore(userInput)
	dataCore.ID = uint(idConv)

	err := delivery.userService.DeleteAvatar(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed delete user's avatar")
	}
	return c.JSON(http.StatusOK, "success delete user's avatar")
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

	dataCore := requestToCore(userInput)
	dataCore.ImageName = file.Filename

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

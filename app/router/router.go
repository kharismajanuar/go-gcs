package router

import (
	_imageData "go-gcs/features/images/data"
	_imageHandler "go-gcs/features/images/delivery"
	_imageService "go-gcs/features/images/service"
	_userData "go-gcs/features/users/data"
	_userHandler "go-gcs/features/users/delivery"
	_userService "go-gcs/features/users/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	// users
	userData := _userData.New(db)
	userService := _userService.New(userData)
	userHandler := _userHandler.New(userService)
	e.POST("/users", userHandler.Add)
	e.DELETE("/users/:id", userHandler.DeleteAvatar)

	// images
	imageData := _imageData.New(db)
	imageService := _imageService.New(imageData)
	imageHandler := _imageHandler.New(imageService)
	e.POST("/users/:id/images", imageHandler.Add)
}

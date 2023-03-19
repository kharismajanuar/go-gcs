package router

import (
	_userData "go-gcs/features/users/data"
	_userHandler "go-gcs/features/users/delivery"
	_userService "go-gcs/features/users/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	userData := _userData.New(db)
	userService := _userService.New(userData)
	userHandler := _userHandler.New(userService)
	e.POST("/users", userHandler.Add)
}

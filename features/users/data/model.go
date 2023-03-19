package data

import (
	_modelImage "go-gcs/features/images/data"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name   string `gorm:"not null"`
	Images []_modelImage.Image
}

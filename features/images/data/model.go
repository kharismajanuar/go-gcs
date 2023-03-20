package data

import (
	"go-gcs/features/images"

	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	UserID uint   `gorm:"not null"`
	Url    string `gorm:"not null"`
}

func CoreToModel(dataCore images.Core) Image {
	return Image{
		UserID: dataCore.UserID,
		Url:    dataCore.Url,
	}
}

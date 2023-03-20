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

func ModelToCore(dataModel Image) images.Core {
	return images.Core{
		ID:        dataModel.ID,
		UserID:    dataModel.UserID,
		Url:       dataModel.Url,
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
	}
}

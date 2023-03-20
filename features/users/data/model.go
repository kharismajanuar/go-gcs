package data

import (
	_modelImage "go-gcs/features/images/data"
	"go-gcs/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name   string `gorm:"not null"`
	Avatar string
	Images []_modelImage.Image
	Image  Image
}

type Image struct {
	gorm.Model
	UserID uint
	Url    string `gorm:"not null"`
}

func CoreToModel(dataCore users.Core) User {
	return User{
		Name:   dataCore.Name,
		Avatar: dataCore.Avatar,
	}
}

func ModelToCore(dataModel User) users.Core {
	return users.Core{
		ID:        dataModel.ID,
		Name:      dataModel.Name,
		Avatar:    dataModel.Avatar,
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
		Image: users.ImageCore{
			Url: dataModel.Image.Url,
		},
	}
}

func ListModelToCore(dataModel []User) []users.Core {
	var dataCore []users.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, ModelToCore(v))
	}
	return dataCore
}

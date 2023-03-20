package data

import (
	"go-gcs/features/images"

	"gorm.io/gorm"
)

type imageQuery struct {
	db *gorm.DB
}

// Delete implements images.ImageData
func (*imageQuery) Delete(data images.Core) error {
	panic("unimplemented")
}

// Insert implements images.ImageData
func (*imageQuery) Insert(input images.Core) error {
	panic("unimplemented")
}

func New(db *gorm.DB) images.ImageData {
	return &imageQuery{
		db: db,
	}
}

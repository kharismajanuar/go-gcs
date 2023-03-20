package data

import (
	"errors"
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
func (repo *imageQuery) Insert(input images.Core) error {
	dataModel := CoreToModel(input)
	tx := repo.db.Create(&dataModel)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert error, row affected = 0")
	}
	return nil
}

func New(db *gorm.DB) images.ImageData {
	return &imageQuery{
		db: db,
	}
}

package data

import (
	"errors"
	"go-gcs/features/images"

	"gorm.io/gorm"
)

type imageQuery struct {
	db *gorm.DB
}

// SelectById implements images.ImageData
func (repo *imageQuery) SelectById(id uint) (images.Core, error) {
	dataModel := Image{}
	tx := repo.db.Where("id = ?", id).First(&dataModel)
	if tx.Error != nil {
		return images.Core{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return images.Core{}, errors.New("select error, row affected = 0")
	}
	return ModelToCore(dataModel), nil
}

// Delete implements images.ImageData
func (repo *imageQuery) Delete(id uint) error {
	tx := repo.db.Where("id = ?", id).Delete(&Image{})
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("delete error, row affected = 0")
	}
	return nil
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

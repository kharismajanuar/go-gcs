package data

import (
	"errors"
	"go-gcs/features/users"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

// DeleteAvatar implements users.UserData
func (repo *userQuery) DeleteAvatar(data users.Core) error {
	dataModel := CoreToModel(data)
	tx := repo.db.Model(&dataModel).Where("id = ?", data.ID).Update("avatar", "https://storage.googleapis.com/go-gcs-13/test-files/default_avatar")
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("update error, row affected = 0")
	}
	return nil
}

// Insert implements users.UserData
func (repo *userQuery) Insert(input users.Core) error {
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

func New(db *gorm.DB) users.UserData {
	return &userQuery{
		db: db,
	}
}

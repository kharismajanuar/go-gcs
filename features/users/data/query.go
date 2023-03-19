package data

import (
	"go-gcs/features/users"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

// Insert implements users.UserData
func (*userQuery) Insert(input users.Core) error {
	panic("unimplemented")
}

func New(db *gorm.DB) users.UserData {
	return &userQuery{
		db: db,
	}
}

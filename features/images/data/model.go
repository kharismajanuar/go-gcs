package data

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	UserID uint
	Url    string `gorm:"not null"`
}

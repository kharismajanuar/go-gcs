package data

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	Url    string `gorm:"not null"`
	UserID uint
}

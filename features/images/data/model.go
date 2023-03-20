package data

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	UserID uint   `gorm:"not null"`
	Url    string `gorm:"not null"`
}

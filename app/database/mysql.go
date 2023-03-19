package database

import (
	"fmt"
	"go-gcs/app/config"
	"log"

	_modelImage "go-gcs/features/images/data"
	_modelUser "go-gcs/features/users/data"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg config.AppConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOSTNAME, cfg.DB_PORT, cfg.DB_NAME)
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Println("error connect to DB", err.Error())
		return nil
	}

	return db
}

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(_modelUser.User{})
	db.AutoMigrate(_modelImage.Image{})
}

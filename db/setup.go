package db

import (
	"go_example/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	dsn := "host=localhost " +
		"user=srkn " +
		"dbname=go_example " +
		"port=5432 " +
		"password=1234 " +
		"sslmode=disable " +
		"TimeZone=Europe/Moscow"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Database connection error: " + err.Error())
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return
	}

	DB = db
}

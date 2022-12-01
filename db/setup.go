package db

import (
	"go_example/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func Init() {

	dsn := "host=localhost user=srkn password=1234 dbname=go_example port=5432 sslmode=disable TimeZone=Europe/Moscow"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed database connection")
	}
	db.AutoMigrate(&model.User{})
}

func DbManager() *gorm.DB {
	return db
}

package model

type User struct {
	Base
	Username string `gorm:"unique_index;not null"`
	Email    string `gorm:"unique_index;notnull"`
	Password string `gorm:"not null"`
}

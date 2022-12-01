package model

type User struct {
	Base
	Username string `json:"username" gorm:"not null" validate:"required"`
	Email    string `json:"email" gorm:"unique_index;not null" validate:"required,email,unique"`
	Password string `json:"password" gorm:"not null" validate:"required"`
}

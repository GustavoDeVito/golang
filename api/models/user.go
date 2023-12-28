package models

import (
	"api/helpers"
)

type User struct {
	helpers.CustomModel
	Name     string `json:"name" validate:"required"`
	Age      int16  `json:"age" validate:"required,gte=18"`
	Username string `json:"username" gorm:"unique;<-:create" validate:"required"`
	Password string `json:"password" gorm:"->"`
}

type UserPassword struct {
	User
	Password string `json:"password" validate:"required"`
}

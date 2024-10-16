package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string
	WxOpenId    string
	PhoneNumber string
	Address     string
}

func CreateUser(user User) {
	db.Create(&user)
}

package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	WxOpenId    string
	PhoneNumber string
	Address     string
}

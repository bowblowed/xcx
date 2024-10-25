package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string
	WxNumber    string
	WxOpenId    string
	PhoneNumber string
	Address     string
}

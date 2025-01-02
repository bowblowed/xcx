package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Type        string
	Name        string
	WxNumber    string
	WxOpenId    string
	PhoneNumber string
	Address     string
	ShopCartId  uint
}

var UserTypeAdmin = "admin"
var UserTypeUserUnsigned = "userUnSigned"
var UserTypeUserSigned = "userSigned"

package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	CategoryName Category
	Name         string
	Price        uint32
	Count        uint32
	Description  string
	PicUrl       string
}

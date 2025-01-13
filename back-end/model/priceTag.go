package model

import (
	"gorm.io/gorm"
)

type PriceTag struct {
	gorm.Model
	ProductId uint
	Price     uint
	PropPairs map[string]string `gorm:"serializer:json"`
}

package model

import (
	"gorm.io/gorm"
)

type OrderItem struct {
	ProductId uint
	Props     map[string]string
	Count     uint
	Price     uint
}

type OrderItems []OrderItem

type Order struct {
	gorm.Model
	UserId     uint
	OrderItems OrderItems `gorm:"serializer:json"`
	Address    string
	TotalPrice uint
	Stataus    string
}

var OrderStatusDefault = "未付款"
var OderStatusCancel = "已取消"
var OrderStatusPayed = "已付款"
var OrderStatusShip = "已发货"
var OrderStatusFinish = "已完成"

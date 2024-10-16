package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

type OrderItem struct {
	ProductId uint
	Count     uint
}
type OrderItems []OrderItem

func (s OrderItems) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s *OrderItems) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("type assertion to []byte failed")
	}
	return json.Unmarshal(b, s)
}

type Order struct {
	gorm.Model
	UserId     uint
	User       User
	OrderItems OrderItems
}

func CreateOrder(order Order) {
	db.Create(&order)
}

func GetOrder(id uint) Order {
	var order Order
	db.First(&order)
	return order
}

package service

import "back-end/model"

func CreateOrder(order model.Order) error {
	if order.Stataus == "" {
		order.Stataus = model.OrderStatusDefault
	}
	return model.DB.Create(&order).Error
}

func DeleteOrder(id uint) error {
	return model.DB.Delete(&model.Order{}, id).Error
}

func UpdateOrder(order model.Order) error {
	return model.DB.Updates(&order).Error
}

func ListOrders() ([]model.Order, error) {
	var orders []model.Order
	err := model.DB.Find(&orders).Error
	return orders, err
}

func ListOrdersByUserId(userId uint) ([]model.Order, error) {
	var orders []model.Order
	err := model.DB.Where("user_id = ?", userId).Find(&orders).Error
	return orders, err
}

func GetOrder(id uint) (model.Order, error) {
	var order model.Order
	err := model.DB.First(&order, id).Error
	return order, err
}

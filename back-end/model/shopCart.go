package model

import (
	"gorm.io/gorm"
)

type ShopCart struct {
	gorm.Model
	UserId     uint
	OrderItems []OrderItem `gorm:"serializer:json"`
	Products   []Product   `gorm:"-"`
}

func (s *ShopCart) AfterFind(tx *gorm.DB) error {
	// 收集所有不重复的产品ID
	productIdsMap := make(map[uint]struct{})
	for _, item := range s.OrderItems {
		productIdsMap[item.ProductId] = struct{}{}
	}

	// 转换为切片
	uniqueProductIds := make([]uint, 0, len(productIdsMap))
	for id := range productIdsMap {
		uniqueProductIds = append(uniqueProductIds, id)
	}

	// 一次性查询所有唯一产品
	var uniqueProducts []Product
	err := DB.Find(&uniqueProducts, uniqueProductIds).Error
	if err != nil {
		return err
	}

	// 建立产品ID到产品的映射
	productMap := make(map[uint]Product)
	for _, p := range uniqueProducts {
		productMap[p.ID] = p
	}

	// 按OrderItems的顺序填充Products
	s.Products = make([]Product, len(s.OrderItems))
	for i, item := range s.OrderItems {
		s.Products[i] = productMap[item.ProductId]
	}

	return nil
}

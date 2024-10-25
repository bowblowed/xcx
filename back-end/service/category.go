package service

import (
	"back-end/model"
	"fmt"

	"gorm.io/gorm"
)

func GetCategory(id uint) (model.Category, error) {
	var category model.Category
	err := model.DB.First(&category, id).Error
	return category, fmt.Errorf("failed to get category: %w", err)
}

func CreateCategory(name string) error {
	// 查询已有的分类,确保名字不能重复
	var cat model.Category
	model.DB.Unscoped().Where("name = ?", name).Find(&cat)
	if cat.ID != 0 {
		if !cat.DeletedAt.Time.IsZero() {
			// 分类已存在且已被软删除，可以选择恢复该分类
			cat.DeletedAt = gorm.DeletedAt{}
			if err := model.DB.Save(&cat).Error; err != nil {
				return fmt.Errorf("failed to restore category: %w", err)
			}
			return fmt.Errorf("category with name %s already exists and has been restored", name)
		} else {
			// 分类已存在且未被删除
			return fmt.Errorf("category with name %s already exists", name)
		}
	}
	// 创建新的分类
	if err := model.DB.Create(&model.Category{Name: name}).Error; err != nil {
		return fmt.Errorf("failed to create category: %w", err)
	}
	return nil
}

func ListCategories() ([]model.Category, error) {
	var categories []model.Category
	err := model.DB.Find(&categories).Error
	return categories, fmt.Errorf("failed to list categories: %w", err)
}

func DeleteCategory(id uint) error {
	return model.DB.Delete(&model.Category{}, id).Error
}

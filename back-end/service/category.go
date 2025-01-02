package service

import (
	"back-end/model"
	"fmt"
	"mime/multipart"

	"gorm.io/gorm"
)

func GetCategory(id uint) (model.Category, error) {
	var category model.Category
	err := model.DB.First(&category, id).Error
	if err != nil {
		return category, fmt.Errorf("failed to get category: %w", err)
	}
	return category, nil
}

func CreateCategory(name string, file ...*multipart.FileHeader) error {
	// 查询已有的分类,确保名字不能重复
	var cat model.Category
	model.DB.Unscoped().Where("name = ?", name).Find(&cat)
	if cat.ID != 0 {
		if !cat.DeletedAt.Time.IsZero() {
			// 分类已存在且已被软删除，可以选择恢复该分类
			cat.DeletedAt = gorm.DeletedAt{}
			if err := model.DB.Save(&cat).Error; err != nil {
				return nil
			}
			return fmt.Errorf("category with name %s already exists and has been restored", name)
		} else {
			// 分类已存在且未被删除
			return fmt.Errorf("category with name %s already exists", name)
		}
	}
	// 创建新的分类
	var newCat model.Category
	newCat.Name = name
	if err := model.DB.Create(&newCat).Error; err != nil {
		return fmt.Errorf("failed to create category: %w", err)
	}
	if len(file) != 0 && file[0] != nil {
		err := UploadCategoryPic(file[0], newCat.ID)
		return err
	}
	return nil
}

func ListCategories() ([]model.Category, error) {
	var categories []model.Category
	err := model.DB.Find(&categories).Error
	if err != nil {
		return categories, fmt.Errorf("failed to list categories: %w", err)
	}
	return categories, nil
}

func DeleteCategory(id uint) error {
	cate, err := GetCategory(id)
	if err != nil {
		return err
	}
	if cate.PicId != 0 {
		DeleteFile(cate.PicId)
	}
	return model.DB.Delete(&model.Category{}, id).Error
}

func UpdateCategory(category model.Category) error {
	return model.DB.Updates(&category).Error
}

func UploadCategoryPic(file *multipart.FileHeader, cateId uint) error {
	cat, err := GetCategory(cateId)
	if err != nil {
		return fmt.Errorf("failed to upload cate pic: %w", err)
	}
	if cat.PicId != 0 {
		DeleteFile(cat.PicId)
	}
	picId, err := CreateFile(file, "category")
	if err != nil {
		return fmt.Errorf("failed to upload cate pic: %w", err)
	}
	cat.PicId = picId
	return UpdateCategory(cat)
}

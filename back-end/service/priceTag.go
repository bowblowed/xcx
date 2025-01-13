package service

import "back-end/model"

func ListPriceTagByProductId(productId uint) ([]model.PriceTag, error) {
	var priceTags []model.PriceTag
	err := model.DB.Find(&priceTags, "product_id = ?", productId).Error
	return priceTags, err
}

func CreatePriceTag(priceTag model.PriceTag) error {
	return model.DB.Create(&priceTag).Error
}

func DeletePriceTag(id uint) error {
	return model.DB.Delete(&model.PriceTag{}, id).Error
}

func ListPriceTag() ([]model.PriceTag, error) {
	var priceTags []model.PriceTag
	err := model.DB.Find(&priceTags).Error
	return priceTags, err
}

func UpdatePriceTag(tag model.PriceTag) error {
	return model.DB.Updates(&tag).Error
}

package service

import (
	"back-end/model"
	"fmt"
)

func GetBanner(id uint) (model.Banner, error) {
	var banner model.Banner
	err := model.DB.First(&banner, id).Error
	return banner, fmt.Errorf("failed to get banner: %w", err)
}
func CreateBanner(PicUrl string) error {
	err := model.DB.Create(&model.Banner{PicUrl: PicUrl}).Error
	return fmt.Errorf("failed to create banner: %w", err)
}

func ListBanners() ([]model.Banner, error) {
	var banners []model.Banner
	err := model.DB.Find(&banners).Error
	return banners, fmt.Errorf("failed to list banners: %w", err)
}

func DeleteBanner(id uint) error {
	err := model.DB.Delete(&model.Banner{}, id).Error
	return fmt.Errorf("failed to delete banner: %w", err)
}

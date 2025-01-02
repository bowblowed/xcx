package service

import (
	"back-end/model"
	"fmt"
	"mime/multipart"
)

func GetBanner(id uint) (model.Banner, error) {
	var banner model.Banner
	err := model.DB.First(&banner, id).Error
	if err != nil {
		return banner, fmt.Errorf("failed to get banner: %w", err)
	}
	return banner, nil
}
func CreateBanner(file *multipart.FileHeader) error {
	picId, err := CreateFile(file, "banner")
	if err != nil {
		return err
	}
	err = model.DB.Create(&model.Banner{PicId: picId}).Error
	if err != nil {
		return fmt.Errorf("failed to create banner: %w", err)
	}
	return nil
}

func ListBanners() ([]model.Banner, error) {
	var banners []model.Banner
	err := model.DB.Find(&banners).Error
	if err != nil {
		return banners, fmt.Errorf("failed to list banners: %w", err)
	}
	return banners, nil
}

func DeleteBanner(id uint) error {
	banner, err := GetBanner(id)
	if err != nil {
		return err
	}
	if banner.PicId != 0 {
		err = DeleteFile(banner.PicId)
		if err != nil {
			return fmt.Errorf("failed to delete banner: %w", err)
		}
	}
	err = model.DB.Delete(&model.Banner{}, id).Error
	if err != nil {
		return fmt.Errorf("failed to delete banner: %w", err)
	}
	return nil
}

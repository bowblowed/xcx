package model

import "gorm.io/gorm"

type Banner struct {
	gorm.Model
	PicId  uint
	PicUrl string `gorm:"-"`
}

func (u *Banner) AfterFind(tx *gorm.DB) (err error) {
	if u.PicId != 0 {
		var pic File
		err := DB.First(&pic, u.PicId).Error
		if err != nil {
			return err
		}
		u.PicUrl = pic.Url
	}
	return nil
}

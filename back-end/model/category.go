package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name   string
	PicId  uint
	PicUrl string `gorm:"-"`
}

func (c *Category) AfterFind(tx *gorm.DB) (err error) {
	if c.PicId != 0 {
		var pic File
		err := DB.First(&pic, c.PicId).Error
		if err != nil {
			return err
		}
		c.PicUrl = pic.Url
	}
	return nil
}

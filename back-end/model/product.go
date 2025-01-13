package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	CategoryId         uint
	CategoryName       string `gorm:"-"`
	Name               string
	DefaultPrice       uint
	Count              uint32
	Description        string
	Content            string
	PicId              uint
	PicUrl             string              `gorm:"-"`
	DescriptionPicIds  []uint              `gorm:"serializer:json"`
	DescriptionPicUrls []string            `gorm:"-"`
	Props              map[string][]string `gorm:"serializer:json"`
}

func (p *Product) AfterFind(tx *gorm.DB) (err error) {
	if p.PicId != 0 {
		var pic File
		err := DB.First(&pic, p.PicId).Error
		if err != nil {
			return err
		}
		p.PicUrl = pic.Url
	}
	p.DescriptionPicUrls = []string{}
	if len(p.DescriptionPicIds) != 0 {
		var pics []File
		err := DB.Find(&pics, p.DescriptionPicIds).Error
		if err != nil {
			return err
		}
		for _, pic := range pics {
			p.DescriptionPicUrls = append(p.DescriptionPicUrls, pic.Url)
		}
	}
	var category Category
	err = DB.First(&category, p.CategoryId).Error
	if err != nil {
		return err
	}
	p.CategoryName = category.Name
	return nil
}

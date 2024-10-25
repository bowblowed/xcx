package model

import "gorm.io/gorm"

type Banner struct {
	gorm.Model
	PicUrl string `json:"picUrl"`
}

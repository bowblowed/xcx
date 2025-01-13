package model

import (
	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	Path string
	Url  string
	Name string
	Type string
}

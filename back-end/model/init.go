package model

import (
	"back-end/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", config.Database.Host, config.Database.User, config.Database.Password, config.Database.Database, config.Database.Port)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("无法连接到数据库：" + err.Error())
	}
	fmt.Println("连接数据库成功")
	DB.AutoMigrate(&Category{}, &Product{}, &Order{}, &OrderItem{}, &User{}, &Banner{})
}

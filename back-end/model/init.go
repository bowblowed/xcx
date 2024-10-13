package model

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	dsn := "host=localhost user=xcx password=123456 dbname=xcx port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("无法连接到数据库：" + err.Error())
	}
	fmt.Println(db)
}

func F() {

}

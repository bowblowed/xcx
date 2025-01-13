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
	err = DB.AutoMigrate(&Category{}, &Product{}, &Order{}, &User{}, &Banner{}, &File{}, &PriceTag{}, &ShopCart{})
	if err != nil {
		panic("无法迁移数据库：" + err.Error())
	}
}

func ClearDB() {
	migrator := DB.Migrator()
	tableNames, err := migrator.GetTables()
	if err != nil {
		panic(fmt.Sprintf("获取表名失败：%v", err))
	}

	for _, tableName := range tableNames {
		if err := migrator.DropTable(tableName); err != nil {
			fmt.Printf("删除表 %s 失败：%v\n", tableName, err)
		} else {
			fmt.Printf("成功删除表 %s\n", tableName)
		}
	}
}

package common

import (
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"updateTool/model"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	// 根据配置获取sqlite文件位置
	dbPath := viper.GetString("datasource.path")

	// 创建数据库连接
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	db.AutoMigrate(&model.UpdateHistory{})
	db.AutoMigrate(&model.User{})

	return db
}

func GetDB() *gorm.DB {
	if DB == nil {
		DB = InitDB()
	}
	return DB
}

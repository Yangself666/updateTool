package common

import (
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"updateTool/model"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	// 根据配置获取sqlite文件位置
	dbPath := viper.GetString("datasource.path")

	// 创建数据库连接
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}

	db.AutoMigrate(&model.UpdateHistory{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.ProjectPath{})
	db.AutoMigrate(&model.Project{})
	db.AutoMigrate(&model.Server{})
	db.AutoMigrate(&model.ProjectServerCon{})
	db.AutoMigrate(&model.ProjectUserCon{})
	db.AutoMigrate(&model.Permission{})
	db.AutoMigrate(&model.UserPermissionCon{})

	return db
}

func GetDB() *gorm.DB {
	if DB == nil {
		DB = InitDB()
	}
	return DB
}

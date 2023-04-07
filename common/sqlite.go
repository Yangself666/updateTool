package common

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"updateTool/model"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	// 根据配置获取sqlite文件位置
	dbPath := viper.GetString("datasource.path")

	// 创建数据库连接
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		// 设置日志打印sql
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}

	db.AutoMigrate(&model.UpdateHistory{})
	db.AutoMigrate(&model.ProjectPath{})
	db.AutoMigrate(&model.Project{})
	db.AutoMigrate(&model.Server{})
	db.AutoMigrate(&model.ProjectServerCon{})
	db.AutoMigrate(&model.ProjectUserCon{})
	if exists := db.Migrator().HasTable(&model.User{}); !exists {
		db.AutoMigrate(&model.User{})
		// 如果表不存在，默认添加管理员账号
		DefaultAdminGenerator(db)
	} else {
		db.AutoMigrate(&model.User{})
	}
	if exists := db.Migrator().HasTable(&model.Permission{}); !exists {
		db.AutoMigrate(&model.Permission{})
		// 如果表不存在，默认添加所有内置权限
		DefaultPermissionsGenerator(db)
	} else {
		db.AutoMigrate(&model.Permission{})
	}
	db.AutoMigrate(&model.UserPermissionCon{})

	return db
}

func GetDB() *gorm.DB {
	if DB == nil {
		DB = InitDB()
	}
	return DB
}

// DefaultPermissionsGenerator 默认权限生成器
func DefaultPermissionsGenerator(db *gorm.DB) {
	var permissions = []model.Permission{
		{PermissionName: "添加用户", MenuName: "user", PermissionPath: "/api/user/add"},
		{PermissionName: "修改用户信息", MenuName: "user", PermissionPath: "/api/user/edit"},
		{PermissionName: "修改用户密码", MenuName: "user", PermissionPath: "/api/user/editPassword"},
		{PermissionName: "删除用户", MenuName: "user", PermissionPath: "/api/user/del"},
		{PermissionName: "获取所有用户列表", MenuName: "user", PermissionPath: "/api/user/list"},
		{PermissionName: "批量编辑用户权限", MenuName: "user", PermissionPath: "/api/user/editPermission"},
		{PermissionName: "根据用户ID获取用户权限列表", MenuName: "user", PermissionPath: "/api/user/userPermission"},
		{PermissionName: "设置用户为管理员", MenuName: "user", PermissionPath: "/api/user/setUserAsAdmin"},
		{PermissionName: "设置用户为非管理员", MenuName: "user", PermissionPath: "/api/user/setUserAsNonAdmin"},
		{PermissionName: "添加项目", MenuName: "project", PermissionPath: "/api/project/add"},
		{PermissionName: "删除项目", MenuName: "project", PermissionPath: "/api/project/del"},
		{PermissionName: "删除前检测项目是否可以删除", MenuName: "project", PermissionPath: "/api/project/delCheck"},
		{PermissionName: "修改项目信息", MenuName: "project", PermissionPath: "/api/project/edit"},
		{PermissionName: "获取项目列表", MenuName: "project", PermissionPath: "/api/project/list"},
		{PermissionName: "获取单个项目信息", MenuName: "project", PermissionPath: "/api/project/info"},
		{PermissionName: "通过项目ID获取路径信息", MenuName: "upload", PermissionPath: "/api/project/path"},
		{PermissionName: "通过用户ID获取绑定的项目", MenuName: "project", PermissionPath: "/api/project/userId"},
		{PermissionName: "批量编辑用户关联", MenuName: "project", PermissionPath: "/api/project/user/edit"},
		{PermissionName: "获取项目中的用户列表", MenuName: "project", PermissionPath: "/api/project/user/list"},
		{PermissionName: "批量编辑服务器关联", MenuName: "project", PermissionPath: "/api/project/server/edit"},
		{PermissionName: "获取项目中的服务器列表", MenuName: "project", PermissionPath: "/api/project/server/list"},
		{PermissionName: "添加路径信息", MenuName: "project", PermissionPath: "/api/project/path/add"},
		{PermissionName: "删除路径信息", MenuName: "project", PermissionPath: "/api/project/path/del"},
		{PermissionName: "修改路径信息", MenuName: "project", PermissionPath: "/api/project/path/edit"},
		{PermissionName: "添加服务器", MenuName: "server", PermissionPath: "/api/server/add"},
		{PermissionName: "删除服务器", MenuName: "server", PermissionPath: "/api/server/del"},
		{PermissionName: "修改服务器信息", MenuName: "server", PermissionPath: "/api/server/edit"},
		{PermissionName: "获取服务器列表", MenuName: "server", PermissionPath: "/api/server/list"},
		{PermissionName: "检查服务器是否可以连接", MenuName: "server", PermissionPath: "/api/server/check"},
		{PermissionName: "添加权限信息", MenuName: "permission", PermissionPath: "/api/permission/add"},
		{PermissionName: "删除权限信息", MenuName: "permission", PermissionPath: "/api/permission/del"},
		{PermissionName: "修改权限信息", MenuName: "permission", PermissionPath: "/api/permission/edit"},
		{PermissionName: "获取权限列表", MenuName: "permission", PermissionPath: "/api/permission/list"},
		{PermissionName: "添加权限信息", MenuName: "permission", PermissionPath: "/api/permission/add"},
		{PermissionName: "上传文件接口", MenuName: "upload", PermissionPath: "/api/upload/file"},
		{PermissionName: "获取上传历史", MenuName: "history", PermissionPath: "/api/history/get"},
		{PermissionName: "回滚到历史文件", MenuName: "history", PermissionPath: "/api/history/rollback"},
		{PermissionName: "下载历史文件", MenuName: "history", PermissionPath: "/api/history/download"},
	}
	db.Create(&permissions)
}

// DefaultAdminGenerator 默认用户生成器
func DefaultAdminGenerator(db *gorm.DB) {
	defaultPassword := "123456"
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(defaultPassword), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("默认管理密码生成失败")
	}
	var admin = model.User{
		Name:     "Admin",
		Email:    "admin@qq.com",
		Password: string(hashPassword),
		IsAdmin:  true,
	}
	fmt.Printf("创建%v", admin)
	db.Create(&admin)
}

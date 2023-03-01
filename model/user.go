package model

import "gorm.io/gorm"

// User 用户结构体
type User struct {
	gorm.Model
	// 用户名
	Name string `json:"name" gorm:"type:varchar(20);not null;comment:'用户名''"`
	// 用户登陆邮箱
	Email string `json:"email" gorm:"type:varchar(50);not null;unique;comment:'用户登陆邮箱''"`
	// 加密的用户密码
	Password string `json:"password" gorm:"size:255;not null;comment:'加密的用户密码''"`
	// 是否为管理员
	IsAdmin bool `json:"isAdmin" gorm:"type:tinyint(2);not null;default:0;comment:'是否为管理员''"`
}

package model

import "gorm.io/gorm"

// Permission 权限实体类
type Permission struct {
	gorm.Model
	// 权限名称
	PermissionName string `json:"permissionName" gorm:"type:varchar(50);not null;comment: '权限名称''"`
	// 菜单标识
	MenuName string `json:"menuName" gorm:"type:varchar(50);not null;comment: '菜单标识''"`
	// 权限路径
	PermissionPath string `json:"permissionPath" gorm:"type:varchar(200);not null;comment: '权限路径''"`
}

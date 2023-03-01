package model

import "gorm.io/gorm"

// UserPermissionCon 用户和权限关联表
type UserPermissionCon struct {
	gorm.Model
	// 用户ID
	UserId uint `json:"userId" gorm:"type:int;not null;comment:'用户ID'"`
	// 权限ID
	PermissionId uint `json:"permissionId" gorm:"type:int;not null;comment:'权限ID'"`
}

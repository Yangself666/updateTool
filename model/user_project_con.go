package model

import "gorm.io/gorm"

// UserProjectCon 用户和项目关联模型
type UserProjectCon struct {
	gorm.Model
	UserId    uint `json:"userId" gorm:"type:int;not null;comment:'用户ID'"`
	ProjectId uint `json:"projectId" gorm:"type:int;not null;comment:'项目ID'"`
}

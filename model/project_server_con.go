package model

import "gorm.io/gorm"

// ProjectServerCon 项目和服务器关联模型
type ProjectServerCon struct {
	gorm.Model
	ProjectId uint `json:"projectId" gorm:"type:int;not null;comment:'项目ID'"`
	ServerId  uint `json:"serverId" gorm:"type:int;not null;comment:'服务器ID'"`
}

package model

import (
	"gorm.io/gorm"
	"reflect"
)

// 项目路径模型（上传文件白名单）

type ProjectPath struct {
	gorm.Model
	// 关联的项目ID
	ProjectId uint `json:"projectId" gorm:"type:int;not null;comment: '关联的项目ID'"`
	// 路径名称
	PathName string `json:"pathName" gorm:"type:varchar(50);not null;comment: '路径名称'"`
	// 绝对路径
	Path string `json:"path" gorm:"type:varchar(500);not null;comment: '绝对路径'"`
	// 是否包含子路径
	HasSubPath int `json:"hasSubPath" gorm:"type:int(4);not null;default:1;comment: '是否包含子路径 1:包含 2:不包含'"`
}

func (projectPath ProjectPath) IsEmpty() bool {
	return reflect.DeepEqual(projectPath, ProjectPath{})
}

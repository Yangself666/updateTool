package model

import "gorm.io/gorm"

// UpdateHistory 更新历史
type UpdateHistory struct {
	gorm.Model
	// 上传的远程路径
	RemotePath string `json:"remotePath" gorm:"type:varchar(500);not null;comment: '上传的远程路径'"`
	// 本地存储路径
	LocalPath string `json:"localPath" gorm:"type:varchar(500);not null;comment: '本地存储路径'"`
	// 文件名称
	FileName string `json:"fileName" gorm:"type:varchar(500);not null;comment: '文件名称'"`
	// 唯一文件名
	UniqueFileName string `json:"uniqueFileName" gorm:"type:varchar(500);not null;comment: '唯一文件名'"`
	// 所属项目ID
	ProjectId uint `json:"projectId" gorm:"type:int;not null;default:0;comment: '所属项目ID'"`
	// 所属路径ID
	PathId uint `json:"pathId" gorm:"type:int;not null;default:0;comment: '所属路径ID'"`
	// 更新服务器相关信息
	ServerInfo string `json:"serverInfo" gorm:"type:varchar(3000);null;comment: '更新服务器相关信息'"`
	// 备注信息
	OtherInfo string `json:"otherInfo" gorm:"type:varchar(500);null;comment: '备注信息'"`
}

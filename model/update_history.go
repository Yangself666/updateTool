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
	// 备注信息
	OtherInfo string `json:"otherInfo" gorm:"type:varchar(500);null;comment: '备注信息'"`
}

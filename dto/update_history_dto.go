package dto

import (
	"gorm.io/gorm"
	"updateTool/model"
)

// UpdateHistoryDto 更新历史
type UpdateHistoryDto struct {
	gorm.Model
	// 操作用户ID
	UserId uint `json:"userId"`
	// 上传的远程路径
	RemotePath string `json:"remotePath"`
	// 本地存储路径
	LocalPath string `json:"localPath"`
	// 文件名称
	FileName string `json:"fileName"`
	// 唯一文件名
	UniqueFileName string `json:"uniqueFileName"`
	// 所属项目ID
	ProjectId uint `json:"projectId"`
	// 所属路径ID
	PathId uint `json:"pathId"`
	// 更新服务器相关信息
	ServerInfo string `json:"serverInfo"`
	// 备注信息
	OtherInfo string `json:"otherInfo"`

	// 用户信息
	UserInfo UserDto
	// 项目信息
	ProjectInfo model.Project
	// 路径信息
	PathInfo model.ProjectPath
}

func ToUpdateHistoryDto(history model.UpdateHistory) UpdateHistoryDto {
	return UpdateHistoryDto{
		Model:          history.Model,
		UserId:         history.UserId,
		RemotePath:     history.RemotePath,
		LocalPath:      history.LocalPath,
		FileName:       history.FileName,
		UniqueFileName: history.UniqueFileName,
		ProjectId:      history.ProjectId,
		PathId:         history.PathId,
		ServerInfo:     history.ServerInfo,
		OtherInfo:      history.OtherInfo,
	}
}

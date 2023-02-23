package dto

import (
	"gorm.io/gorm"
	"updateTool/model"
)

type ProjectDto struct {
	gorm.Model
	// 项目名称
	ProjectName string `json:"projectName" gorm:"type:varchar(50);not null;comment: '项目名称''"`
	// 项目简介
	ProjectIntro string `json:"projectIntro" gorm:"type:varchar(500);null;comment: '项目简介''"`
	// 项目绑定的服务器ID列表
	ServerIdList []int `json:"serverIdList"`
	// 项目绑定的服务器信息
	ServerList []model.Server `json:"serverList"`
	// 项目绑定的路径信息
	ProjectPathList []model.ProjectPath `json:"projectPathList"`
}

func ProjectDtoToProject(projectDto ProjectDto) model.Project {
	return model.Project{
		Model:        projectDto.Model,
		ProjectName:  projectDto.ProjectName,
		ProjectIntro: projectDto.ProjectIntro,
	}
}

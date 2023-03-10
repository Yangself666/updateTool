package dto

import (
	"gorm.io/gorm"
	"updateTool/model"
)

type ProjectDto struct {
	gorm.Model
	// 项目名称
	ProjectName string `json:"projectName"`
	// 项目简介
	ProjectIntro string `json:"projectIntro"`
	// 项目绑定的服务器ID列表
	ServerIdList []int `json:"serverIdList"`
	// 项目绑定的服务器信息
	ServerList []model.Server `json:"serverList"`
	// 项目绑定的用户ID列表
	UserIdList []int `json:"userIdList"`
	// 项目绑定的用户信息列表
	UserList []model.User `json:"userList"`
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

func ToProjectDto(project model.Project) ProjectDto {
	return ProjectDto{
		Model:        project.Model,
		ProjectName:  project.ProjectName,
		ProjectIntro: project.ProjectIntro,
	}
}

package model

import "gorm.io/gorm"

// 项目配置模型

type Project struct {
	gorm.Model
	// 项目名称
	ProjectName string `json:"projectName" gorm:"type:varchar(50);not null;comment: '项目名称''"`
	// 项目简介
	ProjectIntro string `json:"projectIntro" gorm:"type:varchar(500);null;comment: '项目简介''"`
}

package model

import "gorm.io/gorm"

// UpdateHistory 更新历史
type UpdateHistory struct {
	gorm.Model
	RemotePath     string `gorm:"type:varchar(500);not null" json:"remotePath"`
	LocalPath      string `gorm:"type:varchar(500);not null" json:"localPath"`
	FileName       string `gorm:"type:varchar(500);not null" json:"fileName"`
	UniqueFileName string `gorm:"type:varchar(500);not null" json:"uniqueFileName"`
	OtherInfo      string `gorm:"type:varchar(500)" json:"otherInfo"`
}

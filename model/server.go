package model

import "gorm.io/gorm"

// 服务器配置模型

// Server 服务器配置
type Server struct {
	gorm.Model
	// 服务器名称
	ServerName string `json:"serverName" gorm:"type:varchar(50);not null;comment: '服务器名称'"`
	// 服务器认证类型 1:password 2:secret
	ServerType int `json:"serverType" gorm:"type:int(4);not null;default:1;comment: '服务器认证类型 1:password 2:secret'"`
	// 服务器主机地址
	Host string `json:"host" gorm:"type:varchar(50);not null;comment: '服务器主机地址'"`
	// 服务器端口
	Port int `json:"port" gorm:"type:int(10);not null;comment: '服务器端口'"`
	// 服务器登陆用户名（服务器类型为0时所需要）
	Username string `json:"username" gorm:"type:varchar(50);null;comment: '服务器登陆用户名（服务器类型为1时所需要）'"`
	// 服务器登陆密码（服务器类型为0时所需要）
	Password string `json:"password" gorm:"type:varchar(500);null;comment: '服务器登陆密码（服务器类型为1时所需要）'"`
	// 服务器ssh secret（服务器类型为1时所需要）
	PrivateKey string `json:"privateKey" gorm:"type:varchar(2000);null;comment: '服务器ssh secret（服务器类型为2时所需要）'"`
}

package dto

import "updateTool/model"

type ServerDto struct {
	ID uint
	// 服务器名称
	ServerName string `json:"serverName"`
	// 服务器认证类型 1:Password 2:PrivateKey
	ServerType int `json:"serverType"`
	// 服务器主机地址
	Host string `json:"host"`
	// 服务器端口
	Port int `json:"port"`
	// 服务器登陆用户名（服务器类型为0时所需要）
	Username string `json:"username"`
}

func ToServerDto(server model.Server) ServerDto {
	return ServerDto{
		ID:         server.ID,
		ServerName: server.ServerName,
		ServerType: server.ServerType,
		Host:       server.Host,
		Port:       server.Port,
		Username:   server.Username,
	}
}

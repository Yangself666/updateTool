package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"updateTool/common"
	"updateTool/sftp"
	"updateTool/util"
)

func main() {
	// 初始化配置文件
	InitConfig()

	// 获取数据库连接
	common.GetDB()

	// 检查配置中服务器状态
	// CheckServers()

	r := gin.Default()
	// 不记录静态文件日志
	gin.LoggerWithWriter(gin.DefaultWriter, "/assets/*")
	// 加载路由
	r = CollectRoute(r)

	// 获取启动端口
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("resource/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func CheckServers() {
	var serverStr string
	servers := util.GetServers()
	for _, item := range servers {
		client, err := sftp.GetSftpClient(item.Username, item.Password, item.Host, item.Port)
		if err != nil {
			panic(item.Host + " 无法连接，请检查该服务器的参数及配置")
		}
		// 创建连接后首先defer进行关闭操作，防止遗忘
		client.Close()
		serverStr += item.Host + " "
	}
	log.Println("所有服务器配置检查通过 [ " + serverStr + "]")
}

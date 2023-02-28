package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"updateTool/common"
)

func main() {
	// 初始化配置文件
	InitConfig()

	// 获取数据库连接
	common.GetDB()

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

// InitConfig 初始化读取配置文件
func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("resource/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

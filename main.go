package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io"
	"log"
	"os"
	"path"
	"updateTool/common"
)

func main() {
	// 初始化配置文件
	InitConfig()

	// 获取数据库连接
	common.GetDB()

	r := gin.New()
	r.Use(gin.Recovery())
	r = InitLog(r)

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

func InitLog(r *gin.Engine) *gin.Engine {
	var (
		logPath = "logs"
		logName = "updateTool.log"
		file    *os.File
		err     error
	)
	// 文件夹是否存在
	// 文件是否存在
	_, err = os.Stat(logPath)
	if err != nil {
		// 文件夹不存在，进行创建
		os.MkdirAll(logPath, 0644)
	}
	filePath := path.Join(logPath, logName)

	file, err = os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Println("Log File Open Fail")
		r.Use(gin.Logger())
		return r
	}
	r.Use(gin.LoggerWithWriter(io.MultiWriter(file, os.Stdout)))
	return r
}

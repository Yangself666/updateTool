package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"os"
	"updateTool/response"
)

func Reload(c *gin.Context) {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("配置文件重载失败:", err)
		response.Fail(c, nil, "配置文件重载失败，请尝试重启服务")
		return
	}
	response.Success(c, nil, "配置文件重载成功")
}

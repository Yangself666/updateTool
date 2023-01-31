package main

import (
	"github.com/gin-gonic/gin"
	"updateTool/controller"
	"updateTool/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.GET("/", controller.Index)
	r.POST("/uploadFile", controller.UploadFile)
	r.POST("/getAllHistory", controller.GetAllHistory)
	r.POST("/getHistoryByRemotePath", controller.GetHistoryByRemotePath)
	r.POST("/rollback", controller.Rollback)
	r.GET("/reload", controller.Reload)
	return r
}

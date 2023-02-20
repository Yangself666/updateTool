package main

import (
	"github.com/gin-gonic/gin"
	"updateTool/controller"
	"updateTool/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.GET("/", controller.Index)
	r.POST("/login", controller.Login)
	r.POST("/uploadFile", controller.UploadFile)
	r.POST("/getHistory", controller.GetHistory)
	r.POST("/rollback", controller.Rollback)
	r.GET("/reload", controller.Reload)
	return r
}

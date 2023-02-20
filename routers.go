package main

import (
	"github.com/gin-gonic/gin"
	"updateTool/controller"
	"updateTool/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	//r.GET("/", controller.Index)
	r.POST("/login", controller.Login)
	r.POST("/uploadFile", middleware.AuthMiddleware(), controller.UploadFile)
	r.POST("/getHistory", middleware.AuthMiddleware(), controller.GetHistory)
	r.POST("/rollback", middleware.AuthMiddleware(), controller.Rollback)
	//r.GET("/reload", controller.Reload)
	r.StaticFile("/", "index.html")
	return r
}

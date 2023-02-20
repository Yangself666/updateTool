package main

import (
	"github.com/gin-gonic/gin"
	"updateTool/controller"
	"updateTool/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	// 静态文件路由
	r.StaticFile("/", "resource/web/index.html")
	r.Static("/static", "resource/web/static")
	// 解决vue等前端路由问题（gin路由不存在返回首页）
	r.NoRoute(func(c *gin.Context) {
		c.Request.URL.Path = "/"
		r.HandleContext(c)
	})

	// 服务接口
	apiRoutes := r.Group("/api")
	apiRoutes.POST("/login", controller.Login)
	apiRoutes.POST("/uploadFile", middleware.AuthMiddleware(), controller.UploadFile)
	apiRoutes.POST("/getHistory", middleware.AuthMiddleware(), controller.GetHistory)
	apiRoutes.POST("/rollback", middleware.AuthMiddleware(), controller.Rollback)
	apiRoutes.GET("/reload", controller.Reload)

	return r
}

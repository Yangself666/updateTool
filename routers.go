package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
		c.Redirect(http.StatusFound, "/")
	})

	// 服务接口组
	apiRoutes := r.Group("/api")
	// 登陆
	apiRoutes.POST("/login", middleware.ExposeHeaderMiddleware(), controller.Login)

	// 用户相关接口
	userApi := apiRoutes.Group("/user", middleware.AuthMiddleware())
	// 登陆用户获取信息
	userApi.POST("/info", controller.Info)

	// 项目管理相关接口
	projectApi := apiRoutes.Group("/project", middleware.AuthMiddleware())
	// 添加项目信息
	projectApi.POST("/add", controller.AddProject)
	// 删除项目信息
	projectApi.POST("/del", controller.DelProject)
	// 修改项目信息
	projectApi.POST("/edit", controller.EditProject)
	// 获取项目列表
	projectApi.POST("/list", controller.GetProjectList)
	// 获取单个项目信息
	projectApi.POST("/info", controller.GetProjectById)
	// 通过项目ID获取路径信息
	projectApi.POST("/path", controller.GetPathListByProjectId)

	// 服务器管理相关接口
	serverApi := apiRoutes.Group("/server", middleware.AuthMiddleware())
	// 添加服务器信息
	serverApi.POST("/add", controller.AddServer)
	// 删除服务器信息
	serverApi.POST("/del", controller.DelServer)
	// 修改服务器信息
	serverApi.POST("/edit", controller.EditServer)
	// 获取服务器列表
	serverApi.POST("/list", controller.GetServerList)

	// 上传文件
	// apiRoutes.POST("/uploadFile", middleware.AuthMiddleware(), controller.UploadFile)
	// 获取上传历史
	apiRoutes.POST("/getHistory", middleware.AuthMiddleware(), controller.GetHistory)
	// 回滚
	// apiRoutes.POST("/rollback", middleware.AuthMiddleware(), controller.Rollback)
	// 重新读取配置文件
	apiRoutes.GET("/reload", controller.Reload)

	return r
}

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
	r.StaticFile("/vite.svg", "resource/web/vite.svg")
	r.Static("/assets", "resource/web/assets")
	// 解决vue等前端路由问题（gin路由不存在返回首页）
	r.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/")
	})
	r = apiRouter(r)

	return r
}

// 后台接口
func apiRouter(r *gin.Engine) *gin.Engine {
	// 服务接口组
	apiRoutes := r.Group("/api")
	// 登陆
	apiRoutes.POST("/login", middleware.ExposeHeaderMiddleware(), controller.Login)

	// 用户相关接口
	userApi := apiRoutes.Group("/user", middleware.AuthMiddleware())
	// 登陆用户获取信息
	userApi.POST("/info", controller.Info)
	// 添加用户
	userApi.POST("/add", controller.AddUser)
	// 编辑用户信息
	userApi.POST("/edit", controller.EditUser)
	// 修改用户密码
	userApi.POST("/editPassword", controller.EditUserPassword)
	// 删除用户及其关联
	userApi.POST("/del", controller.DelUser)
	// 获取所有用户列表
	userApi.POST("/list", controller.ListUser)

	// 项目管理相关接口
	projectApi := apiRoutes.Group("/project", middleware.AuthMiddleware())
	// 添加项目信息
	projectApi.POST("/add", controller.AddProject)
	// 删除项目信息
	projectApi.POST("/del", controller.DelProject)
	// 删除前检测项目是否可以删除
	projectApi.POST("/delCheck", controller.DelCheckProject)
	// 修改项目信息
	projectApi.POST("/edit", controller.EditProject)
	// 获取项目列表
	projectApi.POST("/list", controller.GetProjectList)
	// 获取单个项目信息
	projectApi.POST("/info", controller.GetProjectById)
	// 通过项目ID获取路径信息
	projectApi.POST("/path", controller.GetPathListByProjectId)
	// 通过用户ID获取绑定的项目
	projectApi.POST("/userId", controller.GetProjectListByUserId)
	// 获取登陆用户绑定的项目
	projectApi.POST("/get", controller.GetProjectListByLoginUser)

	// 项目中的用户管理
	projectUserApi := projectApi.Group("/user", middleware.AuthMiddleware())
	// 添加用户关联
	projectUserApi.POST("/add", controller.AddProjectUser)
	// 删除用户关联
	projectUserApi.POST("/del", controller.DelProjectUser)
	// 批量编辑用户关联
	projectUserApi.POST("/edit", controller.EditProjectUser)
	// 获取项目中的用户列表
	projectUserApi.POST("/list", controller.GetUserListByProjectId)

	// 项目中的路径管理
	projectPathApi := projectApi.Group("/path", middleware.AuthMiddleware())
	// 添加路径信息
	projectPathApi.POST("/add", controller.AddProjectPath)
	// 删除路径信息
	projectPathApi.POST("/del", controller.DelProjectPath)
	// 修改路径信息
	projectPathApi.POST("/edit", controller.EditProjectPath)

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
	// 检查服务器是否可以连接
	serverApi.POST("/check", controller.CheckServer)

	// 权限相关接口
	permissionApi := apiRoutes.Group("/permission", middleware.AuthMiddleware())
	// 添加权限信息
	permissionApi.POST("/add", controller.AddPermission)
	// 删除权限信息
	permissionApi.POST("/del", controller.DelPermission)
	// 修改权限信息
	permissionApi.POST("/edit", controller.EditPermission)
	// 获取权限列表
	permissionApi.POST("/list", controller.GetPermissionList)

	// 上传文件
	uploadApi := apiRoutes.Group("/upload", middleware.AuthMiddleware())
	// 上传文件接口
	uploadApi.POST("/file", controller.UploadFile)

	// 获取上传历史
	historyGroup := apiRoutes.Group("/history", middleware.AuthMiddleware())
	historyGroup.POST("/get", controller.GetHistory)
	// 回滚
	historyGroup.POST("/rollback", controller.Rollback)

	// 重新读取配置文件
	apiRoutes.GET("/reload", controller.Reload)

	return r
}

package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"path"
	"updateTool/common"
	"updateTool/dto"
	"updateTool/model"
	"updateTool/response"
)

/*
项目路径管理Controller
*/

// AddProjectPath 新增项目路径
func AddProjectPath(c *gin.Context) {
	var projectDto = dto.ProjectDto{}
	err := c.BindJSON(&projectDto)
	if err != nil {
		response.Fail(c, nil, "参数不正确")
		return
	}

	if projectDto.ProjectName == "" {
		response.Fail(c, nil, "项目名称不能为空")
		return
	}
	project := dto.ProjectDtoToProject(projectDto)

	DB := common.GetDB()
	DB.Create(&project)

	// 如果填写了服务器就进行添加操作
	if projectDto.ServerIdList != nil && len(projectDto.ServerIdList) > 0 {
		// 查询绑定的服务器
		var servers []model.Server
		DB.Model(&model.Server{}).Where("id in ?", projectDto.ServerIdList).Find(&servers)

		// 添加服务器
		if len(servers) > 0 {
			cons := make([]model.ProjectServerCon, 0)
			for _, server := range servers {
				cons = append(cons, model.ProjectServerCon{ProjectId: project.ID, ServerId: server.ID})
			}
			DB.Create(&cons)
		}
	}

	// 如果填写了服务器路径就进行添加操作
	if projectDto.ProjectPathList != nil && len(projectDto.ProjectPathList) > 0 {
		// 去重
		pathMap := make(map[string]model.ProjectPath)
		for _, projectPath := range projectDto.ProjectPathList {
			cleanPath := path.Clean(projectPath.Path)
			projectPath.Path = cleanPath
			// 放置项目ID
			projectPath.ProjectId = project.ID
			// 放置默认值
			if projectPath.HasSubPath == 0 {
				projectPath.HasSubPath = 1
			}
			// 说明path不存在
			if pathMap[cleanPath].IsEmpty() {
				// 存放路径
				pathMap[cleanPath] = projectPath
			}
		}
		// 创建新的数组
		newPathList := make([]model.ProjectPath, 0)
		for _, value := range pathMap {
			newPathList = append(newPathList, value)
		}

		// 保存路径数组
		DB.Create(&newPathList)
	}

	// 如果填写了用户就进行添加操作
	if projectDto.UserIdList != nil && len(projectDto.UserIdList) > 0 {
		// 查询绑定的服务器
		var users []model.User
		DB.Model(&model.User{}).Where("id in ?", projectDto.UserIdList).Find(&users)

		// 添加用户
		if len(users) > 0 {
			cons := make([]model.ProjectUserCon, 0)
			for _, user := range users {
				cons = append(cons, model.ProjectUserCon{ProjectId: project.ID, UserId: user.ID})
			}
			DB.Create(&cons)
		}
	}

	response.Success(c, nil, "请求成功")
}

// DelProjectPath 删除项目路径
func DelProjectPath(c *gin.Context) {
	var param = make(map[string]int, 0)
	err := c.BindJSON(&param)
	if err != nil {
		log.Println("参数接收发生错误 -> ", err)
		response.Fail(c, nil, "参数不正确")
		return
	}

	if param["id"] == 0 {
		response.Fail(c, nil, "项目ID不能为空")
		return
	}
	projectId := param["id"]

	DB := common.GetDB()
	var count int64
	DB.Model(&model.Project{}).Where("id = ?", projectId).Count(&count)
	if count <= 0 {
		response.Fail(c, nil, "该项目不存在")
		return
	}

	// 删除关联的服务器
	DB.Unscoped().Where("project_id = ?", projectId).Delete([]model.ProjectServerCon{})

	// 删除路径的关联
	DB.Unscoped().Where("project_id = ?", projectId).Delete([]model.ProjectPath{})

	// 删除用户的关联
	DB.Unscoped().Where("project_id = ?", projectId).Delete([]model.ProjectUserCon{})

	// 删除项目
	DB.Where("id = ?", projectId).Delete(model.Project{})

	response.Success(c, nil, "删除成功")
}

// EditProjectPath 编辑项目路径
func EditProjectPath(c *gin.Context) {
	var projectDto = dto.ProjectDto{}
	err := c.BindJSON(&projectDto)
	if err != nil {
		response.Fail(c, nil, "参数不正确")
		return
	}

	if projectDto.ProjectName == "" {
		response.Fail(c, nil, "项目名称不能为空")
		return
	}
	if projectDto.ID == 0 {
		response.Fail(c, nil, "项目ID不能为空")
		return
	}
	project := dto.ProjectDtoToProject(projectDto)

	DB := common.GetDB()
	var count int64
	DB.Model(&model.Project{}).Where("id = ?", project.ID).Count(&count)
	if count <= 0 {
		response.Fail(c, nil, "修改的项目不存在")
		return
	}

	// 删除之前绑定的服务器
	DB.Unscoped().Where("project_id = ?", project.ID).Delete([]model.ProjectServerCon{})

	// 删除之前绑定的路径
	DB.Unscoped().Where("project_id = ?", project.ID).Delete([]model.ProjectPath{})

	// 删除之前绑定的路径
	DB.Unscoped().Where("project_id = ?", project.ID).Delete([]model.ProjectUserCon{})

	// 如果填写了服务器就进行添加操作
	if projectDto.ServerIdList != nil && len(projectDto.ServerIdList) > 0 {
		// 查询绑定的服务器
		var servers []model.Server
		DB.Model(&model.Server{}).Where("id in ?", projectDto.ServerIdList).Find(&servers)

		// 添加服务器
		if len(servers) > 0 {
			cons := make([]model.ProjectServerCon, 0)
			for _, server := range servers {
				cons = append(cons, model.ProjectServerCon{ProjectId: project.ID, ServerId: server.ID})
			}
			DB.Create(&cons)
		}
	}

	// 如果填写了服务器路径就进行添加操作
	if projectDto.ProjectPathList != nil && len(projectDto.ProjectPathList) > 0 {
		// 去重
		pathMap := make(map[string]model.ProjectPath)
		for _, projectPath := range projectDto.ProjectPathList {
			cleanPath := path.Clean(projectPath.Path)
			projectPath.Path = cleanPath
			// 放置项目ID
			projectPath.ProjectId = project.ID
			// 放置默认值
			if projectPath.HasSubPath == 0 {
				projectPath.HasSubPath = 1
			}
			// 说明path不存在
			if pathMap[cleanPath].IsEmpty() {
				// 存放路径
				pathMap[cleanPath] = projectPath
			}
		}
		// 创建新的数组
		newPathList := make([]model.ProjectPath, 0)
		for _, value := range pathMap {
			newPathList = append(newPathList, value)
		}

		// 保存路径数组
		DB.Create(&newPathList)
	}

	// 如果填写了用户就进行添加操作
	if projectDto.UserIdList != nil && len(projectDto.UserIdList) > 0 {
		// 查询绑定的服务器
		var users []model.User
		DB.Model(&model.User{}).Where("id in ?", projectDto.UserIdList).Find(&users)

		// 添加用户
		if len(users) > 0 {
			cons := make([]model.ProjectUserCon, 0)
			for _, user := range users {
				cons = append(cons, model.ProjectUserCon{ProjectId: project.ID, UserId: user.ID})
			}
			DB.Create(&cons)
		}
	}

	DB.Updates(&project)
	response.Success(c, nil, "请求成功")
}

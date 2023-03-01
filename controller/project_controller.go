package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"updateTool/common"
	"updateTool/dto"
	"updateTool/model"
	"updateTool/response"
)

/*
项目管理Controller
*/

// AddProject 新增项目
func AddProject(c *gin.Context) {
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

	/*
		// 如果填写了服务器路径就进行添加操作
		if projectDto.ProjectPathList != nil && len(projectDto.ProjectPathList) > 0 {
			// 去重
			pathMap := make(map[string]model.ProjectPath)
			for _, projectPath := range projectDto.ProjectPathList {
				// 路径检查
				if !strings.HasSuffix(projectPath.Path, "/") {
					projectPath.Path = projectPath.Path + "/"
				}
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
	*/

	response.Success(c, nil, "请求成功")
}

// DelProject 删除项目
func DelProject(c *gin.Context) {
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
	DB.Unscoped().Delete(&model.ProjectServerCon{}, "project_id = ?", projectId)

	// 删除路径的关联
	DB.Unscoped().Delete(&model.ProjectPath{}, "project_id = ?", projectId)

	// 删除用户的关联
	DB.Unscoped().Delete(&model.ProjectUserCon{}, "project_id = ?", projectId)

	// 删除项目
	DB.Delete(&model.Project{}, "id = ?", projectId)

	response.Success(c, nil, "删除成功")
}

// DelCheckProject 删除前检测项目是否可以删除
func DelCheckProject(c *gin.Context) {
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
	var projectCount int64
	DB.Model(&model.Project{}).Where("id = ?", projectId).Count(&projectCount)
	if projectCount <= 0 {
		response.Fail(c, nil, "该项目不存在")
		return
	}

	// 查询项目是否有关联服务器
	var count int64
	DB.Model(&model.ProjectServerCon{}).Where("project_id = ?", projectId).Count(&count)
	if count > 0 {
		response.Success(c, false, "查询成功")
		return
	}
	// 查询项目是否有关联路径
	DB.Model(&model.ProjectPath{}).Where("project_id = ?", projectId).Count(&count)
	if count > 0 {
		response.Success(c, false, "查询成功")
		return
	}
	// 查询项目是否有关联用户
	DB.Model(&model.ProjectUserCon{}).Where("project_id = ?", projectId).Count(&count)
	if count > 0 {
		response.Success(c, false, "查询成功")
		return
	}
	response.Success(c, true, "查询成功")
}

// EditProject 编辑项目
func EditProject(c *gin.Context) {
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
	DB.Unscoped().Delete(&model.ProjectServerCon{}, "project_id = ?", project.ID)
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

	/*
		// 删除之前绑定的路径
		DB.Unscoped().Delete(&model.ProjectPath{}, "project_id = ?", project.ID)

		// 删除之前绑定的用户
		DB.Unscoped().Delete(&model.ProjectUserCon{}, "project_id = ?", project.ID)

		// 如果填写了服务器路径就进行添加操作
		if projectDto.ProjectPathList != nil && len(projectDto.ProjectPathList) > 0 {
			// 去重
			pathMap := make(map[string]model.ProjectPath)
			for _, projectPath := range projectDto.ProjectPathList {
				// 路径检查
				if !strings.HasSuffix(projectPath.Path, "/") {
					projectPath.Path = projectPath.Path + "/"
				}
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
			// 查询绑定的用户
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
	*/

	DB.Updates(&project)
	response.Success(c, nil, "请求成功")
}

// GetProjectList 获取所有项目
func GetProjectList(c *gin.Context) {
	var project = model.Project{}
	err := c.BindJSON(&project)
	if err != nil {
		log.Println("参数解析失败 -> ", err)
		response.Fail(c, nil, "参数不正确")
		return
	}

	DB := common.GetDB()
	tx := DB.Model(&model.Project{})
	var projectList = make([]model.Project, 0)

	if project.ProjectName != "" {
		tx.Where("project_name like ?", "%"+project.ProjectName+"%")
	}
	if project.ID != 0 {
		tx.Where("id = ?", project.ID)
	}
	tx.Find(&projectList)
	response.Success(c, projectList, "请求成功")
}

// GetProjectById 通过ID查询单个项目信息
func GetProjectById(c *gin.Context) {
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

	// 查询项目
	DB := common.GetDB()
	var project model.Project
	DB.Where("id = ?", projectId).First(&project)
	if project.ID == 0 {
		// 未查询出项目
		response.Fail(c, nil, "该项目不存在")
		return
	}
	projectDto := dto.ToProjectDto(project)

	// 查询关联路径
	paths := make([]model.ProjectPath, 0)
	DB.Where("project_id = ?", projectId).Find(&paths)
	projectDto.ProjectPathList = paths

	// 查询关联服务器
	var servers []model.Server
	DB.Model(&model.ProjectServerCon{}).Select("servers.*").Joins("left join servers on project_server_cons.server_id = servers.id").Where("project_id = ?", projectId).Find(&servers)
	projectDto.ServerList = servers

	response.Success(c, projectDto, "请求成功")
}

// GetPathListByProjectId 通过项目ID获取路径列表
func GetPathListByProjectId(c *gin.Context) {
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

	// 查询项目下的路径
	pathList := make([]model.ProjectPath, 0)
	DB.Model(&model.ProjectPath{}).Where("project_id = ?", projectId).Find(&pathList)

	response.Success(c, pathList, "请求成功")
}

// GetProjectListByUserId 通过用户ID获取项目列表(管理员)
func GetProjectListByUserId(c *gin.Context) {
	var param = make(map[string]int, 0)
	err := c.BindJSON(&param)
	if err != nil {
		log.Println("参数接收发生错误 -> ", err)
		response.Fail(c, nil, "参数不正确")
		return
	}

	if param["id"] == 0 {
		response.Fail(c, nil, "用户ID不能为空")
		return
	}
	userId := param["id"]
	projectDtoList, err := getProjectListByUserIdService(userId)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, projectDtoList, "请求成功")
}

// GetProjectListByLoginUser 获取登陆用户绑定的项目
func GetProjectListByLoginUser(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		response.Unauthorized(c)
		return
	}
	// 断言类型
	userId := (int)(user.(model.User).ID)
	projectDtoList, err := getProjectListByUserIdService(userId)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, projectDtoList, "请求成功")
}

// 通过用户ID获取项目列表
func getProjectListByUserIdService(userId int) ([]dto.ProjectDto, error) {
	DB := common.GetDB()
	var count int64
	DB.Model(&model.User{}).Where("id = ?", userId).Count(&count)
	if count <= 0 {
		return nil, common.Error("该用户不存在")
	}

	// 查询用户绑定的路径
	projectIds := DB.Select("project_id").Where("user_id = ?", userId).Table("project_user_cons")
	projectList := make([]model.Project, 0)
	DB.Model(&model.Project{}).Where("id in (?)", projectIds).Find(&projectList)

	projectDtoList := make([]dto.ProjectDto, 0)
	for _, project := range projectList {
		projectDto := dto.ToProjectDto(project)
		// 查询关联路径
		paths := make([]model.ProjectPath, 0)
		DB.Where("project_id = ?", projectDto.ID).Find(&paths)
		projectDto.ProjectPathList = paths
		projectDtoList = append(projectDtoList, projectDto)
	}
	return projectDtoList, nil
}

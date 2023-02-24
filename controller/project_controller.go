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
	DB.Unscoped().Where("project_id = ?", projectId).Delete([]model.ProjectServerCon{})

	// 删除路径的关联
	DB.Unscoped().Where("project_id = ?", projectId).Delete([]model.ProjectPath{})

	// 删除项目
	DB.Where("id = ?", projectId).Delete(model.Project{})

	response.Success(c, nil, "删除成功")
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
	DB.Unscoped().Where("project_id = ?", project.ID).Delete([]model.ProjectServerCon{})

	// 删除之前绑定的路径
	DB.Unscoped().Where("project_id = ?", project.ID).Delete([]model.ProjectPath{})

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
	var projectList []model.Project

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

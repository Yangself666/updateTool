package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"updateTool/common"
	"updateTool/model"
	"updateTool/response"
)

/*
项目管理Controller
*/

// AddProject 新增项目
func AddProject(c *gin.Context) {
	var project = model.Project{}
	err := c.BindJSON(&project)
	if err != nil {
		response.Fail(c, nil, "参数不正确")
		return
	}

	if project.ProjectName == "" {
		response.Fail(c, nil, "项目名称不能为空")
		return
	}

	DB := common.GetDB()
	DB.Create(&project)

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

	DB := common.GetDB()
	var count int64
	DB.Model(&model.Project{}).Where("id = ?", param["id"]).Count(&count)
	if count <= 0 {
		response.Fail(c, nil, "该项目不存在")
		return
	}

	DB.Model(model.Project{}).Delete("id = ?", param["id"])

	// todo	需要删除路径关联

	response.Success(c, nil, "删除成功")
}

// EditProject 编辑项目
func EditProject(c *gin.Context) {
	var project = model.Project{}
	err := c.BindJSON(&project)
	if err != nil {
		response.Fail(c, nil, "参数不正确")
		return
	}

	if project.ProjectName == "" {
		response.Fail(c, nil, "项目名称不能为空")
		return
	}
	if project.ID == 0 {
		response.Fail(c, nil, "项目ID不能为空")
		return
	}

	DB := common.GetDB()
	var count int64
	DB.Model(model.Project{}).Where("id = ?", project.ID).Count(&count)
	if count <= 0 {
		response.Fail(c, nil, "修改的项目不存在")
		return
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

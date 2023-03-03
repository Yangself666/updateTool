package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"path"
	"strings"
	"updateTool/common"
	"updateTool/model"
	"updateTool/response"
)

/*
项目路径管理Controller
*/

// AddProjectPath 新增项目路径
func AddProjectPath(c *gin.Context) {
	var projectPath = model.ProjectPath{}
	err := c.BindJSON(&projectPath)
	if err != nil {
		response.Fail(c, nil, "参数不正确")
		return
	}

	if projectPath.ProjectId == 0 ||
		projectPath.HasSubPath == 0 ||
		projectPath.Path == "" ||
		projectPath.PathName == "" {
		response.Fail(c, nil, "参数不完整")
		return
	}
	// 路径检查
	if !strings.HasPrefix(projectPath.Path, "/") {
		projectPath.Path = "/" + projectPath.Path
	}
	projectPath.Path = path.Clean(projectPath.Path)

	DB := common.GetDB()
	var count int64
	DB.Model(&model.Project{}).Where("id = ?", projectPath.ProjectId).Count(&count)
	if count <= 0 {
		response.Fail(c, nil, "该项目不存在")
		return
	}

	// 检查路径是否存在
	var pathCount int64
	DB.Model(&model.ProjectPath{}).Where(
		"project_id = ? and path = ?",
		projectPath.ProjectId, path.Clean(projectPath.Path)).Count(&pathCount)
	if pathCount > 0 {
		response.Fail(c, nil, "该路径已存在")
		return
	}
	// 添加路径
	DB.Create(&projectPath)

	response.Success(c, nil, "添加成功")
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
		response.Fail(c, nil, "项目路径ID不能为空")
		return
	}
	pathId := param["id"]

	DB := common.GetDB()
	var count int64
	DB.Model(&model.ProjectPath{}).Where("id = ?", pathId).Count(&count)
	if count <= 0 {
		response.Fail(c, nil, "该项目路径不存在")
		return
	}

	// 删除路径的关联
	DB.Delete(&model.ProjectPath{}, "id = ?", pathId)

	response.Success(c, nil, "删除成功")
}

// EditProjectPath 编辑项目路径
func EditProjectPath(c *gin.Context) {
	var projectPath = model.ProjectPath{}
	err := c.BindJSON(&projectPath)
	if err != nil {
		response.Fail(c, nil, "参数不正确")
		return
	}

	if projectPath.ID == 0 ||
		projectPath.ProjectId == 0 ||
		projectPath.HasSubPath == 0 ||
		projectPath.Path == "" ||
		projectPath.PathName == "" {
		response.Fail(c, nil, "参数不完整")
		return
	}
	// 路径检查
	if !strings.HasPrefix(projectPath.Path, "/") {
		projectPath.Path = "/" + projectPath.Path
	}
	projectPath.Path = path.Clean(projectPath.Path)

	DB := common.GetDB()
	var count int64
	// 查询项目路径是否存在
	DB.Model(&model.ProjectPath{}).Where("id = ?", projectPath.ID).Count(&count)
	if count <= 0 {
		response.Fail(c, nil, "该项目路径信息不存在")
		return
	}

	// 查询项目是否存在
	DB.Model(&model.Project{}).Where("id = ?", projectPath.ProjectId).Count(&count)
	if count <= 0 {
		response.Fail(c, nil, "该项目不存在")
		return
	}

	// 检查路径是否存在
	var pathCount int64
	DB.Model(&model.ProjectPath{}).Where(
		"id <> ? and project_id = ? and path = ?", projectPath.ID,
		projectPath.ProjectId, path.Clean(projectPath.Path)).Count(&pathCount)
	if pathCount > 0 {
		response.Fail(c, nil, "该路径已存在")
		return
	}
	// 添加路径
	DB.Updates(&projectPath)

	response.Success(c, nil, "修改成功")
}

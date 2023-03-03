package controller

import (
	"github.com/gin-gonic/gin"
	"updateTool/common"
	"updateTool/dto"
	"updateTool/model"
	"updateTool/response"
)

/*
项目关联服务器管理Controller
*/

// EditProjectServer 批量修改项目关联服务器
func EditProjectServer(c *gin.Context) {
	type Param struct {
		ProjectId    uint
		ServerIdList []uint
	}
	param := Param{}
	err := c.BindJSON(&param)
	if err != nil {
		response.Fail(c, nil, "参数不正确")
		return
	}

	if param.ProjectId == 0 {
		response.Fail(c, nil, "参数不正确")
		return
	}
	DB := common.GetDB()
	var projectCount int64
	// 查询项目是否存在
	DB.Model(&model.Project{}).Where("id = ?", param.ProjectId).Count(&projectCount)
	if projectCount <= 0 {
		response.Fail(c, nil, "该项目不存在")
		return
	}
	// 删除所有项目关联服务器
	DB.Unscoped().Delete(&model.ProjectServerCon{}, "project_id = ?", param.ProjectId)
	// 如果传入了新的服务器ID
	if len(param.ServerIdList) > 0 {
		// 查询存在的服务器
		var servers []model.Server
		DB.Model(&model.Server{}).Where("id in ?", param.ServerIdList).Find(&servers)

		// 添加服务器
		if len(servers) > 0 {
			cons := make([]model.ProjectServerCon, 0)
			for _, server := range servers {
				cons = append(cons, model.ProjectServerCon{ProjectId: param.ProjectId, ServerId: server.ID})
			}
			DB.Create(&cons)
		}
	}
	response.Success(c, nil, "请求成功")
}

// GetServerListByProjectId 通过项目ID获取服务器列表
func GetServerListByProjectId(c *gin.Context) {
	var project model.Project
	err := c.BindJSON(&project)
	if err != nil {
		response.Fail(c, nil, "参数不正确")
		return
	}
	if project.ID == 0 {
		response.Fail(c, nil, "参数不完整")
		return
	}

	DB := common.GetDB()
	var projectCount int64
	// 查询项目是否存在
	DB.Model(&model.Project{}).Where("id = ?", project.ID).Count(&projectCount)
	if projectCount <= 0 {
		response.Fail(c, nil, "该项目不存在")
		return
	}

	// 查询项目绑定服务器信息
	serverList := make([]model.Server, 0)
	DB.Select("servers.*").Model(&model.ProjectServerCon{}).Joins("left join servers on project_server_cons.server_id = servers.id").Where("project_server_cons.project_id = ?", project.ID).Find(&serverList)
	var serverDtoList []dto.ServerDto
	for _, server := range serverList {
		serverDto := dto.ToServerDto(server)
		serverDtoList = append(serverDtoList, serverDto)
	}

	response.Success(c, serverDtoList, "请求成功")
}

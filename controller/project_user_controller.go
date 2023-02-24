package controller

import (
	"github.com/gin-gonic/gin"
	"updateTool/common"
	"updateTool/model"
	"updateTool/response"
)

/*
项目关联用户管理Controller
*/

// AddProjectUser 新增项目关联用户
func AddProjectUser(c *gin.Context) {
	type Param struct {
		ProjectId int
		UserId    int
	}
	param := Param{}
	err := c.BindJSON(&param)
	if err != nil {
		response.Fail(c, nil, "参数不正确")
		return
	}

	if param.ProjectId == 0 || param.UserId == 0 {
		response.Fail(c, nil, "参数不完整")
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
	// 查询用户是否存在
	var userCount int64
	DB.Model(&model.ProjectUserCon{}).Where("project_id = ? and user_id = ?", param.ProjectId, param.UserId).Count(&userCount)

	if userCount > 0 {
		response.Fail(c, nil, "该用户已与项目绑定")
		return
	}

	// 开始关联
	DB.Create(&model.ProjectUserCon{ProjectId: uint(param.ProjectId), UserId: uint(param.UserId)})

	response.Success(c, nil, "请求成功")
}

// DelProjectUser 删除项目关联用户
func DelProjectUser(c *gin.Context) {
	// var param = make(map[string]int, 0)
	// err := c.BindJSON(&param)
	// if err != nil {
	// 	log.Println("参数接收发生错误 -> ", err)
	// 	response.Fail(c, nil, "参数不正确")
	// 	return
	// }
	//
	// if param["id"] == 0 {
	// 	response.Fail(c, nil, "项目ID不能为空")
	// 	return
	// }
	// projectId := param["id"]
	//
	// DB := common.GetDB()
	// var count int64
	// DB.Model(&model.Project{}).Where("id = ?", projectId).Count(&count)
	// if count <= 0 {
	// 	response.Fail(c, nil, "该项目不存在")
	// 	return
	// }
	//
	// // 删除关联的服务器
	// DB.Unscoped().Where("project_id = ?", projectId).Delete([]model.ProjectServerCon{})
	//
	// // 删除路径的关联
	// DB.Unscoped().Where("project_id = ?", projectId).Delete([]model.ProjectPath{})
	//
	// // 删除用户的关联
	// DB.Unscoped().Where("project_id = ?", projectId).Delete([]model.ProjectUserCon{})
	//
	// // 删除项目
	// DB.Where("id = ?", projectId).Delete(model.Project{})
	//
	// response.Success(c, nil, "删除成功")
}

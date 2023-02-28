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
	DB.Model(&model.User{}).Where("id = ?", param.UserId).Count(&userCount)

	if userCount <= 0 {
		response.Fail(c, nil, "该用户不存在")
		return
	}

	// 查询用户是否已经关联存在
	var userConCount int64
	DB.Model(&model.ProjectUserCon{}).Where("project_id = ? and user_id = ?", param.ProjectId, param.UserId).Count(&userConCount)

	if userConCount > 0 {
		response.Fail(c, nil, "该用户已与项目绑定")
		return
	}

	// 开始关联
	DB.Create(&model.ProjectUserCon{ProjectId: uint(param.ProjectId), UserId: uint(param.UserId)})

	response.Success(c, nil, "请求成功")
}

// DelProjectUser 删除项目关联用户
func DelProjectUser(c *gin.Context) {
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

	// 查询用户是否已经关联存在
	var userConCount int64
	DB.Model(&model.ProjectUserCon{}).Where("project_id = ? and user_id = ?", param.ProjectId, param.UserId).Count(&userConCount)

	if userConCount <= 0 {
		response.Fail(c, nil, "该用户未与项目绑定")
		return
	}
	// 删除用户的关联
	DB.Unscoped().Where("project_id = ? and user_id = ?", param.ProjectId, param.UserId).Delete(model.ProjectUserCon{})

	response.Success(c, nil, "删除成功")
}

// EditProjectUser 批量修改项目关联用户
func EditProjectUser(c *gin.Context) {
	type Param struct {
		ProjectId  uint
		UserIdList []uint
	}
	param := Param{}
	c.BindJSON(&param)

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
	// 删除所有项目关联用户
	DB.Unscoped().Where("project_id = ?", param.ProjectId).Delete([]model.ProjectUserCon{})
	// 如果传入了新的用户ID
	if len(param.UserIdList) > 0 {
		// 查询存在的用户
		var users []model.User
		DB.Model(&model.User{}).Where("id in ?", param.UserIdList).Find(&users)

		// 添加用户
		if len(users) > 0 {
			cons := make([]model.ProjectUserCon, 0)
			for _, user := range users {
				cons = append(cons, model.ProjectUserCon{ProjectId: param.ProjectId, UserId: user.ID})
			}
			DB.Create(&cons)
		}
	}
	response.Fail(c, nil, "请求成功")
}

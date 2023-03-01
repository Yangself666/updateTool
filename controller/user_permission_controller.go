package controller

import (
	"github.com/gin-gonic/gin"
	"updateTool/common"
	"updateTool/model"
	"updateTool/response"
)

// 用户权限Controller

// EditUserPermission 批量编辑权限
func EditUserPermission(c *gin.Context) {
	type Param struct {
		UserId           uint
		PermissionIdList []uint
	}
	param := Param{}
	err := c.BindJSON(&param)
	if err != nil {
		response.Fail(c, nil, "参数不正确")
		return
	}

	if param.UserId == 0 {
		response.Fail(c, nil, "参数不正确")
		return
	}
	DB := common.GetDB()
	var userCount int64
	// 查询用户是否存在
	DB.Model(&model.User{}).Where("id = ?", param.UserId).Count(&userCount)
	if userCount <= 0 {
		response.Fail(c, nil, "该用户不存在")
		return
	}
	// 删除所有用户关联权限
	DB.Unscoped().Delete(&model.UserPermissionCon{}, "user_id = ?", param.UserId)
	// 如果传入了新的权限ID
	if len(param.PermissionIdList) > 0 {
		// 查询存在的权限
		var permissions []model.Permission
		DB.Model(&model.Permission{}).Where("id in ?", param.PermissionIdList).Find(&permissions)

		// 添加权限
		if len(permissions) > 0 {
			cons := make([]model.UserPermissionCon, 0)
			for _, permission := range permissions {
				cons = append(cons, model.UserPermissionCon{UserId: param.UserId, PermissionId: permission.ID})
			}
			DB.Create(&cons)
		}
	}
	response.Success(c, nil, "请求成功")
}

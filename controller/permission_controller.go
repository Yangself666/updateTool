package controller

import (
	"github.com/gin-gonic/gin"
	"updateTool/common"
	"updateTool/model"
	"updateTool/response"
)

// 权限Controller

// AddPermission 添加权限
func AddPermission(c *gin.Context) {
	var permission model.Permission
	err := c.BindJSON(&permission)
	if err != nil {
		response.Fail(c, nil, "参数不正确")
		return
	}
	// 检查参数
	if permission.PermissionName == "" || permission.PermissionPath == "" || permission.MenuName == "" {
		response.Fail(c, nil, "参数不完整")
		return
	}
	DB := common.GetDB()
	var count int64
	DB.Model(&model.Permission{}).Where(
		"permission_name = ?",
		permission.PermissionName).Or(
		"permission_path = ?",
		permission.PermissionPath).Count(&count)
	if count > 0 {
		response.Fail(c, nil, "权限名或权限路径已存在")
		return
	}

	// 添加权限
	DB.Create(&permission)
	response.Success(c, nil, "请求成功")
}

// EditPermission 修改权限
func EditPermission(c *gin.Context) {
	var permission model.Permission
	err := c.BindJSON(&permission)
	if err != nil {
		response.Fail(c, nil, "参数不正确")
		return
	}
	// 检查参数
	if permission.ID == 0 || permission.PermissionName == "" || permission.PermissionPath == "" || permission.MenuName == "" {
		response.Fail(c, nil, "参数不完整")
		return
	}
	DB := common.GetDB()
	var count int64
	// 检查权限是否存在
	DB.Model(&model.Permission{}).Where("id = ?", permission.ID).Count(&count)
	if count <= 0 {
		response.Fail(c, nil, "该权限不存在")
		return
	}
	DB.Model(&model.Permission{}).Where(
		"id <> ? and (permission_name = ? or permission_path = ?)",
		permission.ID, permission.PermissionName, permission.PermissionPath).Count(&count)
	if count > 0 {
		response.Fail(c, nil, "权限名或权限路径已存在")
		return
	}

	// 添加权限
	DB.Updates(&permission)
	response.Success(c, nil, "请求成功")
}

// DelPermission 删除权限
func DelPermission(c *gin.Context) {
	var permission model.Permission
	err := c.BindJSON(&permission)
	if err != nil {
		response.Fail(c, nil, "参数不正确")
		return
	}
	// 检查参数
	if permission.ID == 0 {
		response.Fail(c, nil, "参数不完整")
		return
	}
	DB := common.GetDB()
	var count int64
	// 检查权限是否存在
	DB.Model(&model.Permission{}).Where("id = ?", permission.ID).Count(&count)
	if count <= 0 {
		response.Fail(c, nil, "该权限不存在")
		return
	}
	// 删除关联表
	DB.Unscoped().Delete(&model.UserPermissionCon{}, "permission_id = ?", permission.ID)

	// 删除权限
	DB.Delete(&model.Permission{}, "id = ?", permission.ID)
	response.Success(c, nil, "请求成功")
}

// GetPermissionList 获取权限列表
func GetPermissionList(c *gin.Context) {

}

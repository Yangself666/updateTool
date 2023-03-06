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
	err := c.ShouldBindJSON(&param)
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

// GetPermissionListByLoginUser 获取登陆用户自己的权限
func GetPermissionListByLoginUser(c *gin.Context) {
	user, exists := c.Get("user")
	var userId uint
	if !exists {
		response.Unauthorized(c)
		return
	} else {
		userId = user.(model.User).ID
		if userId == 0 {
			response.Fail(c, nil, "用户未登陆")
			return
		}
	}

	DB := common.GetDB()
	var userInfo model.User
	// 查询用户是否存在
	DB.Model(&model.User{}).Where("id = ?", userId).First(&userInfo)
	if userInfo.ID == 0 {
		response.Fail(c, nil, "该用户不存在")
		return
	}

	var (
		permissionList = make([]model.Permission, 0)
		menuList       = make([]string, 0)
	)
	if userInfo.IsAdmin {
		// 获取所有权限
		DB.Model(&model.Permission{}).Find(&permissionList)
		// 获取全部菜单列表
		DB.Select("menu_name").Model(&model.Permission{}).Group("menu_name").Find(&menuList)
	} else {
		// 获取权限
		permissionList = GetPermissionListByUserId(userId)
		// 获取菜单列表
		DB.Select("menu_name").Model(&model.UserPermissionCon{}).Joins("left join permissions on user_permission_cons.permission_id = permissions.id").Where("user_permission_cons.user_id = ?", userId).Group("permissions.menu_name").Find(&menuList)
	}

	response.Success(c, gin.H{"menuList": menuList, "permissionList": permissionList}, "请求成功")
}

// GetPermissionListByUser 通过用户ID获取权限列表
func GetPermissionListByUser(c *gin.Context) {
	type Param struct {
		ID uint
	}
	var param Param
	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.Fail(c, nil, "参数不正确")
		return
	}
	if param.ID == 0 {
		response.Fail(c, nil, "参数不完整")
		return
	}
	DB := common.GetDB()
	var user model.User
	// 查询用户是否存在
	DB.Model(&model.User{}).Where("id = ?", param.ID).First(&user)
	if user.ID == 0 {
		response.Fail(c, nil, "该用户不存在")
		return
	}

	// 权限列表
	var permissionList = make([]model.Permission, 0)
	if user.IsAdmin {
		DB.Model(&model.Permission{}).Find(&permissionList)
	} else {
		// 查询权限列表
		permissionList = GetPermissionListByUserId(param.ID)
	}

	response.Success(c, permissionList, "请求成功")
}

// GetPermissionListByUserId 通过用户ID获取权限列表
func GetPermissionListByUserId(userId uint) []model.Permission {
	var permissionList = make([]model.Permission, 0)
	if userId != 0 {
		DB := common.GetDB()
		DB.Model(&model.UserPermissionCon{}).Select("permissions.*").Joins("left join permissions on user_permission_cons.permission_id = permissions.id").Where("user_permission_cons.user_id = ?", userId).Find(&permissionList)
	}

	return permissionList
}

// HasPermissionByPath 通过用户ID和访问路径查询是否有权限
func HasPermissionByPath(userId uint, path string) bool {
	result := false
	if userId == 0 || path == "" {
		return result
	}
	DB := common.GetDB()
	var count int64
	DB.Model(&model.UserPermissionCon{}).Select("permissions.*").Joins("left join permissions on user_permission_cons.permission_id = permissions.id").Where("user_permission_cons.user_id = ? and permissions.permission_path = ?", userId, path).Count(&count)
	// 如果包含该权限
	if count > 0 {
		result = true
	}
	return result
}

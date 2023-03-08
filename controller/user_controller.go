package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"updateTool/common"
	"updateTool/dto"
	"updateTool/model"
	"updateTool/response"
)

/*
用户管理Controller
*/

// Info 获取用户信息
func Info(c *gin.Context) {
	user, exists := c.Get("user")
	userId := user.(model.User).ID
	if !exists || userId == 0 {
		response.Unauthorized(c)
		return
	}
	response.Success(c, dto.ToUserDto(user.(model.User)), "请求成功")
}

// ListUser 获取用户列表
func ListUser(c *gin.Context) {
	var user = model.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.Fail(c, nil, "参数不正确")
		return
	}
	id := user.ID
	name := user.Name
	email := user.Email

	DB := common.GetDB()
	var userDtoList = make([]dto.UserDto, 0)
	tx := DB.Model(&model.User{})
	if id != 0 {
		tx.Where("id = ?", id)
	}

	if name != "" {
		tx.Where("name like ?", "%"+name+"%")
	}

	if email != "" {
		tx.Where("email like ?", "%"+email+"%")
	}
	tx.Find(&userDtoList)

	response.Success(c, userDtoList, "请求成功")
}

// AddUser 添加用户
func AddUser(c *gin.Context) {
	var user = model.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.Fail(c, nil, "参数不正确")
		return
	}
	name := user.Name
	email := user.Email
	password := user.Password

	// 检查参数是否传递
	if name == "" || email == "" || password == "" {
		response.Fail(c, nil, "参数不完整")
		return
	}

	// 检查密码位数
	if len(password) < 6 {
		response.Fail(c, nil, "密码需要大于等于6位")
		return
	}

	// 检查邮箱地址
	DB := common.GetDB()
	var count int64
	DB.Model(&model.User{}).Where("email = ?", email).Count(&count)

	if count > 0 {
		response.Fail(c, nil, "该邮箱地址已被使用")
		return
	}

	// 新建用户
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Fail(c, nil, "密码加密错误，添加失败")
		return
	}
	// 设置密码
	user.Password = string(hashPassword)

	// 添加用户
	DB.Omit("is_admin").Create(&user)

	response.Success(c, nil, "添加成功")
}

// EditUser 修改用户
func EditUser(c *gin.Context) {
	var user = model.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.Fail(c, nil, "参数不正确")
		return
	}

	// 检查参数是否传递
	if user.ID == 0 || user.Name == "" || user.Email == "" {
		response.Fail(c, nil, "参数不完整")
		return
	}

	// 检查邮件地址使用情况
	DB := common.GetDB()
	var count int64
	DB.Model(&model.User{}).Where("email = ? and id <> ?", user.Email, user.ID).Count(&count)
	if count > 0 {
		response.Fail(c, nil, "该邮箱地址已被使用")
		return
	}

	// 修改用户
	DB.Omit("is_admin").Model(&model.User{}).Select("name", "email").Where("id = ?", user.ID).Updates(&user)

	response.Success(c, nil, "修改成功")
}

// SetUserAsAdmin 设置为管理员
func SetUserAsAdmin(c *gin.Context) {
	type Param struct {
		ID uint
	}
	var param Param
	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.Fail(c, nil, "参数不正确")
		return
	}

	// 检查用户是否存在
	DB := common.GetDB()
	var count int64
	DB.Model(&model.User{}).Where("id = ?", param.ID).Count(&count)
	if count <= 0 {
		response.Fail(c, nil, "该用户不存在")
		return
	}

	// 设置用户为管理员
	DB.Model(&model.User{}).Where("id = ?", param.ID).Update("is_admin", true)

	response.Success(c, nil, "请求成功")
}

// SetUserAsNonAdmin 设置为非管理员
func SetUserAsNonAdmin(c *gin.Context) {
	type Param struct {
		ID uint
	}
	var param Param
	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.Fail(c, nil, "参数不正确")
		return
	}

	// 不能设置自己为非管理员
	// 查询登陆用户是否为传入的用户
	user, exists := c.Get("user")
	userId := user.(model.User).ID
	if !exists || userId == 0 {
		response.Unauthorized(c)
		return
	}

	if userId == param.ID {
		response.Fail(c, nil, "不能设置自己为非管理员")
		return
	}

	// 检查用户是否存在
	DB := common.GetDB()
	var count int64
	DB.Model(&model.User{}).Where("id = ?", param.ID).Count(&count)
	if count <= 0 {
		response.Fail(c, nil, "该用户不存在")
		return
	}

	// 设置用户为非管理员
	DB.Model(&model.User{}).Where("id = ?", param.ID).Update("is_admin", false)

	response.Success(c, nil, "请求成功")
}

// EditUserPassword 修改用户密码
func EditUserPassword(c *gin.Context) {
	var user = model.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.Fail(c, nil, "参数不正确")
		return
	}

	// 检查参数是否传递
	if user.ID == 0 || user.Password == "" {
		response.Fail(c, nil, "参数不完整")
		return
	}

	// 检查密码位数
	if len(user.Password) < 6 {
		response.Fail(c, nil, "密码需要大于等于6位")
		return
	}

	// 检查用户是否存在
	DB := common.GetDB()
	var count int64
	DB.Model(&model.User{}).Where("id = ?", user.ID).Count(&count)
	if count <= 0 {
		response.Fail(c, nil, "该用户不存在")
		return
	}

	// 修改用户密码
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		response.Fail(c, nil, "密码加密错误，添加失败")
		return
	}
	// 设置密码
	user.Password = string(hashPassword)

	// 修改密码
	DB.Model(&model.User{}).Select("password").Where("id = ?", user.ID).Updates(&user)

	// 删除缓存中的数据
	cache := common.GetCache()
	cache.Delete(string(user.ID))

	response.Success(c, nil, "修改成功")
}

// DelUser 删除用户
func DelUser(c *gin.Context) {
	var user = model.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.Fail(c, nil, "参数不正确")
		return
	}

	// 检查参数是否传递
	if user.ID == 0 {
		response.Fail(c, nil, "参数不完整")
		return
	}

	// 删除用户关联项目
	DB := common.GetDB()
	DB.Unscoped().Delete(&model.ProjectUserCon{}, "user_id = ?", user.ID)

	// 删除用户
	DB.Delete(&model.User{}, "id = ?", user.ID)

	// 删除缓存中的数据
	cache := common.GetCache()
	cache.Delete(string(user.ID))

	response.Success(c, nil, "删除成功")
}

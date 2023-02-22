package controller

import (
	"github.com/gin-gonic/gin"
	"updateTool/dto"
	"updateTool/model"
	"updateTool/response"
)

/*
用户管理Controller
*/

// Info 获取用户信息
func Info(c *gin.Context) {
	user, _ := c.Get("user")
	response.Success(c, dto.ToUserDto(user.(model.User)), "请求成功")
}

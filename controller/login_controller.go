package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"updateTool/common"
	"updateTool/model"
	"updateTool/response"
)

/*
登陆相关Controller
*/

// Login 用户登陆
func Login(c *gin.Context) {
	// 获取登陆参数中的邮件地址和密码
	// 获取json
	param := make(map[string]string)
	c.BindJSON(&param)
	email := param["email"]
	password := param["password"]

	// 检查参数是否传递
	if email == "" || password == "" {
		response.Fail(c, nil, "参数不完整")
		return
	}

	// 首先通过邮件地址查找是否有该用户
	DB := common.GetDB()
	var userByEmail = model.User{}

	// 使用邮箱查询用户是否存在
	DB.First(&userByEmail, model.User{Email: email})
	if userByEmail.ID == 0 {
		response.Fail(c, nil, "用户不存在")
		return
	}

	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(userByEmail.Password), []byte(password)); err != nil {
		response.Response(c, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}

	// 密码校验通过，发放Token
	token, err := common.ReleaseToken(userByEmail)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error : %v", err)
		return
	}

	// 将Token放置到Header中返回
	c.Header("Authorization", "Bearer "+token)
	// 返回结果
	response.Success(c, nil, "登录成功")
	return
}

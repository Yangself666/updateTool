package dto

import "updateTool/model"

type UserDto struct {
	ID uint
	// 用户名
	Name string `json:"name"`
	// 用户登陆邮箱
	Email string `json:"email"`
	// 是否为管理员
	IsAdmin bool `json:"isAdmin"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		IsAdmin: user.IsAdmin,
	}
}

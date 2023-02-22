package dto

import "updateTool/model"

type UserDto struct {
	// 用户名
	Name string `json:"name"`
	// 用户登陆邮箱
	Email string `json:"email"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:  user.Name,
		Email: user.Email,
	}
}

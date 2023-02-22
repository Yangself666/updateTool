package test

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestLogin(t *testing.T) {
	// 生成密码
	password, _ := bcrypt.GenerateFromPassword([]byte("wxn123456"), bcrypt.DefaultCost)
	fmt.Println(string(password))
}

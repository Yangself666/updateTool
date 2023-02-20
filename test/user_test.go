package test

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestLogin(t *testing.T) {
	password, _ := bcrypt.GenerateFromPassword([]byte("testtest"), bcrypt.DefaultCost)
	fmt.Println(string(password))
}

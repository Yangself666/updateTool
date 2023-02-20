package common

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"updateTool/model"
)

// 设置密钥
var jwtKey = []byte("update_tool_secret")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// ReleaseToken 生成Token
func ReleaseToken(user model.User) (string, error) {
	// 过期时间设置为7天
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "updateTool",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken 解析Token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})

	return token, claims, err
}

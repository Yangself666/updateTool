package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"strings"
	"updateTool/common"
	"updateTool/model"
	"updateTool/response"
)

// AuthMiddleware 权限验证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取authorization header
		tokenString := c.GetHeader("Authorization")

		log.Println(tokenString)

		// 检查token格式是否正确
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			response.Unauthorized(c)
			c.Abort()
			return
		}

		tokenString = tokenString[7:]

		// 解析Token
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			response.Unauthorized(c)
			c.Abort()
			return
		}

		// 验证通过后获取claim中的userId
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		// 用户
		if user.ID == 0 {
			response.Unauthorized(c)
			c.Abort()
			return
		}

		// 用户存在 将user的信息写入上下文
		c.Set("user", user)

		c.Next()
	}
}

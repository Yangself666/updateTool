package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"updateTool/common"
	"updateTool/controller"
	"updateTool/model"
	"updateTool/response"
)

// NoAuthPathMap 不要权限校验的接口
var NoAuthPathMap = map[string]interface{}{
	"/api/user/info":       nil,
	"/api/user/permission": nil,
}

// AuthMiddleware 权限验证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取authorization header
		tokenString := c.GetHeader("Authorization")

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
		_, exists := NoAuthPathMap[c.FullPath()]
		// 如果用户ID不是1，并且接口不再跳过列表中，检查用户权限
		if user.ID != 1 && !exists {
			// 不再无需校验权限接口中
			// 查询是否有对应权限
			hasPermission := controller.HasPermissionByPath(user.ID, c.FullPath())
			if !hasPermission {
				// 没有权限，返回403
				response.Forbidden(c)
				c.Abort()
				return
			}
		}

		// 用户存在 将user的信息写入上下文
		c.Set("user", user)

		c.Next()
	}
}

// ExposeHeaderMiddleware 前端暴露header设置
func ExposeHeaderMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Authorization")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

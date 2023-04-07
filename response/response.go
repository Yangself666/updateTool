package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Html(c *gin.Context, httpStatus int, template string, data interface{}) {
	c.HTML(httpStatus, template, data)
}

func Response(c *gin.Context, httpStatus int, code int, data interface{}, msg string) {
	c.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}

func ResponseFile(c *gin.Context, filePath string, fileName string) {
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Type", "application/octet-stream")
	c.File(filePath)
}

func Success(c *gin.Context, data interface{}, msg string) {
	Response(c, http.StatusOK, 200, data, msg)
}

func Fail(c *gin.Context, data interface{}, msg string) {
	Response(c, http.StatusOK, 500, data, msg)
}

func Unauthorized(c *gin.Context) {
	Response(c, http.StatusOK, http.StatusUnauthorized, nil, "凭证已过期")
}

func Forbidden(c *gin.Context) {
	Response(c, http.StatusOK, http.StatusForbidden, nil, "权限不足")
}

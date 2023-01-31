package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Html(c *gin.Context, httpStatus int, template string, data gin.H) {
	c.HTML(httpStatus, template, data)
}

func Response(c *gin.Context, httpStatus int, code int, data interface{}, msg string) {
	c.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}

func Success(c *gin.Context, data interface{}, msg string) {
	Response(c, http.StatusOK, 200, data, msg)
}

func Fail(c *gin.Context, data interface{}, msg string) {
	Response(c, http.StatusOK, 500, data, msg)
}

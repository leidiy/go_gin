package response

import (
	"github.com/gin-gonic/gin"
	"go_gin/pkg/e"
	"net/http"
)

func ReturnJSON(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func ReturnError(c *gin.Context, code int, msg string) {
	ReturnJSON(c, code, msg, nil)
}

func ReturnSuccess(c *gin.Context) {
	ReturnJSON(c, e.SUCCESS, "ok", nil)
}

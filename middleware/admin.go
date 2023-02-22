package middleware

import (
	"github.com/Yuzuki616/Aws-Panel/data"
	"github.com/gin-gonic/gin"
)

func AdminCheck(c *gin.Context) {
	username := c.GetString("email")
	if !data.IsAdmin(username) {
		c.JSON(403, gin.H{
			"code": 403,
			"msg":  "没有权限",
		})
		c.Abort()
	}
}

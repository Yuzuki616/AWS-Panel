package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SessionCheck(c *gin.Context) string {
	session := sessions.Default(c)
	loginUser := session.Get("loginuser")
	if loginUser == nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "用户未登录",
		})
		return ""
	}
	return loginUser.(string)
}

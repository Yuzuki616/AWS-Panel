package middleware

import (
	"github.com/Yuzuki616/Aws-Panel/cache"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func UserCheck(c *gin.Context) {
	s := sessions.Default(c)
	id := s.Get("loginSession")
	if id == nil {
		c.JSON(401, gin.H{
			"code": 401,
			"msg":  "用户未登录",
		})
		c.Abort()
	}
	username, _ := cache.Get(id.(string))
	if username == "" {
		c.JSON(401, gin.H{
			"code": 401,
			"msg":  "用户未登录",
		})
		c.Abort()
	}
	c.Set("username", username)
	return
}

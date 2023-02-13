package middleware

import (
	"github.com/Yuzuki616/Aws-Panel/cache"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/modern-go/reflect2"
)

func UserCheck(c *gin.Context) {
	s := sessions.Default(c)
	id := s.Get("loginSession")
	if reflect2.IsNil(id) {
		c.JSON(401, gin.H{
			"code": 401,
			"msg":  "用户未登录",
		})
		c.Abort()
	}
	username, ok := cache.Get(id.(string))
	if !ok {
		c.JSON(401, gin.H{
			"code": 401,
			"msg":  "用户未登录",
		})
		c.Abort()
	}
	c.Set("username", username)
	return
}

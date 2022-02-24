package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetSessionId(c *gin.Context) int {
	s := sessions.Default(c)
	id := s.Get("loginSession")
	if id == nil {
		return 0
	}
	return id.(int)
}

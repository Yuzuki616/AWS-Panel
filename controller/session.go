package controller

import (
	"github.com/Yuzuki616/Aws-Panel/cache"
	"github.com/Yuzuki616/Aws-Panel/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func addLoginSession(c *gin.Context, username string) error {
	sessionId := utils.GenRandomString(16)
	if _, e := cache.Get(sessionId); e {
		return addLoginSession(c, username)
	}
	cache.Set(sessionId, username, 0)
	s := sessions.Default(c)
	s.Set("loginSession", sessionId)
	saveErr := s.Save()
	if saveErr != nil {
		return saveErr
	}
	return nil
}

func delLoginSession(c *gin.Context, sessionId string) error {
	cache.Delete(sessionId)
	s := sessions.Default(c)
	s.Delete("loginSession")
	err := s.Save()
	if err != nil {
		return err
	}
	return nil
}

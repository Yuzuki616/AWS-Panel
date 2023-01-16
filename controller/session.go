package controller

import (
	"encoding/hex"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"math/rand"
)

func getSessionId(c *gin.Context) string {
	s := sessions.Default(c)
	id := s.Get("loginSession")
	if id == nil {
		return ""
	}
	return id.(string)
}

func genRandomString(size int) string {
	tmp := make([]byte, size)
	rand.Read(tmp)
	return hex.EncodeToString(tmp)
}
func addLoginSession(c *gin.Context, username string) error {
	sessionId := genRandomString(16)
	if _, e := cacheClient.Get(sessionId); e {
		return addLoginSession(c, username)
	}
	s := sessions.Default(c)
	s.Set("loginSession", sessionId)
	saveErr := s.Save()
	if saveErr != nil {
		return saveErr
	}
	return nil
}

func delLoginSession(c *gin.Context, sessionId string) error {
	cacheClient.Delete(sessionId)
	s := sessions.Default(c)
	s.Delete("loginSession")
	err := s.Save()
	if err != nil {
		return err
	}
	return nil
}

package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/yuzuki999/Aws-Panel/data"
	"github.com/yuzuki999/Aws-Panel/utils"
)

func LoginVerify(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "信息填写不完整",
		})
	}
	loginErr := data.LoginVerify(username, utils.Md5Encode(password))
	if loginErr == nil {
		session := sessions.Default(c)
		session.Set("loginuser", username)
		saveErr := session.Save()
		if saveErr != nil {
			log.Error("")
		}
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "登录成功",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  loginErr.Error(),
		})
	}
}

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "信息填写不完整",
		})
	}
	registerErr := data.Register(username, utils.Md5Encode(password))
	if registerErr == nil {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "注册成功",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  registerErr.Error(),
		})
	}
}

func ChangePassword(c *gin.Context) {
	username := SessionCheck(c)
	oldPassword := c.PostForm("oldPassword")
	newPassword := c.PostForm("newPassword")
	if oldPassword == "" || newPassword == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "信息填写不完整",
		})
	}
	if username == "" {
		return
	}
	changeErr := data.ChangePassword(username, utils.Md5Encode(oldPassword), utils.Md5Encode(newPassword))
	if changeErr == nil {
		session := sessions.Default(c)
		session.Clear()
		_ = session.Save()
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "修改成功",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  changeErr.Error(),
		})
	}
}

func GetUserInfo(c *gin.Context) {
	username := SessionCheck(c)
	if username == "" {
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"name": username,
	})
}

func Logout(c *gin.Context) {
	username := SessionCheck(c)
	if username == "" {
		return
	}
	session := sessions.Default(c)
	session.Clear()
	_ = session.Save()
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "退出成功",
	})
}

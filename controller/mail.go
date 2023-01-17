package controller

import (
	"github.com/Yuzuki616/Aws-Panel/cache"
	"github.com/Yuzuki616/Aws-Panel/conf"
	"github.com/Yuzuki616/Aws-Panel/mail"
	"github.com/Yuzuki616/Aws-Panel/utils"
	"github.com/gin-gonic/gin"
	"time"
)

func SendMailVerify(c *gin.Context) {
	email := c.PostForm("email")
	to := c.PostForm("to")
	if !conf.Config.EnableMailVerify {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "未开启邮箱验证",
		})
	}
	if _, e := cache.Get(email + "|codeLimit"); e {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "发送间隔太短，请稍后再试",
		})
	} else {
		code := utils.GenRandomString(6)
		cache.Set(email+"|code", code, time.Minute*5)
		cache.Set(email+"|codeLimit", nil, time.Second*60)
		sendErr := mail.SendMail(email, to, code)
		if sendErr == nil {
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "发送成功",
			})
		}
	}
}

func IsEnableEmailVerify(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  conf.Config.EnableMailVerify,
	})
}

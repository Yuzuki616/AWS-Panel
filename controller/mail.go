package controller

import (
	"github.com/Yuzuki616/Aws-Panel/data"
	"github.com/Yuzuki616/Aws-Panel/mail"
	"github.com/Yuzuki616/Aws-Panel/session"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
)

func SendMailVerify(c *gin.Context) {
	email := c.PostForm("email")
	to := c.PostForm("to")
	config := data.GetMailConfig()
	if config.EmailVerity == 1 {
		if session.GetMailLast(email) != "" {
			c.JSON(400, gin.H{
				"code": 400,
				"msg":  "发送间隔太短，请稍后再试",
			})
		} else {
			code := strconv.Itoa(100000 + rand.Intn(999999-100000+1))
			session.CreateMailCode(email, code)
			sendErr := mail.SendMail(config.Email, to, code)
			if sendErr == nil {
				c.JSON(200, gin.H{
					"code": 200,
					"msg":  "发送成功",
				})
			}
		}
	} else {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "未开启邮箱验证",
		})
	}
}

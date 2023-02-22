package controller

import (
	"github.com/Yuzuki616/Aws-Panel/cache"
	"github.com/Yuzuki616/Aws-Panel/conf"
	"github.com/Yuzuki616/Aws-Panel/data"
	"github.com/Yuzuki616/Aws-Panel/mail"
	"github.com/Yuzuki616/Aws-Panel/request"
	"github.com/Yuzuki616/Aws-Panel/utils"
	"github.com/gin-gonic/gin"
	"time"
)

func Login(c *gin.Context) {
	params := &request.Login{}
	if err := c.ShouldBind(params); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	err := data.VerifyUser(params.Email, params.Password)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	err = addLoginSession(c, params.Email)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"isAdmin": data.IsAdmin(params.Email),
		"msg":     "登录成功",
	})
}

func SendMailVerify(c *gin.Context) {
	p := request.SendMailVerify{}
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	if !conf.Config.EnableMailVerify {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "未开启邮箱验证",
		})
		return
	}
	if _, e := cache.Get(p.Email + "|codeLimit"); e {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "发送间隔太短，请稍后再试",
		})
		return
	}
	code := utils.GenRandomString(6)
	cache.Set(p.Email+"|code", code, time.Minute*5)
	cache.Set(p.Email+"|codeLimit", nil, time.Second*60)
	sendErr := mail.SendMail(p.Email, code)
	if sendErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "发送失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "发送成功",
	})
}

func Register(c *gin.Context) {
	params := &request.Register{}
	if err := c.ShouldBind(params); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	if conf.Config.EnableMailVerify {
		code := c.PostForm("code")
		if savedCode, e := cache.Get(code + "|code"); !e || savedCode.(string) != code {
			c.JSON(400, gin.H{
				"code": 400,
				"msg":  "验证码不正确或已失效",
			})
			return
		}
	}
	registerErr := data.CreateUser(params.Email, params.Email, params.Password, 0)
	if registerErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  registerErr.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "注册成功",
	})
}

func ChangeEmail(c *gin.Context) {
	params := &request.ChangeUsername{}
	if err := c.ShouldBind(params); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	err := data.ChangeEmail(params.OldEmail, params.NewEmail, params.Password)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	Logout(c)
	if c.IsAborted() {
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "修改成功",
	})
}

func ChangePassword(c *gin.Context) {
	username := c.GetString("email")
	params := &request.ChangePassword{}
	if err := c.ShouldBind(params); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	changeErr := data.ChangeUserPassword(username, params.OldPassword, params.NewPassword)
	if changeErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  changeErr.Error(),
		})
	}
	Logout(c)
	if c.IsAborted() {
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "修改成功",
	})
}

func GetUserInfo(c *gin.Context) {
	email := c.GetString("email")
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": email,
	})
}

func Logout(c *gin.Context) {
	err := delLoginSession(c)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "退出成功",
	})
}

func IsAdmin(c *gin.Context) {
	email := c.GetString("email")
	if data.IsAdmin(email) {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  true,
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  false,
		})
	}
}

func DeleteUser(c *gin.Context) {
	params := &request.DeleteUser{}
	if err := c.ShouldBind(params); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	err := data.DeleteUser(params.Email)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}

func BanUser(c *gin.Context) {
	params := &request.BanUser{}
	if err := c.ShouldBind(params); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	err := data.BanUser(params.Email)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "封禁成功",
	})
}

func UnBanUser(c *gin.Context) {
	params := &request.BanUser{}
	if err := c.ShouldBind(params); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	err := data.UnBanUser(params.Email)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "解封成功",
	})
}

func GetUserList(c *gin.Context) {
	list, listErr := data.GetUserList()
	if listErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  listErr.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": list,
	})
}

package controller

import (
	"errors"
	"github.com/Yuzuki616/Aws-Panel/cache"
	"github.com/Yuzuki616/Aws-Panel/conf"
	"github.com/Yuzuki616/Aws-Panel/data"
	"github.com/Yuzuki616/Aws-Panel/mail"
	"github.com/Yuzuki616/Aws-Panel/request"
	"github.com/Yuzuki616/Aws-Panel/utils"
	"github.com/gin-gonic/gin"
	"time"
)

func Login(c *gin.Context) error {
	params := &request.Login{}
	if err := c.ShouldBind(params); err != nil {
		return err
	}
	err := data.VerifyUser(params.Email, params.Password)
	if err != nil {
		return err
	}
	err = addLoginSession(c, params.Email)
	if err != nil {
		return err
	}
	c.JSON(200, gin.H{
		"code":    200,
		"isAdmin": data.IsAdmin(params.Email),
		"msg":     "登录成功",
	})
	return nil
}

func SendMailVerify(c *gin.Context) error {
	p := request.SendMailVerify{}
	if err := c.ShouldBind(&p); err != nil {
		return err
	}
	if !conf.Config.EnableMailVerify {
		return errors.New("未开启邮箱验证")
	}
	if _, e := cache.Get(p.Email + "|codeLimit"); e {
		return errors.New("发送间隔太短，请稍后再试")
	}
	code := utils.GenRandomString(6)
	cache.Set(p.Email+"|code", code, time.Minute*5)
	cache.Set(p.Email+"|codeLimit", nil, time.Second*60)
	sendErr := mail.SendMail(p.Email, code)
	if sendErr != nil {
		return sendErr
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "发送成功",
	})
	return nil
}

func Register(c *gin.Context) error {
	params := &request.Register{}
	if err := c.ShouldBind(params); err != nil {
		return err
	}
	if conf.Config.EnableMailVerify {
		code := c.PostForm("code")
		if savedCode, e := cache.Get(code + "|code"); !e || savedCode.(string) != code {
			return errors.New("验证码不正确或已失效")
		}
	}
	registerErr := data.CreateUser(params.Email, params.Email, params.Password, 0)
	if registerErr != nil {
		return registerErr
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "注册成功",
	})
	return nil
}

func ChangeEmail(c *gin.Context) error {
	params := &request.ChangeUsername{}
	if err := c.ShouldBind(params); err != nil {
		return err
	}
	err := data.ChangeEmail(params.OldEmail, params.NewEmail, params.Password)
	if err != nil {
		return err
	}
	Logout(c)
	if c.IsAborted() {
		return nil
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "修改成功",
	})
	return nil
}

func ChangePassword(c *gin.Context) error {
	username := c.GetString("email")
	params := &request.ChangePassword{}
	if err := c.ShouldBind(params); err != nil {
		return err
	}
	changeErr := data.ChangeUserPassword(username, params.OldPassword, params.NewPassword)
	if changeErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  changeErr.Error(),
		})
	}
	err := Logout(c)
	if err != nil {
		return err
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "修改成功",
	})
	return nil
}

func GetUserInfo(c *gin.Context) {
	email := c.GetString("email")
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": email,
	})
}

func Logout(c *gin.Context) error {
	err := delLoginSession(c)
	if err != nil {
		return err
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "退出成功",
	})
	return nil
}

func IsAdmin(c *gin.Context) {
	email := c.GetString("email")
	if data.IsAdmin(email) {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  true,
		})
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  false,
	})
}

func DeleteUser(c *gin.Context) error {
	params := &request.DeleteUser{}
	if err := c.ShouldBind(params); err != nil {
		return err
	}
	err := data.DeleteUser(params.Email)
	if err != nil {
		return err
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
	return nil
}

func BanUser(c *gin.Context) error {
	params := &request.BanUser{}
	if err := c.ShouldBind(params); err != nil {
		return err
	}
	err := data.BanUser(params.Email)
	if err != nil {
		return err
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "封禁成功",
	})
	return nil
}

func UnBanUser(c *gin.Context) error {
	params := &request.BanUser{}
	if err := c.ShouldBind(params); err != nil {
		return err
	}
	err := data.UnBanUser(params.Email)
	if err != nil {
		return err
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "解封成功",
	})
	return nil
}

func GetUserList(c *gin.Context) error {
	list, listErr := data.GetUserList()
	if listErr != nil {
		return listErr
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": list,
	})
	return nil
}

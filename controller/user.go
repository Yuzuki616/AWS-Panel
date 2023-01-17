package controller

import (
	"github.com/Yuzuki616/Aws-Panel/cache"
	"github.com/Yuzuki616/Aws-Panel/conf"
	"github.com/Yuzuki616/Aws-Panel/data"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoginVerify(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "信息填写不完整",
		})
	}
	err := data.VerifyUser(username, password)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	err = addLoginSession(c, username)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"isAdmin": data.IsAdmin(username),
		"msg":     "登录成功",
	})
}

func Register(c *gin.Context) {
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "信息填写不完整",
		})
	}
	if conf.Config.EnableMailVerify {
		code := c.PostForm("code")
		if savedCode, e := cache.Get(email + "|code"); !e || savedCode.(string) != code {
			c.JSON(400, gin.H{
				"code": 400,
				"msg":  "验证码不正确或已失效",
			})
			return
		}
	}
	registerErr := data.CreateUser(username, email, password, 0)
	if registerErr == nil {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "注册成功",
		})
	} else {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  registerErr.Error(),
		})
	}
}

func ChangeUsername(c *gin.Context) {
	username, _ := c.Get("username")
	oldName := c.PostForm("oldUsername")
	newName := c.PostForm("newUsername")
	password := c.PostForm("password")
	if username == "" {
		return
	}
	err := data.ChangeUsername(oldName, newName, password)
	Logout(c)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "修改成功",
		})
	}
}

func ChangePassword(c *gin.Context) {
	username, _ := c.Get("username")
	oldPassword := c.PostForm("oldPassword")
	newPassword := c.PostForm("newPassword")
	if oldPassword == "" || newPassword == "" {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "信息填写不完整",
		})
	}
	if username == "" {
		return
	}
	changeErr := data.ChangeUserPassword(username.(string), oldPassword, newPassword)
	if changeErr == nil {
		s := sessions.Default(c)
		s.Clear()
		_ = s.Save()
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "修改成功",
		})
	} else {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  changeErr.Error(),
		})
	}
}

func GetUserInfo(c *gin.Context) {
	username, _ := c.Get("username")
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
	s := sessions.Default(c)
	id := s.Get("loginSession")
	err := delLoginSession(c, id.(string))
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "退出成功",
	})
}

func IsAdmin(c *gin.Context) {
	username, _ := c.Get("username")
	if username == "" {
		return
	}
	if data.IsAdmin(username.(string)) {
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
	user := c.PostForm("username")
	username, _ := c.Get("username")
	if username == "" {
		return
	}
	if data.IsAdmin(username.(string)) {
		err := data.DeleteUser(user)
		if err != nil {
			c.JSON(400, gin.H{
				"code": 400,
				"msg":  err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "删除成功",
			})
		}
	} else {
		c.JSON(403, gin.H{
			"code": 403,
			"msg":  "没有权限",
		})
	}
}

func BanUser(c *gin.Context) {
	user := c.PostForm("username")
	err := data.BanUser(user)
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
	user := c.PostForm("username")
	err := data.UnBanUser(user)
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

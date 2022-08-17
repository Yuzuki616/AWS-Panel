package controller

import (
	"github.com/Yuzuki616/Aws-Panel/data"
	"github.com/Yuzuki616/Aws-Panel/session"
	"github.com/Yuzuki616/Aws-Panel/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetLoginUser(c *gin.Context) string {
	username := session.GetSession(GetSessionId(c))
	if username == "" {
		c.JSON(401, gin.H{
			"code": 401,
			"msg":  "用户未登录",
		})
		return ""
	}
	return username
}

func LoginVerify(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "信息填写不完整",
		})
	}
	loginErr := data.LoginVerify(username, utils.Md5Encode(password))
	if loginErr == nil {
		s := sessions.Default(c)
		s.Set("loginSession", session.CreateSession(username, 0))
		saveErr := s.Save()
		if saveErr != nil {
			log.Error("")
		}
		c.JSON(200, gin.H{
			"code":    200,
			"isAdmin": data.IsAdmin(username),
			"msg":     "登录成功",
		})
	} else {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  loginErr.Error(),
		})
	}
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
	if data.IsEmailVerity() {
		code := c.PostForm("code")
		if code != session.GetMailCode(email) {
			c.JSON(400, gin.H{
				"code": 400,
				"msg":  "验证码不正确或已失效",
			})
			return
		}
	}
	registerErr := data.Register(username, email, utils.Md5Encode(password))
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
	username := GetLoginUser(c)
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
	username := GetLoginUser(c)
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
	changeErr := data.ChangePassword(username, utils.Md5Encode(oldPassword), utils.Md5Encode(newPassword))
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
	username := GetLoginUser(c)
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
	id := GetSessionId(c)
	session.DeleteSession(id)
	s := sessions.Default(c)
	s.Clear()
	_ = s.Save()
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "退出成功",
	})
}

func IsAdmin(c *gin.Context) {
	username := GetLoginUser(c)
	if username == "" {
		return
	}
	if data.IsAdmin(username) {
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
	username := GetLoginUser(c)
	if username == "" {
		return
	}
	if data.IsAdmin(username) {
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
	username := GetLoginUser(c)
	if username == "" {
		return
	}
	if data.IsAdmin(username) {
		err := data.BanUser(user)
		if err != nil {
			c.JSON(400, gin.H{
				"code": 400,
				"msg":  err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "封禁成功",
			})
		}
	} else {
		c.JSON(403, gin.H{
			"code": 403,
			"msg":  "没有权限",
		})
	}
}

func UnBanUser(c *gin.Context) {
	user := c.PostForm("username")
	username := GetLoginUser(c)
	if username == "" {
		return
	}
	if data.IsAdmin(username) {
		err := data.UnBanUser(user)
		if err != nil {
			c.JSON(400, gin.H{
				"code": 400,
				"msg":  err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "解封成功",
			})
		}
	} else {
		c.JSON(403, gin.H{
			"code": 403,
			"msg":  "没有权限",
		})
	}
}

func GetUserList(c *gin.Context) {
	username := GetLoginUser(c)
	if username == "" {
		return
	}
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

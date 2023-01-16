package controller

import (
	"github.com/Yuzuki616/Aws-Panel/data"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func getLoginUser(c *gin.Context) string {
	username, _ := cacheClient.Get(getSessionId(c))
	if username == "" {
		c.JSON(401, gin.H{
			"code": 401,
			"msg":  "用户未登录",
		})
		return ""
	}
	return username.(string)
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
	if data.IsEmailVerity() {
		code := c.PostForm("code")
		if savedCode, e := cacheClient.Get(email + "|code"); !e || savedCode.(string) != code {
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
	username := getLoginUser(c)
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
	username := getLoginUser(c)
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
	changeErr := data.ChangeUserPassword(username, oldPassword, newPassword)
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
	username := getLoginUser(c)
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
	id := getSessionId(c)
	err := delLoginSession(c, id)
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
	username := getLoginUser(c)
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
	username := getLoginUser(c)
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
	username := getLoginUser(c)
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
	username := getLoginUser(c)
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
	username := getLoginUser(c)
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

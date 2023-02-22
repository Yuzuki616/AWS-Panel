package controller

import (
	"github.com/Yuzuki616/Aws-Panel/data"
	"github.com/Yuzuki616/Aws-Panel/request"
	"github.com/gin-gonic/gin"
)

func AddSecret(c *gin.Context) {
	email := c.GetString("email")
	params := &request.AddSecret{}
	if err := c.ShouldBind(params); err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	addErr := data.AddSecret(email, params.Name, params.Id, params.Secret)
	if addErr != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  addErr.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "添加成功",
	})
}

func ListSecret(c *gin.Context) {
	email := c.GetString("email")
	secret, listErr := data.ListSecret(email)
	if listErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  listErr.Error(),
		})
		return
	}
	var tmp []map[string]string
	for _, v := range secret {
		tmp = append(tmp, map[string]string{
			"name":   v.Name,
			"id":     v.SecretId,
			"secret": v.Secret,
		})
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": tmp,
	})
}

func GetSecret(c *gin.Context) {
	email := c.GetString("email")
	params := &request.GetSecret{}
	secret, getErr := data.GetSecret(email, params.Name)
	if getErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  getErr.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": secret,
	})
}

func DelSecret(c *gin.Context) {
	email := c.GetString("email")
	params := &request.DelSecret{}
	if err := c.ShouldBind(params); err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	delErr := data.DelSecret(email, params.Name)
	if delErr != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  delErr.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}

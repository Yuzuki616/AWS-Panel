package controller

import (
	"github.com/Yuzuki616/Aws-Panel/data"
	"github.com/Yuzuki616/Aws-Panel/request"
	"github.com/gin-gonic/gin"
)

func AddSecret(c *gin.Context) error {
	email := c.GetString("email")
	params := &request.AddSecret{}
	if err := c.ShouldBind(params); err != nil {
		return err
	}
	addErr := data.AddSecret(email, params.Name, params.Id, params.Secret)
	if addErr != nil {
		return addErr
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "添加成功",
	})
	return nil
}

func ListSecret(c *gin.Context) error {
	email := c.GetString("email")
	secret, listErr := data.ListSecret(email)
	if listErr != nil {
		return listErr
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
	return nil
}

func GetSecret(c *gin.Context) error {
	email := c.GetString("email")
	params := &request.GetSecret{}
	secret, getErr := data.GetSecret(email, params.Name)
	if getErr != nil {
		return getErr
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": secret,
	})
	return nil
}

func DelSecret(c *gin.Context) error {
	email := c.GetString("email")
	params := &request.DelSecret{}
	if err := c.ShouldBind(params); err != nil {
		return err
	}
	delErr := data.DelSecret(email, params.Name)
	if delErr != nil {
		return delErr
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
	return nil
}

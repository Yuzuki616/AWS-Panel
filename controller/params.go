package controller

import (
	"github.com/gin-gonic/gin"
)

func GetAndCheckParams(c *gin.Context, params ...string) map[string]string {
	result := make(map[string]string, len(params))
	for _, v := range params {
		param := c.PostForm(v)
		if param == "" {
			return nil
		}
		result[v] = param
	}
	c.JSON(400, gin.H{
		"code": 400,
		"msg":  "信息填写不完整",
	})
	return result
}

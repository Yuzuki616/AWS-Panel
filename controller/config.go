package controller

import (
	"github.com/Yuzuki616/Aws-Panel/conf"
	"github.com/gin-gonic/gin"
)

func IsEnableEmailVerify(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
		"data": conf.Config.EnableMailVerify,
	})
}

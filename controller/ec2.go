package controller

import (
	"github.com/Yuzuki616/Aws-Panel/aws"
	"github.com/Yuzuki616/Aws-Panel/data"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CreateEc2(c *gin.Context) {
	username := GetLoginUser(c)
	if username == "" {
		return
	}
	userdata := c.PostForm("userdata")
	params := GetAndCheckParams(c, "secretName", "region", "ami", "ec2Type", "ec2Name", "disk")
	if len(params) == 0 {
		return
	}
	disk, _ := strconv.ParseInt(params["disk"], 10, 64)
	secret, _ := data.GetSecret(username, params["secretName"])
	client, newErr := aws.New(params["region"], secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	amiTmp, amiErr := client.GetAmiId(params["ami"])
	if amiErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  amiErr.Error(),
		})
		return
	}
	if amiTmp == "" {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "Get ami ID error: Not found ami",
		})
		return
	}
	creRt, creErr := client.CreateEc2(amiTmp, params["ec2Type"], params["ec2Name"], userdata, disk)
	if creErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  creErr.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "创建成功",
		"data": creRt.Key,
	})
}

func ListEc2(c *gin.Context) {
	username := GetLoginUser(c)
	if username == "" {
		return
	}
	params := GetAndCheckParams(c, "secretName", "region")
	if len(params) == 0 {
		return
	}
	secret, _ := data.GetSecret(username, params["secretName"])
	client, newErr := aws.New(params["region"], secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	ec2Info, listErr := client.ListEc2()
	if listErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  listErr.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"data": ec2Info,
	})
}

func GetEc2Info(c *gin.Context) {
	username := GetLoginUser(c)
	if username == "" {
		return
	}
	params := GetAndCheckParams(c, "secretName", "region", "ec2Id")
	if len(params) == 0 {
		return
	}
	secret, _ := data.GetSecret(username, params["secretName"])
	client, newErr := aws.New(params["region"], secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	ec2Info, getErr := client.GetEc2Info(params["ec2Id"])
	if getErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  getErr.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": ec2Info,
	})
}

func ChangeEc2Ip(c *gin.Context) {
	username := GetLoginUser(c)
	if username == "" {
		return
	}
	params := GetAndCheckParams(c, "secretName", "region", "ec2Id")
	if len(params) == 0 {
		return
	}
	secret, _ := data.GetSecret(username, params["secretName"])
	client, newErr := aws.New(params["region"], secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	newIp, changeErr := client.ChangeEc2Ip(params["ec2Id"])
	if changeErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  changeErr.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "更换IP成功",
		"data": newIp,
	})
}

func StopEc2(c *gin.Context) {
	username := GetLoginUser(c)
	if username == "" {
		return
	}
	params := GetAndCheckParams(c, "secretName", "region", "ec2Id")
	if len(params) == 0 {
		return
	}
	secret, _ := data.GetSecret(username, params["secretName"])
	client, newErr := aws.New(params["region"], secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	stopErr := client.StopEc2(params["ec2Id"])
	if stopErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  stopErr.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "启动失败",
	})
}

func StartEc2(c *gin.Context) {
	username := GetLoginUser(c)
	if username == "" {
		return
	}
	params := GetAndCheckParams(c, "secretName", "region", "ec2Id")
	if len(params) == 0 {
		return
	}
	secret, _ := data.GetSecret(username, params["secretName"])
	client, newErr := aws.New(params["region"], secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
	}
	startErr := client.StartEc2(params["ec2Id"])
	if startErr == nil {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "启动成功",
		})
	} else {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  startErr.Error(),
		})
	}
}

func RebootEc2(c *gin.Context) {
	username := GetLoginUser(c)
	if username == "" {
		return
	}
	params := GetAndCheckParams(c, "secretName", "region", "ec2Id")
	if len(params) == 0 {
		return
	}
	secret, _ := data.GetSecret(username, params["secretName"])
	client, newErr := aws.New(params["region"], secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	rebootErr := client.RebootEc2(params["ec2Id"])
	if rebootErr == nil {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "重启成功",
		})
	} else {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  rebootErr.Error(),
		})
	}
}

func DeleteEc2(c *gin.Context) {
	username := GetLoginUser(c)
	if username == "" {
		return
	}
	params := GetAndCheckParams(c, "secretName", "region", "ec2Id")
	if len(params) == 0 {
		return
	}
	secret, _ := data.GetSecret(username, params["secretName"])
	client, newErr := aws.New(params["region"], secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	delErr := client.DeleteEc2(params["ec2Id"])
	if delErr == nil {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "删除成功",
		})
	} else {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  delErr.Error(),
		})
	}
}

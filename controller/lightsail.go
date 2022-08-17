package controller

import (
	"github.com/Yuzuki616/Aws-Panel/aws"
	"github.com/Yuzuki616/Aws-Panel/data"
	"github.com/gin-gonic/gin"
)

func GetRegions(c *gin.Context) {
	username := GetLoginUser(c)
	if username == "" {
		return
	}
	secretName := c.PostForm("secretName")
	if secretName == "" {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "信息填写不完整",
		})
		return
	}
	secret, _ := data.GetSecret(username, secretName)
	client, newErr := aws.New("ap-northeast-1", secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	regions, GetRegionsErr := client.GetRegions()
	if GetRegionsErr == nil {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "查询成功",
			"data": regions,
		})
	} else {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  GetRegionsErr.Error(),
		})
	}
}

func CreateLightsail(c *gin.Context) {
	username := GetLoginUser(c)
	if username == "" {
		return
	}
	params := GetAndCheckParams(c, "name", "zone", "availabilityZone", "blueprintId", "bundleId")
	if len(params) == 0 {
		return
	}
	secret, _ := data.GetSecret(username, params["secretName"])
	client, newErr := aws.New(params["zone"], secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	createRt, createErr := client.CreateLs(params["name"], params["availabilityZone"], params["blueprintId"], params["bundleId"])
	if createErr == nil {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "创建成功",
			"data": createRt.Key,
		})
	} else {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  createErr.Error(),
		})
	}
}

func OpenLightsailPorts(c *gin.Context) {
	username := GetLoginUser(c)
	if username == "" {
		return
	}
	params := GetAndCheckParams(c, "zone", "secretName", "name")
	secret, _ := data.GetSecret(username, params["secretName"])
	client, newErr := aws.New(params["zone"], secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	openErr := client.OpenLsPorts(params["name"])
	if openErr == nil {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "开放成功",
		})
	} else {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  openErr.Error(),
		})
	}
}

func ListLightsail(c *gin.Context) {
	username := GetLoginUser(c)
	if username == "" {
		return
	}
	params := GetAndCheckParams(c, "zone", "secretName")
	if len(params) == 0 {
		return
	}
	secret, _ := data.GetSecret(username, params["secretName"])
	client, newErr := aws.New(params["zone"], secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	listRt, listErr := client.ListLs()
	if listErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  listErr.Error(),
		})
		return
	}
	var instances []*aws.LsInfo
	if listRt == nil {
		instances = append(instances, &aws.LsInfo{})
	} else {

	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": instances,
	})
}

func GetLightsailInfo(c *gin.Context) {
	username := GetLoginUser(c)
	if username == "" {
		return
	}
	params := GetAndCheckParams(c, "zone", "secretName", "name")
	name := c.PostForm("name")
	secret, _ := data.GetSecret(username, params["secretName"])
	client, newErr := aws.New(params["zone"], secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	info, GetInfoErr := client.GetLsInfo(name)
	if GetInfoErr == nil {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "删除成功",
			"data": info,
		})
	} else {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  GetInfoErr.Error(),
		})
	}
}

func StartLightsail(c *gin.Context) {
	username := GetLoginUser(c)
	if username == "" {
		return
	}
	params := GetAndCheckParams(c, "zone", "secretName", "name")
	if len(params) == 0 {
		return
	}
	secret, _ := data.GetSecret(username, params["secretName"])
	client, newErr := aws.New(params["zone"], secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	startErr := client.StartLs(params["name"])
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

func StopLightsail(c *gin.Context) {
	username := GetLoginUser(c)
	if username == "" {
		return
	}
	params := GetAndCheckParams(c, "zone", "secretName", "name")
	if len(params) == 0 {
		return
	}
	secret, _ := data.GetSecret(username, params["secretName"])
	client, newErr := aws.New(params["zone"], secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	stopErr := client.StopLs(params["name"])
	if stopErr == nil {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "停止成功",
		})
	} else {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  stopErr.Error(),
		})
	}
}

func RebootLightsail(c *gin.Context) {
	username := GetLoginUser(c)
	if username == "" {
		return
	}
	params := GetAndCheckParams(c, "zone", "secretName", "name")
	if len(params) == 0 {
		return
	}
	secret, _ := data.GetSecret(username, params["secretName"])
	client, newErr := aws.New(params["zone"], secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	stopErr := client.RebootLs(params["name"])
	if stopErr == nil {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "停止成功",
		})
	} else {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  stopErr.Error(),
		})
	}
}

func ChangeLightsailIp(c *gin.Context) {
	username := GetLoginUser(c)
	if username == "" {
		return
	}
	params := GetAndCheckParams(c, "zone", "secretName", "name")
	if len(params) == 0 {
		return
	}
	secret, _ := data.GetSecret(username, params["secretName"])
	client, newErr := aws.New(params["zone"], secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	changeIpErr := client.ChangeLsIp(params["name"])
	if changeIpErr == nil {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "更换IP成功",
		})
	} else {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  changeIpErr.Error(),
		})
	}
}

func DeleteLightsail(c *gin.Context) {
	username := GetLoginUser(c)
	if username == "" {
		return
	}
	params := GetAndCheckParams(c, "secretName", "zone", "name", "resourceName")
	if len(params) == 0 {
		return
	}
	secret, _ := data.GetSecret(username, params["secretName"])
	client, newErr := aws.New(params["zone"], secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	startErr := client.DeleteLs(params["name"], params["resourceName"])
	if startErr == nil {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "删除成功",
		})
	} else {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  startErr.Error(),
		})
	}
}

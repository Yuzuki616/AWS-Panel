package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yuzuki999/Aws-Panel/aws"
	"github.com/yuzuki999/Aws-Panel/data"
)

func CreateLightsail(c *gin.Context) {
	username := GetLoginUser(c)
	if username == "" {
		return
	}
	secretName := c.PostForm("secretName")
	name := c.PostForm("name")
	zone := c.PostForm("zone")
	blueprintId := c.PostForm("blueprintId")
	bundleId := c.PostForm("bundleId")
	secret, _ := data.GetSecret(username, secretName)
	client, newErr := aws.New("ap-northeast-1", secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	createRt, createErr := client.CreateLs(name, zone, blueprintId, bundleId)
	if createErr == nil {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "创建成功",
		})
	} else {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  createErr.Error(),
			"data": *createRt.Key,
		})
	}
}

func ListLightsail(c *gin.Context) {
	username := GetLoginUser(c)
	if username == "" {
		return
	}
	secretName := c.PostForm("secretName")
	secret, _ := data.GetSecret(username, secretName)
	client, newErr := aws.New("ap-northeast-1", secret.SecretId, secret.Secret, "")
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
		for _, v := range listRt {
			instances = append(instances, &aws.LsInfo{
				Name:   v.Name,
				Type:   v.BundleId,
				Ip:     v.PublicIpAddress,
				Status: v.State.Name,
			})
		}
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
	secretName := c.PostForm("secretName")
	name := c.PostForm("name")
	secret, _ := data.GetSecret(username, secretName)
	client, newErr := aws.New("ap-northeast-1", secret.SecretId, secret.Secret, "")
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
	secretName := c.PostForm("secretName")
	name := c.PostForm("name")
	secret, _ := data.GetSecret(username, secretName)
	client, newErr := aws.New("ap-northeast-1", secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	startErr := client.StartLs(name)
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
	secretName := c.PostForm("secretName")
	name := c.PostForm("name")
	secret, _ := data.GetSecret(username, secretName)
	client, newErr := aws.New("ap-northeast-1", secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	stopErr := client.StopLs(name)
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
	secretName := c.PostForm("secretName")
	name := c.PostForm("name")
	secret, _ := data.GetSecret(username, secretName)
	client, newErr := aws.New("ap-northeast-1", secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	stopErr := client.RebootLs(name)
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
	secretName := c.PostForm("secretName")
	name := c.PostForm("name")
	secret, _ := data.GetSecret(username, secretName)
	client, newErr := aws.New("ap-northeast-1", secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	changeIpErr := client.ChangeLsIp(name)
	if changeIpErr == nil {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "更换IP成功",
		})
	} else {
		c.JSON(200, gin.H{
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
	secretName := c.PostForm("secretName")
	name := c.PostForm("name")
	secret, _ := data.GetSecret(username, secretName)
	client, newErr := aws.New("ap-northeast-1", secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	startErr := client.DeleteLs(name)
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

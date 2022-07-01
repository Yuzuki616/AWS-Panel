package controller

import (
	"github.com/aws/aws-sdk-go/service/lightsail"
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
	quantity := c.PostForm("quantity")
	awsscreenneed := c.PostForm("Awssecretkey")
	secret, _ := data.GetSecret(username, secretName)
	client, newErr := aws.New(zone, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	createRt, createErr := client.CreateLs(name, zone+"a", blueprintId, bundleId, quantity, awsscreenneed)
	if createErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  createErr.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "创建成功",
		"data": *createRt.Key,
	})

}

func ListLightsail(c *gin.Context) {
	username := GetLoginUser(c)

	if username == "" {
		return
	}
	secretName := c.PostForm("secretName")
	zone := c.PostForm("zone")
	if secretName == "" || zone == "" {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "信息填写不完整",
		})
	}
	secret, _ := data.GetSecret(username, secretName)
	client, newErr := aws.New(zone, secret.SecretId, secret.Secret, "")
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

	data.SortInstance(listRt, func(p, q *lightsail.Instance) bool {
		return *q.Name < *p.Name
	})

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
	zone := c.PostForm("zone")
	client, newErr := aws.New(zone, secret.SecretId, secret.Secret, "")
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
	zone := c.PostForm("zone")
	client, newErr := aws.New(zone, secret.SecretId, secret.Secret, "")
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
	zone := c.PostForm("zone")
	client, newErr := aws.New(zone, secret.SecretId, secret.Secret, "")
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
	zone := c.PostForm("zone")
	client, newErr := aws.New(zone, secret.SecretId, secret.Secret, "")
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
	zone := c.PostForm("zone")
	client, newErr := aws.New(zone, secret.SecretId, secret.Secret, "")
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
	zone := c.PostForm("zone")
	client, newErr := aws.New(zone, secret.SecretId, secret.Secret, "")
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

func OffFirewall(c *gin.Context) {
	username := GetLoginUser(c)
	if username == "" {
		return
	}
	secretName := c.PostForm("secretName")
	name := c.PostForm("name")
	zone := c.PostForm("zone")

	secret, _ := data.GetSecret(username, secretName)
	client, newErr := aws.New(zone, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}

	_, createErr := client.OffLightFirewall(name)
	if createErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  createErr.Error(),
		})

	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "关闭成功",
		})
	}

}

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yuzuki999/Aws-Panel/aws"
	"github.com/yuzuki999/Aws-Panel/data"
)

func GetRegions(c *gin.Context) {
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
	secretName := c.PostForm("secretName")
	name := c.PostForm("name")
	zone := c.PostForm("zone")
	availabilityZone := c.PostForm("availabilityZone")
	blueprintId := c.PostForm("blueprintId")
	bundleId := c.PostForm("bundleId")
	secret, _ := data.GetSecret(username, secretName)
	client, newErr := aws.New(zone, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	createRt, createErr := client.CreateLs(name, availabilityZone, blueprintId, bundleId)
	if createErr == nil {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "创建成功",
			"data": *createRt.Key,
		})
	} else {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  createErr.Error(),
		})
	}
}

func OpenLightsailPorts(c *gin.Context) {
	zone := c.PostForm("zone")
	username := GetLoginUser(c)
	if username == "" {
		return
	}
	secretName := c.PostForm("secretName")
	secret, _ := data.GetSecret(username, secretName)
	name := c.PostForm("name")
	client, newErr := aws.New(zone, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	openErr := client.OpenLsPorts(name)
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
	zone := c.PostForm("zone")
	username := GetLoginUser(c)
	if username == "" {
		return
	}
	secretName := c.PostForm("secretName")
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
	var instances []*aws.LsInfo
	if listRt == nil {
		instances = append(instances, &aws.LsInfo{})
	} else {
		var tag *string
		str := ""
		tag = &str
		for _, v := range listRt {
			if len(v.Tags) > 0 {
				tag = v.Tags[0].Value
			}
			instances = append(instances, &aws.LsInfo{
				Name:       v.Name,
				Type:       v.BundleId,
				Ip:         v.PublicIpAddress,
				SourceName: tag,
				Status:     v.State.Name,
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
	zone := c.PostForm("zone")
	username := GetLoginUser(c)
	if username == "" {
		return
	}
	secretName := c.PostForm("secretName")
	name := c.PostForm("name")
	secret, _ := data.GetSecret(username, secretName)
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
	zone := c.PostForm("zone")
	username := GetLoginUser(c)
	if username == "" {
		return
	}
	secretName := c.PostForm("secretName")
	name := c.PostForm("name")
	secret, _ := data.GetSecret(username, secretName)
	client, newErr := aws.New(zone, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
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
	zone := c.PostForm("zone")
	username := GetLoginUser(c)
	if username == "" {
		return
	}
	secretName := c.PostForm("secretName")
	name := c.PostForm("name")
	secret, _ := data.GetSecret(username, secretName)
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
	zone := c.PostForm("zone")
	username := GetLoginUser(c)
	if username == "" {
		return
	}
	secretName := c.PostForm("secretName")
	name := c.PostForm("name")
	secret, _ := data.GetSecret(username, secretName)
	client, newErr := aws.New(zone, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
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
	zone := c.PostForm("zone")
	username := GetLoginUser(c)
	if username == "" {
		return
	}
	secretName := c.PostForm("secretName")
	name := c.PostForm("name")
	secret, _ := data.GetSecret(username, secretName)
	client, newErr := aws.New(zone, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	changeIpErr := client.ChangeLsIp(name, zone)
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
	zone := c.PostForm("zone")
	username := GetLoginUser(c)
	if username == "" {
		return
	}
	secretName := c.PostForm("secretName")
	name := c.PostForm("name")
	IpName := c.PostForm("SourceName")
	secret, _ := data.GetSecret(username, secretName)
	client, newErr := aws.New(zone, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	startErr := client.DeleteLs(name, IpName)
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

package controller

import (
	"github.com/Yuzuki616/Aws-Panel/aws"
	"github.com/Yuzuki616/Aws-Panel/data"
	"github.com/Yuzuki616/Aws-Panel/request"
	"github.com/gin-gonic/gin"
)

func GetRegions(c *gin.Context) {
	username := c.GetString("username")
	params := request.GetRegions{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	secret, _ := data.GetSecret(username, params.SecretName)
	client, newErr := aws.New("ap-northeast-1", secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	regions, GetRegionsErr := client.GetRegions()
	if GetRegionsErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  GetRegionsErr.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": regions,
	})
}

func CreateLightsail(c *gin.Context) {
	username := c.GetString("username")
	params := request.CreateLightsail{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	secret, _ := data.GetSecret(username, params.SecretName)
	client, newErr := aws.New(params.Zone, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	createRt, createErr := client.CreateLs(params.Name,
		params.AvailabilityZone, params.BlueprintId, params.BundleId)
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
		"data": createRt.Key,
	})
}

func OpenLightsailPorts(c *gin.Context) {
	username := c.GetString("username")
	p := request.OpenLightsailPorts{}
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	secret, _ := data.GetSecret(username, p.SecretName)
	client, newErr := aws.New(p.Zone, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	openErr := client.OpenLsPorts(p.Name)
	if openErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  openErr.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "开放成功",
	})
}

func ListLightsail(c *gin.Context) {
	username := c.GetString("username")
	p := request.LightsailAction{}
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	secret, _ := data.GetSecret(username, p.SecretName)
	client, newErr := aws.New(p.Zone, secret.SecretId, secret.Secret, "")
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
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": instances,
	})
}

func GetLightsailInfo(c *gin.Context) {
	username := c.GetString("username")
	p := request.LightsailAction{}
	name := c.PostForm("name")
	secret, _ := data.GetSecret(username, p.SecretName)
	client, newErr := aws.New(p.Zone, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	info, GetInfoErr := client.GetLsInfo(name)
	if GetInfoErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  GetInfoErr.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "删除成功",
		"data": info,
	})
}

func StartLightsail(c *gin.Context) {
	username := c.GetString("username")
	p := request.LightsailAction{}
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	secret, _ := data.GetSecret(username, p.SecretName)
	client, newErr := aws.New(p.Zone, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	startErr := client.StartLs(p.Name)
	if startErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  startErr.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "启动成功",
	})
}

func StopLightsail(c *gin.Context) {
	username := c.GetString("username")
	p := request.LightsailAction{}
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	secret, _ := data.GetSecret(username, p.SecretName)
	client, newErr := aws.New(p.Zone, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	stopErr := client.StopLs(p.Name)
	if stopErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  stopErr.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "停止成功",
	})
}

func RebootLightsail(c *gin.Context) {
	username := c.GetString("username")
	p := request.LightsailAction{}
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	secret, _ := data.GetSecret(username, p.SecretName)
	client, newErr := aws.New(p.Zone, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	stopErr := client.RebootLs(p.Name)
	if stopErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  stopErr.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "停止成功",
	})
}

func ChangeLightsailIp(c *gin.Context) {
	username := c.GetString("username")
	p := request.LightsailAction{}
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	secret, _ := data.GetSecret(username, p.SecretName)
	client, newErr := aws.New(p.Zone, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	changeIpErr := client.ChangeLsIp(p.Name)
	if changeIpErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  changeIpErr.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "更换IP成功",
	})
}

func DeleteLightsail(c *gin.Context) {
	username := c.GetString("username")
	p := request.DeleteLightsail{}
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	secret, _ := data.GetSecret(username, p.SecretName)
	client, newErr := aws.New(p.Zone, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	startErr := client.DeleteLs(p.Name, p.ResourceName)
	if startErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  startErr.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}

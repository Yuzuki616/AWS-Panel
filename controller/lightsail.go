package controller

import (
	"github.com/Yuzuki616/Aws-Panel/aws"
	"github.com/Yuzuki616/Aws-Panel/data"
	"github.com/Yuzuki616/Aws-Panel/request"
	"github.com/gin-gonic/gin"
)

func GetRegions(c *gin.Context) error {
	email := c.GetString("email")
	params := request.GetRegions{}
	if err := c.ShouldBind(&params); err != nil {
		return err
	}
	secret, _ := data.GetSecret(email, params.SecretName)
	client, newErr := aws.New("ap-northeast-1", secret.SecretId, secret.Secret, "")
	if newErr != nil {
		return newErr
	}
	regions, getRegionsErr := client.GetRegions()
	if getRegionsErr != nil {
		return getRegionsErr
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": regions,
	})
	return nil
}

func CreateLightsail(c *gin.Context) error {
	email := c.GetString("email")
	params := request.CreateLightsail{}
	if err := c.ShouldBind(&params); err != nil {
		return err
	}
	secret, _ := data.GetSecret(email, params.SecretName)
	client, newErr := aws.New(params.Zone, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		return newErr
	}
	createRt, createErr := client.CreateLs(params.Name,
		params.AvailabilityZone, params.BlueprintId, params.BundleId)
	if createErr != nil {
		return createErr
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "创建成功",
		"data": createRt.Key,
	})
	return nil
}

func OpenLightsailPorts(c *gin.Context) error {
	email := c.GetString("email")
	p := request.OpenLightsailPorts{}
	if err := c.ShouldBind(&p); err != nil {
		return err
	}
	secret, _ := data.GetSecret(email, p.SecretName)
	client, newErr := aws.New(p.Zone, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		return newErr
	}
	openErr := client.OpenLsPorts(p.Name)
	if openErr != nil {
		return openErr
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "开放成功",
	})
	return nil
}

func ListLightsail(c *gin.Context) error {
	email := c.GetString("email")
	p := request.LightsailAction{}
	if err := c.ShouldBind(&p); err != nil {
		return err
	}
	secret, _ := data.GetSecret(email, p.SecretName)
	client, newErr := aws.New(p.Zone, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		return newErr
	}
	listRt, listErr := client.ListLs()
	if listErr != nil {
		return listErr
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
	return nil
}

func GetLightsailInfo(c *gin.Context) error {
	email := c.GetString("email")
	p := request.LightsailAction{}
	name := c.PostForm("name")
	secret, _ := data.GetSecret(email, p.SecretName)
	client, newErr := aws.New(p.Zone, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		return newErr
	}
	info, getInfoErr := client.GetLsInfo(name)
	if getInfoErr != nil {
		return getInfoErr
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "删除成功",
		"data": info,
	})
	return nil
}

func StartLightsail(c *gin.Context) error {
	email := c.GetString("email")
	p := request.LightsailAction{}
	if err := c.ShouldBind(&p); err != nil {
		return err
	}
	secret, _ := data.GetSecret(email, p.SecretName)
	client, newErr := aws.New(p.Zone, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		return newErr
	}
	startErr := client.StartLs(p.Name)
	if startErr != nil {
		return startErr
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "启动成功",
	})
	return nil
}

func StopLightsail(c *gin.Context) error {
	email := c.GetString("email")
	p := request.LightsailAction{}
	if err := c.ShouldBind(&p); err != nil {
		return err
	}
	secret, _ := data.GetSecret(email, p.SecretName)
	client, newErr := aws.New(p.Zone, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		return newErr
	}
	stopErr := client.StopLs(p.Name)
	if stopErr != nil {
		return stopErr
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "停止成功",
	})
	return nil
}

func RebootLightsail(c *gin.Context) error {
	email := c.GetString("email")
	p := request.LightsailAction{}
	if err := c.ShouldBind(&p); err != nil {
		return err
	}
	secret, _ := data.GetSecret(email, p.SecretName)
	client, newErr := aws.New(p.Zone, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		return newErr
	}
	stopErr := client.RebootLs(p.Name)
	if stopErr != nil {
		return stopErr
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "停止成功",
	})
	return nil
}

func ChangeLightsailIp(c *gin.Context) error {
	email := c.GetString("email")
	p := request.LightsailAction{}
	if err := c.ShouldBind(&p); err != nil {
		return err
	}
	secret, _ := data.GetSecret(email, p.SecretName)
	client, newErr := aws.New(p.Zone, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		return newErr
	}
	changeIpErr := client.ChangeLsIp(p.Name)
	if changeIpErr != nil {
		return changeIpErr
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "更换IP成功",
	})
	return nil
}

func DeleteLightsail(c *gin.Context) error {
	email := c.GetString("email")
	p := request.DeleteLightsail{}
	if err := c.ShouldBind(&p); err != nil {
		return err
	}
	secret, _ := data.GetSecret(email, p.SecretName)
	client, newErr := aws.New(p.Zone, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		return newErr
	}
	startErr := client.DeleteLs(p.Name, p.ResourceName)
	if startErr != nil {
		return startErr
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
	return nil
}

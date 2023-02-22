package controller

import (
	"errors"
	"github.com/Yuzuki616/Aws-Panel/aws"
	"github.com/Yuzuki616/Aws-Panel/data"
	"github.com/Yuzuki616/Aws-Panel/request"
	"github.com/gin-gonic/gin"
)

func CreateEc2(c *gin.Context) error {
	email := c.GetString("email")
	p := request.CreateEc2{}
	if err := c.ShouldBind(&p); err != nil {
		return err
	}
	secret, _ := data.GetSecret(email, p.SecretName)
	client, newErr := aws.New(p.Region, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		return newErr
	}
	amiTmp, amiErr := client.GetAmiId(p.Ami)
	if amiErr != nil {
		return amiErr
	}
	if amiTmp == "" {
		return errors.New("get ami ID error: not found ami")
	}
	creRt, creErr := client.CreateEc2(amiTmp, p.Ec2Type, p.Ec2Name, p.Userdata, p.Disk)
	if creErr != nil {
		return creErr
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "创建成功",
		"data": creRt.Key,
	})
	return nil
}

func ListEc2(c *gin.Context) error {
	email := c.GetString("email")
	p := request.ListEc2{}
	if err := c.ShouldBind(&p); err != nil {
		return err
	}
	secret, _ := data.GetSecret(email, p.SecretName)
	client, newErr := aws.New(p.Region, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		return newErr
	}
	ec2Info, listErr := client.ListEc2()
	if listErr != nil {
		return listErr
	}

	c.JSON(200, gin.H{
		"code": 200,
		"data": ec2Info,
	})
	return nil
}

func GetEc2Info(c *gin.Context) error {
	email := c.GetString("email")
	p := request.Ec2Action{}
	if err := c.ShouldBind(&p); err != nil {
		return err
	}
	secret, _ := data.GetSecret(email, p.SecretName)
	client, newErr := aws.New(p.Region, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		return newErr
	}
	ec2Info, getErr := client.GetEc2Info(p.Ec2Id)
	if getErr != nil {
		return getErr
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": ec2Info,
	})
	return nil
}

func ChangeEc2Ip(c *gin.Context) error {
	email := c.GetString("email")
	p := request.Ec2Action{}
	if err := c.ShouldBind(&p); err != nil {
		return err
	}
	secret, _ := data.GetSecret(email, p.SecretName)
	client, newErr := aws.New(p.Region, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		return newErr
	}
	newIp, changeErr := client.ChangeEc2Ip(p.Ec2Id)
	if changeErr != nil {
		return changeErr
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "更换IP成功",
		"data": newIp,
	})
	return nil
}

func StopEc2(c *gin.Context) error {
	email := c.GetString("email")
	p := request.Ec2Action{}
	if err := c.ShouldBind(&p); err != nil {
		return err
	}
	secret, _ := data.GetSecret(email, p.SecretName)
	client, newErr := aws.New(p.Region, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		return newErr
	}
	stopErr := client.StopEc2(p.Ec2Id)
	if stopErr != nil {
		return stopErr
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "启动失败",
	})
	return nil
}

func StartEc2(c *gin.Context) error {
	email := c.GetString("email")
	p := request.Ec2Action{}
	if err := c.ShouldBind(&p); err != nil {
		return err
	}
	secret, _ := data.GetSecret(email, p.SecretName)
	client, newErr := aws.New(p.Region, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		return newErr
	}
	startErr := client.StartEc2(p.Ec2Id)
	if startErr != nil {
		return startErr
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "启动成功",
	})
	return nil
}

func RebootEc2(c *gin.Context) error {
	email := c.GetString("email")
	p := request.Ec2Action{}
	if err := c.ShouldBind(&p); err != nil {
		return err
	}
	secret, _ := data.GetSecret(email, p.SecretName)
	client, newErr := aws.New(p.Region, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		return newErr
	}
	rebootErr := client.RebootEc2(p.Ec2Id)
	if rebootErr != nil {
		return rebootErr
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "重启成功",
	})
	return nil
}

func DeleteEc2(c *gin.Context) error {
	email := c.GetString("email")
	p := request.Ec2Action{}
	if err := c.ShouldBind(&p); err != nil {
		return err
	}
	secret, _ := data.GetSecret(email, p.SecretName)
	client, newErr := aws.New(p.Region, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		return newErr
	}
	delErr := client.DeleteEc2(p.Ec2Id)
	if delErr != nil {
		return delErr
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
	return nil
}

func CreateEc2SshKey(c *gin.Context) error {
	email := c.GetString("email")
	p := request.CreateEc2SshKey{}
	if err := c.ShouldBind(&p); err != nil {
		return err
	}
	secret, _ := data.GetSecret(email, p.SecretName)
	client, newErr := aws.New(p.Region, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		return newErr
	}
	key, err := client.CreateEc2SshKey(p.KeyName)
	if err != nil {
		return err
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "创建成功",
		"data": key,
	})
	return nil
}

func ListEc2SshKey(c *gin.Context) error {
	email := c.GetString("email")
	p := request.ListEc2SshKey{}
	if err := c.ShouldBind(&p); err != nil {
		return err
	}
	secret, _ := data.GetSecret(email, p.SecretName)
	client, newErr := aws.New(p.Region, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		return newErr
	}
	keys, err := client.ListEc2SshKey()
	if err != nil {
		return err
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": keys,
	})
	return nil
}

func DeleteEc2SshKey(c *gin.Context) error {
	email := c.GetString("email")
	p := request.DeleteEc2SshKey{}
	if err := c.ShouldBind(&p); err != nil {
		return err
	}
	secret, _ := data.GetSecret(email, p.SecretName)
	client, newErr := aws.New(p.Region, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		return newErr
	}
	err := client.DeleteEc2SshKey(p.KeyName)
	if err != nil {
		return err
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
	return nil
}

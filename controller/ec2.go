package controller

import (
	"github.com/Yuzuki616/Aws-Panel/aws"
	"github.com/Yuzuki616/Aws-Panel/data"
	"github.com/Yuzuki616/Aws-Panel/request"
	"github.com/gin-gonic/gin"
)

func CreateEc2(c *gin.Context) {
	username := c.GetString("username")
	p := request.CreateEc2{}
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	secret, _ := data.GetSecret(username, p.SecretName)
	client, newErr := aws.New(p.Region, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	amiTmp, amiErr := client.GetAmiId(p.Ami)
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
	creRt, creErr := client.CreateEc2(amiTmp, p.Ec2Type, p.Ec2Name, p.Userdata, p.Disk)
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
	username := c.GetString("username")
	p := request.ListEc2{}
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	secret, _ := data.GetSecret(username, p.SecretName)
	client, newErr := aws.New(p.Region, secret.SecretId, secret.Secret, "")
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
	username := c.GetString("username")
	p := request.Ec2Action{}
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	secret, _ := data.GetSecret(username, p.SecretName)
	client, newErr := aws.New(p.Region, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	ec2Info, getErr := client.GetEc2Info(p.Ec2Id)
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
	username := c.GetString("username")
	p := request.Ec2Action{}
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	secret, _ := data.GetSecret(username, p.SecretName)
	client, newErr := aws.New(p.Region, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	newIp, changeErr := client.ChangeEc2Ip(p.Ec2Id)
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
	username := c.GetString("username")
	p := request.Ec2Action{}
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	secret, _ := data.GetSecret(username, p.SecretName)
	client, newErr := aws.New(p.Region, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	stopErr := client.StopEc2(p.Ec2Id)
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
	username := c.GetString("username")
	p := request.Ec2Action{}
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	secret, _ := data.GetSecret(username, p.SecretName)
	client, newErr := aws.New(p.Region, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	startErr := client.StartEc2(p.Ec2Id)
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

func RebootEc2(c *gin.Context) {
	username := c.GetString("username")
	p := request.Ec2Action{}
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	secret, _ := data.GetSecret(username, p.SecretName)
	client, newErr := aws.New(p.Region, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	rebootErr := client.RebootEc2(p.Ec2Id)
	if rebootErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  rebootErr.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "重启成功",
	})
}

func DeleteEc2(c *gin.Context) {
	username := c.GetString("username")
	p := request.Ec2Action{}
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	secret, _ := data.GetSecret(username, p.SecretName)
	client, newErr := aws.New(p.Region, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	delErr := client.DeleteEc2(p.Ec2Id)
	if delErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  delErr.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}

func CreateEc2SshKey(c *gin.Context) {
	username := c.GetString("username")
	p := request.CreateEc2SshKey{}
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	secret, _ := data.GetSecret(username, p.SecretName)
	client, newErr := aws.New(p.Region, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	key, err := client.CreateEc2SshKey(p.KeyName)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "创建成功",
		"data": key,
	})
}

func ListEc2SshKey(c *gin.Context) {
	username := c.GetString("username")
	p := request.ListEc2SshKey{}
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	secret, _ := data.GetSecret(username, p.SecretName)
	client, newErr := aws.New(p.Region, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	keys, err := client.ListEc2SshKey()
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": keys,
	})
}

func DeleteEc2SshKey(c *gin.Context) {
	username := c.GetString("username")
	p := request.DeleteEc2SshKey{}
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	secret, _ := data.GetSecret(username, p.SecretName)
	client, newErr := aws.New(p.Region, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	err := client.DeleteEc2SshKey(p.KeyName)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}

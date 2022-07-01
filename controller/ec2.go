package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yuzuki999/Aws-Panel/aws"
	"github.com/yuzuki999/Aws-Panel/data"
	"strconv"
)

func CreateEc2(c *gin.Context) {
	username := GetLoginUser(c)
	if username == "" {
		return
	}
	secretName := c.PostForm("secretName")
	region := c.PostForm("region")
	ami := c.PostForm("ami")
	ec2Type := c.PostForm("ec2Type")
	ec2Name := c.PostForm("ec2Name")
	if secretName == "" || region == "" || ami == "" || ec2Type == "" || ec2Name == "" {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "信息填写不完整",
		})
		return
	}
	disk, _ := strconv.ParseInt(c.PostForm("disk"), 10, 64)
	secret, _ := data.GetSecret(username, secretName)
	client, newErr := aws.New(region, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	amiTmp, amiErr := client.GetAmiId(ami)
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
	creRt, creErr := client.CreateEc2(amiTmp, ec2Type, ec2Name, disk)
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
		"data": *creRt.Key,
	})
}

func ListEc2(c *gin.Context) {
	username := GetLoginUser(c)
	if username == "" {
		return
	}
	secretName := c.PostForm("secretName")
	region := c.PostForm("region")
	if secretName == "" || region == "" {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "信息填写不完整",
		})
		return
	}
	secret, _ := data.GetSecret(username, secretName)
	client, newErr := aws.New(region, secret.SecretId, secret.Secret, "")
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
	var ec2Instances []*aws.Ec2Info
	if ec2Info == nil {
		ec2Instances = nil
	} else {
		for i := range ec2Info {
			ec2Instances = append(ec2Instances, &aws.Ec2Info{
				Name:       aws.CheckNameNil(ec2Info[i].Instances[0].Tags),
				Status:     ec2Info[i].Instances[0].State.Name,
				Type:       ec2Info[i].Instances[0].InstanceType,
				InstanceId: ec2Info[i].Instances[0].InstanceId,
				Ip:         ec2Info[i].Instances[0].PublicIpAddress,
			})
		}
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": ec2Instances,
	})
}

func GetEc2Info(c *gin.Context) {
	username := GetLoginUser(c)
	if username == "" {
		return
	}
	secretName := c.PostForm("secretName")
	region := c.PostForm("region")
	ec2Id := c.PostForm("ec2Id")
	if secretName == "" || region == "" || ec2Id == "" {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "信息填写不完整",
		})
		return
	}
	secret, _ := data.GetSecret(username, secretName)
	client, newErr := aws.New(region, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	ec2Info, getErr := client.GetEc2Info(ec2Id)
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
	secretName := c.PostForm("secretName")
	region := c.PostForm("region")
	ec2Id := c.PostForm("ec2Id")
	if secretName == "" || region == "" || ec2Id == "" {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "信息填写不完整",
		})
		return
	}
	secret, _ := data.GetSecret(username, secretName)
	client, newErr := aws.New(region, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	newIp, changeErr := client.ChangeEc2Ip(ec2Id)
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
	secretName := c.PostForm("secretName")
	region := c.PostForm("region")
	ec2Id := c.PostForm("ec2Id")
	if secretName == "" || region == "" || ec2Id == "" {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "信息填写不完整",
		})
		return
	}
	secret, _ := data.GetSecret(username, secretName)
	client, newErr := aws.New(region, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	stopErr := client.StopEc2(ec2Id)
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
	secretName := c.PostForm("secretName")
	region := c.PostForm("region")
	ec2Id := c.PostForm("ec2Id")
	if secretName == "" || region == "" || ec2Id == "" {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "信息填写不完整",
		})
		return
	}
	secret, _ := data.GetSecret(username, secretName)
	client, newErr := aws.New(region, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
	}
	startErr := client.StartEc2(ec2Id)
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
	secretName := c.PostForm("secretName")
	region := c.PostForm("region")
	ec2Id := c.PostForm("ec2Id")
	if secretName == "" || region == "" || ec2Id == "" {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "信息填写不完整",
		})
		return
	}
	secret, _ := data.GetSecret(username, secretName)
	client, newErr := aws.New(region, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	rebootErr := client.RebootEc2(ec2Id)
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
	secretName := c.PostForm("secretName")
	region := c.PostForm("region")
	ec2Id := c.PostForm("ec2Id")
	if secretName == "" || region == "" || ec2Id == "" {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "信息填写不完整",
		})
		return
	}
	secret, _ := data.GetSecret(username, secretName)
	client, newErr := aws.New(region, secret.SecretId, secret.Secret, "")
	if newErr != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  newErr.Error(),
		})
		return
	}
	delErr := client.DeleteEc2(ec2Id)
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

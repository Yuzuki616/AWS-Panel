package request

type CreateEc2 struct {
	SecretName string `form:"secretName" binding:"required"`
	Region     string `form:"region" binding:"required"`
	Ami        string `form:"ami" binding:"required"`
	Ec2Type    string `form:"ec2Type" binding:"required"`
	Ec2Name    string `form:"ec2Name" binding:"required"`
	Disk       int64  `form:"Disk" binding:"required"`
	Userdata   string `form:"userdata" binding:"required"`
}

type ListEc2 struct {
	SecretName string `form:"secretName" binding:"required"`
	Region     string `form:"region" binding:"required"`
}

type Ec2Action struct {
	SecretName string `form:"secretName" binding:"required"`
	Region     string `form:"region" binding:"required"`
	Ec2Id      string `form:"ec2Id" binding:"required"`
}

type CreateEc2SshKey struct {
	SecretName string `form:"secretName" binding:"required"`
	Region     string `form:"region" binding:"required"`
	KeyName    string `form:"keyName" binding:"required"`
}

type ListEc2SshKey struct {
	SecretName string `form:"secretName" binding:"required"`
	Region     string `form:"region" binding:"required"`
}
type DeleteEc2SshKey struct {
	SecretName string `form:"secretName" binding:"required"`
	Region     string `form:"region" binding:"required"`
	KeyName    string `form:"keyName" binding:"required"`
}

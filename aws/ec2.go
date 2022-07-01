package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"time"
)

type Ec2Info struct {
	Name       *string
	InstanceId *string
	Status     *string
	Type       *string
	Ip         *string
	Key        *string
}

func (p *Aws) CreateEc2(Ami string, Ec2Type string, Name string, DiskSize int64) (*Ec2Info, error) {
	svc := ec2.New(p.Sess)
	dateName := Name + time.Unix(time.Now().Unix(), 0).Format("_2006-01-02_15:04:05")
	keyRt, keyErr := svc.CreateKeyPair(&ec2.CreateKeyPairInput{KeyName: &dateName})
	if keyErr != nil {
		return nil, fmt.Errorf("create key error: %v", keyErr)
	} //创建ssh密钥
	secRt, secErr := svc.CreateSecurityGroup(&ec2.CreateSecurityGroupInput{
		GroupName:   aws.String(dateName + "security"),
		Description: aws.String("A security group for aws manger bot"),
	}) //创建安全组
	if secErr != nil {
		return nil, fmt.Errorf("create security group error: %v", secErr)
	}
	_, authSecInErr := svc.AuthorizeSecurityGroupIngress(&ec2.AuthorizeSecurityGroupIngressInput{
		GroupId: secRt.GroupId,
		IpPermissions: []*ec2.IpPermission{
			{
				IpProtocol: aws.String("-1"),
				IpRanges: []*ec2.IpRange{
					{
						CidrIp: aws.String("0.0.0.0/0"),
					},
				},
				FromPort: aws.Int64(-1),
				ToPort:   aws.Int64(-1),
			},
		},
	}) //添加入站规则
	if authSecInErr != nil {
		return nil, fmt.Errorf(
			"authorize security group ingress error: %v", authSecInErr)
	}
	ebs, ebsErr := p.getAmiEbsMap(Ami)
	if ebsErr != nil {
		return nil, ebsErr
	}
	ebs[0].Ebs.VolumeSize = &DiskSize
	runRt, runErr := svc.RunInstances(&ec2.RunInstancesInput{
		ImageId:             aws.String(Ami),
		InstanceType:        aws.String(Ec2Type),
		MinCount:            aws.Int64(1),
		MaxCount:            aws.Int64(1),
		KeyName:             &dateName,
		BlockDeviceMappings: ebs,
		SecurityGroupIds:    []*string{secRt.GroupId},
	}) //创建ec2实例
	if runErr != nil {
		return nil, fmt.Errorf("run instances error: %v", runErr)
	}
	_, tagErr := svc.CreateTags(&ec2.CreateTagsInput{
		Resources: []*string{runRt.Instances[0].InstanceId},
		Tags: []*ec2.Tag{
			{
				Key:   aws.String("Name"),
				Value: aws.String(Name),
			},
		},
	}) //创建标签
	if tagErr != nil {
		return nil, fmt.Errorf("create tag error: %v", tagErr)
	}
	return &Ec2Info{
		Name:       &Name,
		InstanceId: runRt.Instances[0].InstanceId,
		Status:     runRt.Instances[0].State.Name,
		Key:        keyRt.KeyMaterial,
	}, nil
}

func (p *Aws) ChangeEc2Ip(InstanceId string) (*string, error) {
	svc := ec2.New(p.Sess)
	desRt, desErr := svc.DescribeAddresses(&ec2.DescribeAddressesInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("instance-id"),
				Values: []*string{aws.String(InstanceId)},
			},
		},
	})
	if desErr != nil {
		return nil, fmt.Errorf("describe addresses error: %v", desErr)
	}
	if len(desRt.Addresses) != 0 {
		_, relErr := svc.ReleaseAddress(&ec2.ReleaseAddressInput{AllocationId: desRt.Addresses[0].AllocationId})
		if relErr != nil {
			return nil, fmt.Errorf("release ip error: %s", relErr)
		}
	}
	allRt, allErr := svc.AllocateAddress(&ec2.AllocateAddressInput{})
	if allErr != nil {
		return nil, fmt.Errorf("allocate address error: %v", allErr)
	}
	_, assErr := svc.AssociateAddress(&ec2.AssociateAddressInput{
		AllocationId: allRt.AllocationId,
		InstanceId:   aws.String(InstanceId),
	})
	if assErr != nil {
		return nil, fmt.Errorf("associate address error: %v", assErr)
	}
	return allRt.PublicIp, nil
}

func (p *Aws) GetEc2Info(InstanceId string) (*Ec2Info, error) {
	svc := ec2.New(p.Sess)
	rt, err := svc.DescribeInstances(&ec2.DescribeInstancesInput{InstanceIds: []*string{aws.String(InstanceId)}})
	if err != nil {
		return nil, err
	}
	return &Ec2Info{
		Name:       CheckNameNil(rt.Reservations[0].Instances[0].Tags),
		InstanceId: rt.Reservations[0].Instances[0].InstanceId,
		Status:     rt.Reservations[0].Instances[0].State.Name,
		Ip:         rt.Reservations[0].Instances[0].PublicIpAddress,
	}, nil
}

func (p *Aws) ListEc2() ([]*ec2.Reservation, error) {
	svc := ec2.New(p.Sess)
	rt, err := svc.DescribeInstances(&ec2.DescribeInstancesInput{MaxResults: aws.Int64(100)})
	if err != nil {
		return nil, err
	}
	return rt.Reservations, nil
}

func (p *Aws) StartEc2(InstanceId string) error {
	svc := ec2.New(p.Sess)
	_, err := svc.StartInstances(&ec2.StartInstancesInput{InstanceIds: []*string{aws.String(InstanceId)}})
	if err != nil {
		return err
	}
	return nil
}

func (p *Aws) StopEc2(InstanceId string) error {
	svc := ec2.New(p.Sess)
	_, err := svc.StopInstances(&ec2.StopInstancesInput{InstanceIds: []*string{aws.String(InstanceId)}})
	if err != nil {
		return err
	}
	return nil
}

func (p *Aws) RebootEc2(InstanceId string) error {
	svc := ec2.New(p.Sess)
	_, err := svc.RebootInstances(&ec2.RebootInstancesInput{InstanceIds: []*string{aws.String(InstanceId)}})
	if err != nil {
		return err
	}
	return nil
}

func (p *Aws) DeleteEc2(InstanceId string) error {
	svc := ec2.New(p.Sess)
	ip, ipErr := svc.DescribeAddresses(&ec2.DescribeAddressesInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("instance-id"),
				Values: []*string{aws.String(InstanceId)},
			},
		},
	})
	if ipErr != nil {
		return fmt.Errorf("get ip error: %v", ipErr)
	}
	if len(ip.Addresses) != 0 {
		_, relErr := svc.ReleaseAddress(&ec2.ReleaseAddressInput{AllocationId: ip.Addresses[0].AssociationId})
		if relErr != nil {
			return fmt.Errorf("release ip error: %v", relErr)
		}
	}
	_, err := svc.TerminateInstances(&ec2.TerminateInstancesInput{InstanceIds: []*string{aws.String(InstanceId)}})
	if err != nil {
		return fmt.Errorf("terminate instance error: %v", err)
	}
	return nil
}

func (p *Aws) GetAmiId(AmiName string) (string, error) {
	svc := ec2.New(p.Sess)
	ami, err := svc.DescribeImages(&ec2.DescribeImagesInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("name"),
				Values: []*string{aws.String(AmiName)},
			},
			{
				Name:   aws.String("architecture"),
				Values: []*string{aws.String("x86_64")},
			},
		}})
	if err != nil {
		return "", err
	}
	return *ami.Images[0].ImageId, nil
}

func (p *Aws) getAmiEbsMap(AmiId string) ([]*ec2.BlockDeviceMapping, error) {
	svc := ec2.New(p.Sess)
	ami, err := svc.DescribeImages(&ec2.DescribeImagesInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("image-id"),
				Values: []*string{aws.String(AmiId)},
			},
			{
				Name:   aws.String("architecture"),
				Values: []*string{aws.String("x86_64")},
			},
		}})
	if err != nil {
		return nil, err
	}
	return ami.Images[0].BlockDeviceMappings, nil
}

func (p *Aws) GetWindowsPassword(InstanceId string) (*ec2.GetPasswordDataOutput, error) {
	svc := ec2.New(p.Sess)
	rt, err := svc.GetPasswordData(&ec2.GetPasswordDataInput{InstanceId: aws.String(InstanceId)})
	if err != nil {
		return nil, err
	}
	return rt, nil
}

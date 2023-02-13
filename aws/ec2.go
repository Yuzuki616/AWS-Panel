package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"reflect"
	"time"
)

type Ec2Info struct {
	Name       string
	InstanceId string
	Status     string
	Type       string
	Ip         string
	Key        string
}

func (a *Aws) CreateEc2(Ami, Ec2Type, Name, userdata string, DiskSize int64) (*Ec2Info, error) {
	svc := ec2.New(a.Sess)
	dateName := Name + time.Unix(time.Now().Unix(), 0).Format("_2006-01-02_15:04:05")
	key, keyErr := a.CreateEc2SshKey(dateName + "_key")
	if keyErr != nil {
		return nil, fmt.Errorf("create key error: %v", keyErr)
	} //创建ssh密钥
	secRt, secErr := svc.CreateSecurityGroup(&ec2.CreateSecurityGroupInput{
		GroupName:   aws.String(dateName + "_security"),
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
	ebs, ebsErr := a.getAmiEbsMap(Ami)
	if ebsErr != nil {
		return nil, ebsErr
	}
	ebs[0].Ebs.VolumeSize = &DiskSize
	runRt, runErr := svc.RunInstances(&ec2.RunInstancesInput{
		ImageId:             &Ami,
		InstanceType:        &Ec2Type,
		MinCount:            aws.Int64(1),
		MaxCount:            aws.Int64(1),
		KeyName:             &dateName,
		BlockDeviceMappings: ebs,
		SecurityGroupIds:    []*string{secRt.GroupId},
		UserData:            &userdata,
	}) //创建ec2实例
	if runErr != nil {
		return nil, fmt.Errorf("run instances error: %v", runErr)
	}
	_, tagErr := svc.CreateTags(&ec2.CreateTagsInput{
		Resources: []*string{runRt.Instances[0].InstanceId},
		Tags: []*ec2.Tag{
			{
				Key:   aws.String("Name"),
				Value: &Name,
			},
		},
	}) //创建标签
	if tagErr != nil {
		return nil, fmt.Errorf("create tag error: %v", tagErr)
	}
	return &Ec2Info{
		Name:       Name,
		InstanceId: *runRt.Instances[0].InstanceId,
		Status:     *runRt.Instances[0].State.Name,
		Key:        key,
	}, nil
}

func (a *Aws) ChangeEc2Ip(InstanceId string) (*string, error) {
	svc := ec2.New(a.Sess)
	desRt, desErr := svc.DescribeAddresses(&ec2.DescribeAddressesInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("instance-id"),
				Values: []*string{&InstanceId},
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
		InstanceId:   &InstanceId,
	})
	if assErr != nil {
		return nil, fmt.Errorf("associate address error: %v", assErr)
	}
	return allRt.PublicIp, nil
}

func getNameFromTags(v []*ec2.Tag) string {
	if reflect.ValueOf(v).IsNil() {
		return ""
	}
	return *v[0].Value
}

func (a *Aws) GetEc2Info(InstanceId string) (*Ec2Info, error) {
	svc := ec2.New(a.Sess)
	rt, err := svc.DescribeInstances(&ec2.DescribeInstancesInput{InstanceIds: []*string{aws.String(InstanceId)}})
	if err != nil {
		return nil, err
	}
	return &Ec2Info{
		Name:       getNameFromTags(rt.Reservations[0].Instances[0].Tags),
		InstanceId: *rt.Reservations[0].Instances[0].InstanceId,
		Status:     *rt.Reservations[0].Instances[0].State.Name,
		Ip:         *rt.Reservations[0].Instances[0].PublicIpAddress,
	}, nil
}

func (a *Aws) ListEc2() ([]Ec2Info, error) {
	svc := ec2.New(a.Sess)
	rt, err := svc.DescribeInstances(&ec2.DescribeInstancesInput{MaxResults: aws.Int64(100)})
	if err != nil {
		return nil, err
	}
	ec2Instances := make([]Ec2Info, 0, len(rt.Reservations))
	for _, v := range rt.Reservations {
		ec2Instances = append(ec2Instances, Ec2Info{
			Name:       getNameFromTags(v.Instances[0].Tags),
			Status:     *v.Instances[0].State.Name,
			Type:       *v.Instances[0].InstanceType,
			InstanceId: *v.Instances[0].InstanceId,
			Ip:         *v.Instances[0].PublicIpAddress,
		})
	}
	return ec2Instances, nil
}

func (a *Aws) StartEc2(InstanceId string) error {
	svc := ec2.New(a.Sess)
	_, err := svc.StartInstances(&ec2.StartInstancesInput{
		InstanceIds: []*string{&InstanceId}})
	if err != nil {
		return err
	}
	return nil
}

func (a *Aws) StopEc2(InstanceId string) error {
	svc := ec2.New(a.Sess)
	_, err := svc.StopInstances(&ec2.StopInstancesInput{
		InstanceIds: []*string{&InstanceId}})
	if err != nil {
		return err
	}
	return nil
}

func (a *Aws) RebootEc2(InstanceId string) error {
	svc := ec2.New(a.Sess)
	_, err := svc.RebootInstances(&ec2.RebootInstancesInput{
		InstanceIds: []*string{&InstanceId}})
	if err != nil {
		return err
	}
	return nil
}

func (a *Aws) DeleteEc2(InstanceId string) error {
	svc := ec2.New(a.Sess)
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
		_, relErr := svc.ReleaseAddress(&ec2.ReleaseAddressInput{
			AllocationId: ip.Addresses[0].AssociationId})
		if relErr != nil {
			return fmt.Errorf("release ip error: %v", relErr)
		}
	}
	_, err := svc.TerminateInstances(&ec2.TerminateInstancesInput{
		InstanceIds: []*string{aws.String(InstanceId)}})
	if err != nil {
		return fmt.Errorf("terminate instance error: %v", err)
	}
	return nil
}

func (a *Aws) GetAmiId(AmiName string) (string, error) {
	svc := ec2.New(a.Sess)
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

func (a *Aws) getAmiEbsMap(AmiId string) ([]*ec2.BlockDeviceMapping, error) {
	svc := ec2.New(a.Sess)
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

func (a *Aws) GetEc2WindowsPassword(InstanceId string) (*ec2.GetPasswordDataOutput, error) {
	svc := ec2.New(a.Sess)
	rt, err := svc.GetPasswordData(&ec2.GetPasswordDataInput{
		InstanceId: &InstanceId})
	if err != nil {
		return nil, err
	}
	return rt, nil
}

func (a *Aws) CreateEc2SshKey(name string) (string, error) {
	svc := ec2.New(a.Sess)
	rt, err := svc.CreateKeyPair(&ec2.CreateKeyPairInput{KeyName: &name})
	if err != nil {
		return "", err
	}
	return *rt.KeyMaterial, nil
}

func (a *Aws) ListEc2SshKey() ([]*ec2.KeyPairInfo, error) {
	svc := ec2.New(a.Sess)
	rt, err := svc.DescribeKeyPairs(&ec2.DescribeKeyPairsInput{})
	if err != nil {
		return nil, err
	}
	return rt.KeyPairs, nil
}

func (a *Aws) DeleteEc2SshKey(name string) error {
	svc := ec2.New(a.Sess)
	_, err := svc.DeleteKeyPair(&ec2.DeleteKeyPairInput{KeyName: &name})
	if err != nil {
		return err
	}
	return nil
}

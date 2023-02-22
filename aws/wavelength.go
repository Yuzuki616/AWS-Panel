package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"time"
)

func (a *Aws) getVpcId() (string, error) {
	svc := ec2.New(a.Sess)
	vpc, err := svc.DescribeVpcs(&ec2.DescribeVpcsInput{})
	if err != nil {
		return "", err
	}
	return *vpc.Vpcs[0].VpcId, nil
}

func (a *Aws) CreateWl(Zone string) (string, error) {
	svc := ec2.New(a.Sess)
	vpcId, vpcErr := a.getVpcId()
	if vpcErr != nil {
		return "", vpcErr
	}
	sub, subErr := svc.CreateSubnet(&ec2.CreateSubnetInput{
		AvailabilityZone: aws.String(Zone),
		CidrBlock:        aws.String("172.31.128.0/20"),
		VpcId:            aws.String(vpcId),
		TagSpecifications: []*ec2.TagSpecification{
			{
				ResourceType: aws.String("subnet"),
				Tags: []*ec2.Tag{
					{
						Key:   aws.String("Username"),
						Value: aws.String("aws_manger_subnet"),
					},
				},
			},
		},
	})
	if subErr != nil {
		return "", subErr
	}
	ca, caErr := svc.CreateCarrierGateway(&ec2.CreateCarrierGatewayInput{
		VpcId: aws.String(vpcId),
		TagSpecifications: []*ec2.TagSpecification{
			{
				ResourceType: aws.String("carrier-gateway"),
				Tags: []*ec2.Tag{
					{
						Key:   aws.String("Username"),
						Value: aws.String("aws_manger_gateway"),
					},
				},
			},
		},
	})
	if caErr != nil {
		return "", caErr
	}
	route, routeErr := svc.CreateRouteTable(&ec2.CreateRouteTableInput{
		VpcId: aws.String(vpcId),
		TagSpecifications: []*ec2.TagSpecification{
			{
				ResourceType: aws.String("route-table"),
				Tags: []*ec2.Tag{
					{
						Key:   aws.String("Username"),
						Value: aws.String("aws_manger_route"),
					},
				},
			},
		},
	})
	if routeErr != nil {
		return "", routeErr
	}
	_, assErr := svc.AssociateRouteTable(&ec2.AssociateRouteTableInput{
		RouteTableId: route.RouteTable.RouteTableId,
		GatewayId:    ca.CarrierGateway.CarrierGatewayId,
	})
	if assErr != nil {
		return "", assErr
	}
	return *sub.Subnet.SubnetId, nil
}

func (a *Aws) GetSubnetInfo() (*ec2.DescribeSubnetsOutput, error) {
	svc := ec2.New(a.Sess)
	sub, err := svc.DescribeSubnets(&ec2.DescribeSubnetsInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("tag:Username"),
				Values: []*string{aws.String("kddi")},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return sub, nil
}

func (a *Aws) CreateEc2Wl(SubId string, Ami string, Name string, DiskSize int64) (*Ec2Info, error) {
	svc := ec2.New(a.Sess)
	dateName := Name + time.Unix(time.Now().Unix(), 0).Format("_2006-01-02_15:04:05")
	keyRt, keyErr := svc.CreateKeyPair(&ec2.CreateKeyPairInput{KeyName: &dateName})
	if keyErr != nil {
		return nil, keyErr
	} //创建ssh密钥
	secRt, secErr := svc.CreateSecurityGroup(&ec2.CreateSecurityGroupInput{
		GroupName:   aws.String(dateName + "security"),
		Description: aws.String("A security group for aws manger bot"),
	}) //创建安全组
	if secErr != nil {
		return nil, secErr
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
		return nil, authSecInErr
	}
	runRt, runErr := svc.RunInstances(&ec2.RunInstancesInput{
		ImageId:      aws.String(Ami),
		InstanceType: aws.String("t3.medium"),
		MinCount:     aws.Int64(1),
		MaxCount:     aws.Int64(1),
		KeyName:      &dateName,
		NetworkInterfaces: []*ec2.InstanceNetworkInterfaceSpecification{{
			DeviceIndex:               aws.Int64(0),
			SubnetId:                  aws.String(SubId),
			AssociateCarrierIpAddress: aws.Bool(true),
			Groups:                    []*string{secRt.GroupId},
		}},
		BlockDeviceMappings: []*ec2.BlockDeviceMapping{{DeviceName: aws.String("/dev/sda1"),
			Ebs: &ec2.EbsBlockDevice{VolumeSize: aws.Int64(DiskSize)}}},
	}) //创建ec2实例
	if runErr != nil {
		return nil, runErr
	}
	_, tagErr := svc.CreateTags(&ec2.CreateTagsInput{
		Resources: []*string{runRt.Instances[0].InstanceId},
		Tags: []*ec2.Tag{
			{
				Key:   aws.String("Email"),
				Value: aws.String(Name),
			},
		},
	}) //创建标签
	if tagErr != nil {
		return nil, tagErr
	}
	return &Ec2Info{
		Name:       Name,
		Ip:         *runRt.Instances[0].PrivateIpAddress,
		InstanceId: *runRt.Instances[0].InstanceId,
		Status:     *runRt.Instances[0].State.Name,
		Key:        *keyRt.KeyMaterial,
	}, nil
}

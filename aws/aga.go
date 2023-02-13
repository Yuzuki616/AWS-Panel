package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	aga "github.com/aws/aws-sdk-go/service/globalaccelerator"
	"time"
)

type AgaInfo struct {
	Name     *string
	Status   *string
	Arn      string
	Ip       []*aga.IpSet
	Protocol *string
	Port     []*aga.PortOverride
}

func (a *Aws) CreateAga(Name string, Region string, InstanceId string) (*AgaInfo, error) {
	svc := aga.New(a.Sess)
	IdempotencyToken := time.Unix(time.Now().Unix(), 0).Format("2006-01-02_15:04:05")
	createAccRt, createAccErr := svc.CreateAccelerator(&aga.CreateAcceleratorInput{
		Name:             aws.String(Name),
		Enabled:          aws.Bool(false),
		IdempotencyToken: aws.String(IdempotencyToken),
	})
	if createAccErr != nil {
		return nil, createAccErr
	}
	createLiRt, createLiErr := svc.CreateListener(&aga.CreateListenerInput{
		AcceleratorArn: createAccRt.Accelerator.AcceleratorArn,
		PortRanges: []*aga.PortRange{
			{
				FromPort: aws.Int64(1),
				ToPort:   aws.Int64(65535),
			},
		},
		Protocol: aws.String("TCP"),
	})
	if createLiErr != nil {
		return nil, createLiErr
	}
	createEndRt, createEndErr := svc.CreateEndpointGroup(&aga.CreateEndpointGroupInput{
		EndpointGroupRegion: aws.String(Region),
		IdempotencyToken:    aws.String(IdempotencyToken),
		ListenerArn:         createLiRt.Listener.ListenerArn,
		HealthCheckPort:     aws.Int64(22),
		EndpointConfigurations: []*aga.EndpointConfiguration{
			{
				EndpointId: aws.String(InstanceId),
			},
		},
	})
	if createEndErr != nil {
		return nil, createEndErr
	}
	return &AgaInfo{
		Name:     createAccRt.Accelerator.Name,
		Status:   createAccRt.Accelerator.Status,
		Arn:      *createAccRt.Accelerator.AcceleratorArn,
		Ip:       createAccRt.Accelerator.IpSets,
		Protocol: createLiRt.Listener.Protocol,
		Port:     createEndRt.EndpointGroup.PortOverrides,
	}, nil
}

func (a *Aws) ListAga() ([]*aga.Accelerator, error) {
	svc := aga.New(a.Sess)
	rt, err := svc.ListAccelerators(&aga.ListAcceleratorsInput{})
	if err != nil {
		return nil, err
	}
	return rt.Accelerators, err
}

/* func (p *Aws) GetAgaInfo(AcceleratorArn string) (*AgaInfo, error) {
	svc := aga.New(p.Sess)
	accRt, accErr := svc.DescribeAccelerator(&aga.DescribeAcceleratorInput{AcceleratorArn: aws.String(AcceleratorArn)})
	if accErr != nil {
		return nil, accErr
	}
	liRt, liErr := svc.ListListeners(&aga.ListListenersInput{AcceleratorArn: accRt.Accelerator.AcceleratorArn})
	if liErr != nil {
		return nil, liErr
	}
	endRt, endErr := svc.ListEndpointGroups(&aga.ListEndpointGroupsInput{ListenerArn: liRt.Listeners[0].ListenerArn})
	if endErr != nil {
		return nil, endErr
	}
	return &AgaInfo{
		Username:   accRt.Accelerator.Username,
		Status: accRt.Accelerator.Status,
		Arn: AcceleratorArn + "_" +
			*liRt.Listeners[0].ListenerArn + "_" + *endRt.EndpointGroups[0].EndpointGroupArn,
		Ip:       accRt.Accelerator.IpSets,
		Protocol: liRt.Listeners[0].Protocol,
		Port:     endRt.EndpointGroups[0].PortOverrides,
	}, nil
}*/

func (a *Aws) DeleteAga(AcceleratorArn string) error {
	svc := aga.New(a.Sess)
	_, updateErr := svc.UpdateAccelerator(&aga.UpdateAcceleratorInput{
		AcceleratorArn: aws.String(AcceleratorArn),
		Enabled:        aws.Bool(false),
	})
	if updateErr != nil {
		return updateErr
	}
	_, deleteErr := svc.DeleteAccelerator(&aga.DeleteAcceleratorInput{AcceleratorArn: aws.String(AcceleratorArn)})
	if deleteErr != nil {
		return deleteErr
	}
	return nil
}

package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lightsail"
	log "github.com/sirupsen/logrus"
	"strconv"
)

type LsInfo struct {
	Name   *string
	Ip     *string
	IpName *string
	Type   *string
	Status *string
	Key    *string
}

func (p *Aws) GetBlueprintId() (*lightsail.GetBlueprintsOutput, error) {
	svc := lightsail.New(p.Sess)
	blueID, blueErr := svc.GetBlueprints(&lightsail.GetBlueprintsInput{})
	if blueErr != nil {
		return nil, blueErr
	}
	return blueID, nil
}

func (p *Aws) CreateLs(Name string, Zone string, BlueprintId string, BundleId string, Quantity string) (*LsInfo, error) {
	svc := lightsail.New(p.Sess)
	named := []*string{aws.String(Name)}
	quantity, err := strconv.Atoi(Quantity)
	if err == nil {
		for quantity > 1 {
			quantity -= 1
			named = append(named, aws.String(Name+"_"+strconv.Itoa(quantity)))
		}
	}

	lsRt, lsErr := svc.CreateInstances(&lightsail.CreateInstancesInput{
		AvailabilityZone: aws.String(Zone),
		BlueprintId:      aws.String(BlueprintId),
		BundleId:         aws.String(BundleId),
		InstanceNames:    named,
		UserData:         aws.String(" "),
	})
	if lsErr != nil {
		log.Warning(lsErr.Error() + "lsErr")
		return nil, lsErr
	}

	return &LsInfo{
		Name:   &Name,
		Status: lsRt.Operations[0].Status,
	}, nil
}

func (p *Aws) OffLightFirewall(Name string) (*LsInfo, error) {
	svc := lightsail.New(p.Sess)
	_, PortErr := svc.OpenInstancePublicPorts(&lightsail.OpenInstancePublicPortsInput{
		InstanceName: aws.String(Name),
		PortInfo: &lightsail.PortInfo{
			Protocol: aws.String("all"),
			FromPort: aws.Int64(0),
			ToPort:   aws.Int64(65535),
		},
	})
	if PortErr != nil {
		log.Warning(PortErr.Error() + "PortErr")
		return nil, PortErr
	}
	return nil, nil
}

func (p *Aws) GetLsInfo(Name string) (*LsInfo, error) {
	svc := lightsail.New(p.Sess)
	rt, err := svc.GetInstance(&lightsail.GetInstanceInput{InstanceName: aws.String(Name)})
	if err != nil {
		return nil, err
	}
	return &LsInfo{
		Name:   rt.Instance.Name,
		Ip:     rt.Instance.PublicIpAddress,
		IpName: rt.Instance.Tags[0].Value,
		Status: rt.Instance.State.Name,
	}, nil
}

func (p *Aws) ListLs() ([]*lightsail.Instance, error) {
	svc := lightsail.New(p.Sess)
	rt, err := svc.GetInstances(&lightsail.GetInstancesInput{})
	for rt.NextPageToken != nil {
		new, werr := svc.GetInstances(&lightsail.GetInstancesInput{
			PageToken: rt.NextPageToken,
		})
		if werr != nil {
			log.Warning(werr.Error() + "werr")
			return nil, nil
		}

		rt.Instances = append(rt.Instances, new.Instances...)
		rt.NextPageToken = new.NextPageToken

	}

	if err != nil {
		return nil, err
	}
	return rt.Instances, nil
}

func (p *Aws) ChangeLsIp(Name string) error {
	getRt, getErr := p.GetLsInfo(Name)
	if getErr != nil {
		return getErr
	}
	svc := lightsail.New(p.Sess)
	_, detErr := svc.DetachStaticIp(&lightsail.DetachStaticIpInput{StaticIpName: getRt.IpName})
	if detErr != nil {
		return detErr
	}
	_, relErr := svc.ReleaseStaticIp(&lightsail.ReleaseStaticIpInput{StaticIpName: getRt.IpName})
	if relErr != nil {
		return relErr
	}
	_, allErr := svc.AllocateStaticIp(&lightsail.AllocateStaticIpInput{StaticIpName: getRt.IpName})
	if allErr != nil {
		return allErr
	}
	_, attErr := svc.AttachStaticIp(&lightsail.AttachStaticIpInput{
		StaticIpName: getRt.IpName,
		InstanceName: aws.String(Name),
	})
	if attErr != nil {
		return attErr
	}
	return nil
}

func (p *Aws) StopLs(Name string) error {
	svc := lightsail.New(p.Sess)
	_, err := svc.StopInstance(&lightsail.StopInstanceInput{InstanceName: aws.String(Name)})
	if err != nil {
		return err
	}
	return nil
}

func (p *Aws) StartLs(Name string) error {
	svc := lightsail.New(p.Sess)
	_, err := svc.StartInstance(&lightsail.StartInstanceInput{InstanceName: aws.String(Name)})
	if err != nil {
		return err
	}
	return nil
}

func (p *Aws) RebootLs(Name string) error {
	svc := lightsail.New(p.Sess)
	_, err := svc.RebootInstance(&lightsail.RebootInstanceInput{InstanceName: aws.String(Name)})
	if err != nil {
		return err
	}
	return nil
}

func (p *Aws) DeleteLs(Name string) error {
	svc := lightsail.New(p.Sess)
	_, err := svc.DeleteInstance(&lightsail.DeleteInstanceInput{InstanceName: aws.String(Name)})
	if err != nil {
		return err
	}
	return nil
}

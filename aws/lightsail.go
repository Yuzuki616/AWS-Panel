package aws

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lightsail"
)

type LsInfo struct {
	Name       *string
	Ip         *string
	SourceName *string
	Type       *string
	Status     *string
	Key        *string
}

func (p *Aws) GetBlueprintId() (*lightsail.GetBlueprintsOutput, error) {
	svc := lightsail.New(p.Sess)
	blueID, blueErr := svc.GetBlueprints(&lightsail.GetBlueprintsInput{})
	if blueErr != nil {
		return nil, blueErr
	}
	return blueID, nil
}

func (p *Aws) GetRegions() ([]*lightsail.Region, error) {
	svc := lightsail.New(p.Sess)
	regions, regErr := svc.GetRegions(&lightsail.GetRegionsInput{
		IncludeAvailabilityZones: aws.Bool(true),
	})
	if regErr != nil {
		return nil, regErr
	}
	return regions.Regions, nil
}

func (p *Aws) CreateLs(Name string, AvailabilityZone string, BlueprintId string, BundleId string) (*LsInfo, error) {
	svc := lightsail.New(p.Sess)
	dateName := Name + time.Unix(time.Now().Unix(), 0).
		Format(AvailabilityZone+"2006_01_02_15_04_05_"+strconv.Itoa(rand.Intn(1000)))
	key, keyErr := svc.CreateKeyPair(&lightsail.CreateKeyPairInput{
		KeyPairName: aws.String(dateName)})
	if keyErr != nil {
		return nil, fmt.Errorf("create key error: %v", keyErr)
	}
	lsRt, lsErr := svc.CreateInstances(&lightsail.CreateInstancesInput{
		AvailabilityZone: aws.String(AvailabilityZone),
		BlueprintId:      aws.String(BlueprintId),
		BundleId:         aws.String(BundleId),
		InstanceNames:    []*string{aws.String(Name)},
		KeyPairName:      aws.String(dateName),
		Tags: []*lightsail.Tag{
			{
				Key:   aws.String("SIp"),
				Value: aws.String(dateName),
			},
		},
	})
	if lsErr != nil {
		return nil, fmt.Errorf("create ls error: %v", lsErr)
	}
	/*_, allErr := svc.AllocateStaticIp(&lightsail.AllocateStaticIpInput{
		StaticIpName: aws.String(dateName + "_ip")})
	if allErr != nil {
		return nil, fmt.Errorf("allocate ip error: %v", allErr)
	}
	_, attErr := svc.AttachStaticIp(&lightsail.AttachStaticIpInput{
		StaticIpName: aws.String(dateName + "_ip"),
		InstanceName: aws.String(Name),
	})
	if attErr != nil {
		return nil, fmt.Errorf("attach ip error: %v", attErr)
	}*/
	return &LsInfo{
		Name:   &Name,
		Status: lsRt.Operations[0].Status,
		Key:    key.PrivateKeyBase64,
	}, nil
}

func (p *Aws) OpenLsPorts(Name string) error {
	svc := lightsail.New(p.Sess)
	_, err := svc.OpenInstancePublicPorts(&lightsail.OpenInstancePublicPortsInput{
		InstanceName: aws.String(Name),
		PortInfo: &lightsail.PortInfo{
			Protocol: aws.String("all"),
			FromPort: aws.Int64(0),
			ToPort:   aws.Int64(65535),
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (p *Aws) GetLsInfo(Name string) (*LsInfo, error) {
	svc := lightsail.New(p.Sess)
	rt, err := svc.GetInstance(&lightsail.GetInstanceInput{
		InstanceName: aws.String(Name)})
	if err != nil {
		return nil, err
	}
	var tag *string
	if len(rt.Instance.Tags) > 0 {
		tag = rt.Instance.Tags[0].Value
	}
	return &LsInfo{
		Name:       rt.Instance.Name,
		Ip:         rt.Instance.PublicIpAddress,
		SourceName: tag,
		Status:     rt.Instance.State.Name,
	}, nil
}

func (p *Aws) ListLs() ([]*lightsail.Instance, error) {
	svc := lightsail.New(p.Sess)
	rt, err := svc.GetInstances(&lightsail.GetInstancesInput{})
	if err != nil {
		return nil, err
	}
	return rt.Instances, nil
}

func (p *Aws) ChangeLsIp(Name, Zone string) error {
	svc := lightsail.New(p.Sess)
	getRt, getErr := p.GetLsInfo(Name)
	if getErr != nil {
		return getErr
	}
	if getRt.SourceName == nil {
		*getRt.SourceName = Name + time.Unix(time.Now().Unix(), 0).
			Format(Zone+"2006_01_02_15_04_05_"+strconv.Itoa(rand.Intn(1000)))
		*getRt.SourceName = *getRt.SourceName + "_ip"
		_, detErr := svc.DetachStaticIp(&lightsail.DetachStaticIpInput{
			StaticIpName: getRt.SourceName})
		if detErr != nil {
			return fmt.Errorf("detach ip error: %v", detErr)
		}
		_, relErr := svc.ReleaseStaticIp(&lightsail.ReleaseStaticIpInput{
			StaticIpName: getRt.SourceName})
		if relErr != nil {
			return fmt.Errorf("release ip error: %v", relErr)
		}
	}
	*getRt.SourceName = *getRt.SourceName + "_ip"
	_, allErr := svc.AllocateStaticIp(&lightsail.AllocateStaticIpInput{
		StaticIpName: getRt.SourceName})
	if allErr != nil {
		return fmt.Errorf("allocate ip error: %v", allErr)
	}
	_, attErr := svc.AttachStaticIp(&lightsail.AttachStaticIpInput{
		StaticIpName: getRt.SourceName,
		InstanceName: aws.String(Name),
	})
	if attErr != nil {
		return fmt.Errorf("attach ip error: %v", attErr)
	}
	return nil
}

func (p *Aws) StopLs(Name string) error {
	svc := lightsail.New(p.Sess)
	_, err := svc.StopInstance(&lightsail.StopInstanceInput{
		InstanceName: aws.String(Name)})
	if err != nil {
		return err
	}
	return nil
}

func (p *Aws) StartLs(Name string) error {
	svc := lightsail.New(p.Sess)
	_, err := svc.StartInstance(&lightsail.StartInstanceInput{
		InstanceName: aws.String(Name)})
	if err != nil {
		return err
	}
	return nil
}

func (p *Aws) RebootLs(Name string) error {
	svc := lightsail.New(p.Sess)
	_, err := svc.RebootInstance(&lightsail.RebootInstanceInput{
		InstanceName: aws.String(Name)})
	if err != nil {
		return err
	}
	return nil
}

func (p *Aws) DeleteLs(Name string, SourceName string) error {
	svc := lightsail.New(p.Sess)
	if SourceName != "" {
		ipName := SourceName + "_ip"
		_, detErr := svc.DetachStaticIp(&lightsail.DetachStaticIpInput{
			StaticIpName: aws.String(ipName)})
		if detErr != nil {
			return fmt.Errorf("detach ip error: %v", detErr)
		}
		_, relErr := svc.ReleaseStaticIp(&lightsail.ReleaseStaticIpInput{
			StaticIpName: aws.String(ipName)})
		if relErr != nil {
			return fmt.Errorf("release ip error: %v", relErr)
		}
		_, delErr := svc.DeleteKeyPair(&lightsail.DeleteKeyPairInput{
			KeyPairName: aws.String(SourceName)})
		if delErr != nil {
			return fmt.Errorf("delete key error: %v", delErr)
		}
	}
	_, delInstance := svc.DeleteInstance(&lightsail.DeleteInstanceInput{
		InstanceName: aws.String(Name)})
	if delInstance != nil {
		return fmt.Errorf("delete instance error: %v", delInstance)
	}
	return nil
}

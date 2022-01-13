package aws

import "github.com/aws/aws-sdk-go/service/ec2"

func CheckNameNil(v []*ec2.Tag) *string {
	if v == nil {
		return nil
	}
	return v[0].Value
}

package aws

import (
	"crypto/tls"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"net/http"
	"net/url"
)

type Aws struct {
	Sess *session.Session
}

func New(Region string, Id string, Secret string, Proxy string) (*Aws, error) {
	config := &aws.Config{
		Region:      aws.String(Region),
		Credentials: credentials.NewStaticCredentials(Id, Secret, ""),
	}
	if Proxy != "" {
		config.HTTPClient = &http.Client{
			Transport: &http.Transport{
				Proxy: func(*http.Request) (*url.URL, error) {
					return url.Parse(Proxy)
				},
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		}
	}
	sess, err := session.NewSession(config)
	if err != nil {
		return nil, err
	}
	c := &Aws{
		Sess: sess,
	}
	return c, nil
}

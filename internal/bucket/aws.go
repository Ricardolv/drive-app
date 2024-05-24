package bucket

import (
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

type AwsConfig struct {
	Config         *aws.Config
	BucketDownload string
	BucketUpload   string
}

func NewAwsSession(cfg AwsConfig) *awsSession {
	c := session.New(cfg.Config)

	return &awsSession{
		sess:           c,
		bucketDownload: cfg.BucketDownload,
		bucketUpload:   cfg.BucketUpload,
	}
}

type awsSession struct {
	sess           *session.Session
	bucketDownload string
	bucketUpload   string
}

func (a *awsSession) Download(src, dst string) (*os.File, error) {
	return nil, nil
}

func (a *awsSession) Upload(file io.Reader, key string) error {
	return nil
}

func (a *awsSession) Delete(key string) error {
	return nil
}

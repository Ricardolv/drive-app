package bucket

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

const (
	AwsProvider BucketType = iota
)

type BucketType int

func New(config any, bucketType BucketType) (bucket *Bucket, err error) {

	rt := reflect.TypeOf(config)

	switch bucketType {
	case AwsProvider:

	default:
		return nil, fmt.Errorf("type not implemented")
	}

	return
}

type BucketInterface interface {
	Upload(io.Reader, string) error
	Download(string, string) (*os.File, error)
	Delete(string) error
}

type Bucket struct {
	p BucketInterface
}

func (b *Bucket) Upload(file io.Reader, key string) error {
	return b.p.Upload(file, key)
}

func (b *Bucket) Download(src, dst string) (*os.File, error) {
	return b.p.Download(src, dst)
}

func (b *Bucket) Delete(key string) error {
	return b.p.Delete(key)
}

package s3

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
)

type Aws3 struct {
	service   *s3.S3
	bucket    *string
	accessKey string
	secretKey string
	domain    string
}

func NewS3(bucket *string, ak, sk, url, region string) *Aws3 {
	ossSession := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Credentials:      credentials.NewStaticCredentials(ak, sk, "token"),
			Endpoint:         aws.String(url),
			Region:           aws.String(region),
			S3ForcePathStyle: aws.Bool(true),
		},
	}))

	svc := s3.New(ossSession)
	return &Aws3{
		bucket:    bucket,
		service:   svc,
		accessKey: ak,
		secretKey: sk,
		domain:    url,
	}
}

func (a *Aws3) Upload(key string, src io.ReadSeeker) (string, error) {
	resp, err := a.service.GetObject(&s3.GetObjectInput{
		Bucket: a.bucket,
		Key:    &key,
	})
	if err != nil && err.Error()[:9] != "NoSuchKey" {
		return "", err
	}
	if resp.Body != nil {
		return "", errors.New("Duplicate filename!")
	}
	if _, err := a.service.PutObject(&s3.PutObjectInput{
		Body:   src,
		Bucket: a.bucket,
		Key:    &key,
	}); err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", key), nil
}

func (a *Aws3) Download(key string) (io.ReadCloser, error) {
	if resp, err := a.service.GetObject(&s3.GetObjectInput{
		Bucket: a.bucket,
		Key:    &key,
	}); err != nil {
		return nil, err
	} else {
		return resp.Body, err
	}
}

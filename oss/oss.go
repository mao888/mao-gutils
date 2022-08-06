package oss

import (
	"io"
	"os"
	"path"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func CreateObject(sess *session.Session, bucket *string, key *string, data io.ReadSeeker) error {
	svc := s3.New(sess)

	if _, err := svc.PutObject(&s3.PutObjectInput{
		Body:   data,
		Bucket: bucket,
		Key:    key,
	}); err != nil {
		return err
	}
	return nil
}

func GetObjects(sess *session.Session, bucket *string) (*s3.ListObjectsV2Output, error) {
	svc := s3.New(sess)

	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: bucket})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func DownloadObject(sess *session.Session, dirPath, filename string, bucket, key *string) error {
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return err
	}
	file, err := os.Create(path.Join(dirPath, filename))
	if err != nil {
		return err
	}

	defer file.Close()

	downloader := s3manager.NewDownloader(sess)

	if _, err := downloader.Download(file, &s3.GetObjectInput{
		Bucket: bucket,
		Key:    key,
	}); err != nil {
		return err
	}
	return nil
}

func GetObject(sess *session.Session, bucket, key *string) (*s3.GetObjectOutput, error) {
	svc := s3.New(sess)

	resp, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: bucket,
		Key:    key,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

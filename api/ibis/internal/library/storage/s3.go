package storage

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/takeuchima0/async_serverless_application_demo/api/ibis/internal/configuration"
)

var _ StorageService = (*S3Wrapper)(nil)

type StorageService interface {
	GetUploadPresignedURL(context.Context, string, time.Duration) (string, error)
	GetDownloadPresignedURL(context.Context, string, time.Duration) (string, error)
}

type S3Wrapper struct {
	bucket string
	client *s3.Client
	pre    *s3.PresignClient
}

func NewStorageService(bucketName string) *S3Wrapper {
	s3Client := s3.NewFromConfig(configuration.Get().AWSConfig)
	return &S3Wrapper{
		bucket: bucketName,
		client: s3Client,
		pre:    s3.NewPresignClient(s3Client),
	}
}

package storage

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (s *S3Wrapper) GetUploadPresignedURL(ctx context.Context, filePath string, expire time.Duration) (string, error) {
	input := &s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(filePath),
	}
	output, err := s.pre.PresignPutObject(ctx, input, func(po *s3.PresignOptions) {
		po.Expires = expire
	})
	if err != nil {
		return "", err
	}
	return output.URL, nil
}

package configuration

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AWSConfig   aws.Config
	Env         string `envconfig:"ENV" default:"dev"`
	ServiceName string `envconfig:"SERVICE_NAME" default:"image-maker-batch"`

	S3 struct {
		BucketName struct {
			ThumbnailImage string `envconfig:"S3_THUMBNAIL_IMAGE_BUCKET_NAME"`
		}
	}
}

var globalConfig Config

func Get() Config { return globalConfig }

func Load(ctx context.Context) error {
	envconfig.MustProcess("", &globalConfig)

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := loadAWSConfig(ctx, globalConfig.Env); err != nil {
		return fmt.Errorf("failed to load AWS configuration: %w", err)
	}

	return nil
}

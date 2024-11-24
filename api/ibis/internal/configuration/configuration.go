package configuration

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AWSConfig aws.Config
	API       struct {
		Env         string `envconfig:"API_ENV" default:"dev"`
		Port        string `envconfig:"API_PORT" default:"8080"`
		ServiceName string `envconfig:"API_SERVICE_NAME" default:"ibis-api"`
	}
	S3 struct {
		BucketName struct {
			OriginalImage  string `envconfig:"S3_ORIGINAL_IMAGE_BUCKET_NAME"`
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

	if err := loadAWSConfig(ctx, globalConfig.API.Env); err != nil {
		return fmt.Errorf("failed to load AWS configuration: %w", err)
	}

	return nil
}

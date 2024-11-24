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
		ServiceName string `envconfig:"API_SERVICE_NAME" default:"notification-batch"`
	}
	SQS struct {
		DlqURL string `envconfig:"SQS_DLQ_URL"`
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
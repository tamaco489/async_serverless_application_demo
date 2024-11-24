package configuration

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
)

const awsRegion = "ap-northeast-1"

func loadAWSConfig(ctx context.Context, env string) error {

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(awsRegion))
	if err != nil {
		return err
	}

	switch env {
	case "dev":
		//

	default:
		//
	}

	globalConfig.AWSConfig = cfg

	return nil
}

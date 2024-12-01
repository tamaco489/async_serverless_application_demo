package repository

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoDBWrapper struct {
	client    *dynamodb.Client
	tableName string
}

func NewDynamoDBWrapper(cfg aws.Config, tableName string) *DynamoDBWrapper {
	cfg.Credentials = aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider("dummy", "dummy", ""))

	client := dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
		o.BaseEndpoint = aws.String("http://dynamodb-local:8000")
	})

	return &DynamoDBWrapper{
		client:    client,
		tableName: tableName,
	}
}

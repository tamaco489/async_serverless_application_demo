package repository

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoDBRepository struct {
	client *dynamodb.Client
}

func NewDynamoDBRepository(cfg aws.Config, env string) *DynamoDBRepository {

	var client *dynamodb.Client

	switch env {
	case "dev":
		cfg.Credentials = aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider("dummy", "dummy", ""))
		client = dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
			o.BaseEndpoint = aws.String("http://dynamodb-local:8000")
		})

	case "stg":
		client = dynamodb.NewFromConfig(cfg)
	}

	return &DynamoDBRepository{
		client: client,
	}
}

// Client: NewDynamoDBRepositoryによりDynamoDBRepositoryを生成し、それを満たした状態でこのメソッドを呼び出す。
func (r *DynamoDBRepository) Client() *dynamodb.Client {
	return r.client
}

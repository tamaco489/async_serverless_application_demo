package dynamo_db

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoDBWrapper struct {
	client    *dynamodb.Client
	tableName string
}

type DynamoDBService interface {
	CreateUser(ctx context.Context, user map[string]interface{}) (*dynamodb.PutItemOutput, error)
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

func (w *DynamoDBWrapper) CreateUser(ctx context.Context, user map[string]interface{}) (*dynamodb.PutItemOutput, error) {
	item, err := attributevalue.MarshalMap(user)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal item: %v", err)
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(w.tableName),
		Item:      item,
	}

	output, err := w.client.PutItem(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to put item into DynamoDB: %v", err)
	}

	return output, nil
}

var _ DynamoDBService = (*DynamoDBWrapper)(nil)

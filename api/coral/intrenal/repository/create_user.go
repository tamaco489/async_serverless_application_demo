package repository

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoDBService interface {
	CreateUser(ctx context.Context, user map[string]interface{}) (*dynamodb.PutItemOutput, error)
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

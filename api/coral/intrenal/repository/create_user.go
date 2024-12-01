package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/tamaco489/async_serverless_application_demo/api/coral/intrenal/model"
)

type DynamoDBService interface {
	CreateUser(ctx context.Context, user model.User) (*dynamodb.PutItemOutput, error)
}

func (w *DynamoDBRepository) CreateUser(ctx context.Context, user model.User) (*dynamodb.PutItemOutput, error) {

	item, err := user.DynamoAttributeMapFromUser()
	if err != nil {
		return nil, fmt.Errorf("failed to convert user to attribute map: %v", err)
	}

	output, err := w.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(w.tableName),
		Item:      item,
	})
	if err != nil {
		var ae awserr.Error
		if ok := errors.As(err, &ae); ok {
			return nil, fmt.Errorf("failed to put item into DynamoDB: %v", ae.Error())
		}
		return nil, fmt.Errorf("failed to put item into DynamoDB: %v", err)
	}

	return output, nil
}

var _ DynamoDBService = (*DynamoDBRepository)(nil)

package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/tamaco489/async_serverless_application_demo/api/coral/intrenal/model"
)

type DynamoDBService interface {
	CreateUser(ctx context.Context, tableName string, user *model.User) error
	GetMeUser(ctx context.Context, tableName string, uid string) (*model.User, error)
}

var _ DynamoDBService = (*DynamoDBRepository)(nil)

// CreateUser: ユーザを新規で作成します。
func (w *DynamoDBRepository) CreateUser(ctx context.Context, tableName string, user *model.User) error {

	item, err := user.DynamoAttributeMapFromUser()
	if err != nil {
		return fmt.Errorf("failed to convert user to attribute map: %v", err)
	}

	_, err = w.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      item,
	})
	if err != nil {
		var ae awserr.Error
		if ok := errors.As(err, &ae); ok {
			return fmt.Errorf("failed to put item into dynamodb: %v", ae.Error())
		}
		return fmt.Errorf("failed to put item into dynamodb: %v", err)
	}

	return nil
}

// GetMeUser: 自身のユーザ情報を取得します。
func (w *DynamoDBRepository) GetMeUser(ctx context.Context, tableName string, uid string) (*model.User, error) {

	result, err := w.client.Query(ctx, &dynamodb.QueryInput{
		TableName:              aws.String(tableName),
		KeyConditionExpression: aws.String("user_id = :user_id"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":user_id": &types.AttributeValueMemberS{Value: uid},
		},
	})
	if err != nil {
		var ae awserr.Error
		if ok := errors.As(err, &ae); ok {
			return nil, fmt.Errorf("failed to query dynamodb: %v", ae.Error())
		}
		return nil, fmt.Errorf("failed to query dynamodb: %v", err)
	}

	// クエリ結果が空の場合は、ユーザが見つからなかったことを返す
	if len(result.Items) == 0 {
		return nil, fmt.Errorf("user with uid %s not found", uid)
	}

	// ユニークなユーザのため基本的には複数のレコードが返却されることはない
	if len(result.Items) > 1 {
		return nil, fmt.Errorf("many user with uid %s", uid)
	}

	// DynamoDBのItemをUser構造体に変換
	item := result.Items[0] // 1件しか返されないので、最初の1件を取得
	u := model.NewUser()

	// DynamoDBの属性マップをUser構造体に変換
	user := u.DynamoAttributeMapToUser(item)

	return user, nil
}

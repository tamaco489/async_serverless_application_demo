package model

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type User struct {
	UserID     string `json:"user_id"`
	Email      string `json:"email"`
	Birthday   string `json:"birthday"`
	EkycStatus string `json:"ekyc_status"`
	InviteCode string `json:"invite_code"`
	IsAdmin    bool   `json:"is_admin"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

// NewUser: Userモデルの初期化
func NewUser() *User {
	return &User{
		UserID:     "",
		Email:      "",
		Birthday:   "",
		EkycStatus: "",
		InviteCode: "",
		IsAdmin:    false,
		CreatedAt:  "",
		UpdatedAt:  "",
	}
}

// dynamoAttributeMapFromUser: User型をDynamoDB用の属性マップに変換
func (u User) DynamoAttributeMapFromUser() (map[string]types.AttributeValue, error) {

	item := make(map[string]types.AttributeValue)

	// UserID
	item["user_id"] = &types.AttributeValueMemberS{Value: u.UserID}

	// Email
	item["email"] = &types.AttributeValueMemberS{Value: u.Email}

	// Birthday
	item["birthday"] = &types.AttributeValueMemberS{Value: u.Birthday}

	// EkycStatus
	item["ekyc_status"] = &types.AttributeValueMemberS{Value: u.EkycStatus}

	// InviteCode
	item["invite_code"] = &types.AttributeValueMemberS{Value: u.InviteCode}

	// IsAdmin
	item["is_admin"] = &types.AttributeValueMemberBOOL{Value: u.IsAdmin}

	// CreatedAt
	item["created_at"] = &types.AttributeValueMemberS{Value: u.CreatedAt}

	// UpdatedAt
	item["updated_at"] = &types.AttributeValueMemberS{Value: u.UpdatedAt}

	return item, nil
}

// DynamoAttributeMapToUser: DynamoDBの属性マップをUser構造体に変換
func (u User) DynamoAttributeMapToUser(item map[string]types.AttributeValue) *User {
	// 各属性をマッピング。対応する型にキャストして設定
	setUserAttribute(item, "user_id", &u.UserID)
	setUserAttribute(item, "email", &u.Email)
	setUserAttribute(item, "birthday", &u.Birthday)
	setUserAttribute(item, "ekyc_status", &u.EkycStatus)
	setUserAttribute(item, "invite_code", &u.InviteCode)
	setUserAttribute(item, "is_admin", &u.IsAdmin)
	setUserAttribute(item, "created_at", &u.CreatedAt)
	setUserAttribute(item, "updated_at", &u.UpdatedAt)

	return &u
}

// setUserAttribute: DynamoDBのアイテムからUser構造体のフィールドに値を設定するユーティリティ関数
func setUserAttribute(item map[string]types.AttributeValue, key string, target interface{}) {
	if value, exists := item[key]; exists {
		switch v := value.(type) {
		case *types.AttributeValueMemberS:
			if strTarget, ok := target.(*string); ok {
				*strTarget = v.Value
			}
		case *types.AttributeValueMemberBOOL:
			if boolTarget, ok := target.(*bool); ok {
				*boolTarget = v.Value
			}
		default:
			log.Printf("unexpected type for key: %s, value type: %T", key, v)
		}
	}
}

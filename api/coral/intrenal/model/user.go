package model

import (
	"time"

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
func NewUser(userID, email, birthday, ekycStatus, inviteCode string, isAdmin bool, now time.Time) *User {

	currentTime := now.Format(time.RFC3339) // 現在の時刻をISO 8601形式で取得

	return &User{
		UserID:     userID,
		Email:      email,
		Birthday:   birthday,
		EkycStatus: ekycStatus,
		InviteCode: inviteCode,
		IsAdmin:    isAdmin,
		CreatedAt:  currentTime,
		UpdatedAt:  currentTime,
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

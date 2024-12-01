package model

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type User struct {
	UserID     string `json:"UserID"`
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
func (user User) DynamoAttributeMapFromUser() (map[string]types.AttributeValue, error) {
	item := make(map[string]types.AttributeValue)

	// UserID
	item["UserID"] = &types.AttributeValueMemberS{Value: user.UserID}

	// Email
	item["Email"] = &types.AttributeValueMemberS{Value: user.Email}

	// Birthday
	item["Birthday"] = &types.AttributeValueMemberS{Value: user.Birthday}

	// EkycStatus
	item["EkycStatus"] = &types.AttributeValueMemberS{Value: user.EkycStatus}

	// InviteCode
	item["InviteCode"] = &types.AttributeValueMemberS{Value: user.InviteCode}

	// IsAdmin
	item["IsAdmin"] = &types.AttributeValueMemberBOOL{Value: user.IsAdmin}

	// CreatedAt
	item["CreatedAt"] = &types.AttributeValueMemberS{Value: user.CreatedAt}

	// UpdatedAt
	item["UpdatedAt"] = &types.AttributeValueMemberS{Value: user.UpdatedAt}

	return item, nil
}

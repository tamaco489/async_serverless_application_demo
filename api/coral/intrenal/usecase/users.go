package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/tamaco489/async_serverless_application_demo/api/coral/intrenal/model"
	"github.com/tamaco489/async_serverless_application_demo/api/coral/intrenal/repository"
)

type userUseCase struct {
	dynamoRepo *repository.DynamoDBRepository
	tableName  string
}

type UserUseCase interface {
	CreateUser(ctx context.Context, user map[string]interface{}) (*model.User, error)
	GetMeUser(ctx context.Context, uid string) (*model.User, error)
	GetUserByID(ctx context.Context, uid string) (*model.User, error)
}

var _ UserUseCase = (*userUseCase)(nil)

func NewUserUseCase(dynamoRepo *repository.DynamoDBRepository) *userUseCase {
	const tableName = "users"
	return &userUseCase{
		dynamoRepo: dynamoRepo,
		tableName:  tableName,
	}
}

// CreateUser: ユーザを新規で作成します。
func (uc *userUseCase) CreateUser(ctx context.Context, user map[string]interface{}) (*model.User, error) {

	u, err := uc.convertToUserModel(user)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to user model: %v", err)
	}

	// 現在の時刻をISO 8601形式で取得
	currentTime := time.Now().Format(time.RFC3339)
	u.CreatedAt = currentTime
	u.UpdatedAt = currentTime

	if err = uc.dynamoRepo.CreateUser(ctx, uc.tableName, u); err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	return &model.User{
		UserID:     u.UserID,
		Email:      u.Email,
		Birthday:   u.Birthday,
		EkycStatus: u.EkycStatus,
		InviteCode: u.InviteCode,
		IsAdmin:    u.IsAdmin,
		CreatedAt:  currentTime,
		UpdatedAt:  currentTime,
	}, nil
}

// GetMeUser: 自身のユーザ情報を取得します。
func (uc *userUseCase) GetMeUser(ctx context.Context, uid string) (*model.User, error) {
	return uc.dynamoRepo.GetMeUser(ctx, uc.tableName, uid)
}

func (uc *userUseCase) GetUserByID(ctx context.Context, uid string) (*model.User, error) {
	return &model.User{
		UserID: uid,
	}, nil
}

// convertToUserModel: リクエストボディに指定された内容をUserモデルに変換
func (uc *userUseCase) convertToUserModel(user map[string]interface{}) (*model.User, error) {

	// 必要なフィールドが存在しない場合や型が不一致の場合にエラーを返す
	email, ok := user["email"].(string)
	if !ok {
		return nil, fmt.Errorf("email is required and should be a string")
	}

	birthday, ok := user["birthday"].(string)
	if !ok {
		return nil, fmt.Errorf("birthday is required and should be a string")
	}

	ekycStatus, ok := user["ekyc_status"].(string)
	if !ok {
		return nil, fmt.Errorf("ekyc_status is required and should be a string")
	}

	inviteCode, ok := user["invite_code"].(string)
	if !ok {
		return nil, fmt.Errorf("invite_code is required and should be a string")
	}

	isAdmin, ok := user["is_admin"].(bool)
	if !ok {
		return nil, fmt.Errorf("is_admin is required and should be an boolean")
	}

	now := time.Now()
	uuid, err := uuid.NewV7AtTime(now)
	if err != nil {
		return nil, fmt.Errorf("failed generate uuid error: %v", err)
	}

	return &model.User{
		UserID:     uuid.String(),
		Email:      email,
		Birthday:   birthday,
		EkycStatus: ekycStatus,
		InviteCode: inviteCode,
		IsAdmin:    isAdmin,
	}, nil
}

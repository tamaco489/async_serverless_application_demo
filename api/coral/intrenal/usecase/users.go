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
	CreateUser(ctx context.Context, user map[string]interface{}) (model.User, error)
}

func NewUserUseCase(dynamoRepo *repository.DynamoDBRepository) *userUseCase {
	const tableName = "users"
	return &userUseCase{
		dynamoRepo: dynamoRepo,
		tableName:  tableName,
	}
}

func (uc *userUseCase) CreateUser(ctx context.Context, user map[string]interface{}) (model.User, error) {

	userModel, err := uc.convertToUserModel(user)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to convert to user model: %v", err)
	}

	if _, err = uc.dynamoRepo.CreateUser(ctx, uc.tableName, userModel); err != nil {
		return model.User{}, fmt.Errorf("failed to create user: %v", err)
	}

	return userModel, nil
}

var _ UserUseCase = (*userUseCase)(nil)

// convertToUserModel: リクエストボディに指定された内容をUserモデルに変換
func (uc *userUseCase) convertToUserModel(user map[string]interface{}) (model.User, error) {

	// 必要なフィールドが存在しない場合や型が不一致の場合にエラーを返す
	email, ok := user["email"].(string)
	if !ok {
		return model.User{}, fmt.Errorf("email is required and should be a string")
	}

	birthday, ok := user["birthday"].(string)
	if !ok {
		return model.User{}, fmt.Errorf("birthday is required and should be a string")
	}

	ekycStatus, ok := user["ekyc_status"].(string)
	if !ok {
		return model.User{}, fmt.Errorf("ekyc_status is required and should be a string")
	}

	inviteCode, ok := user["invite_code"].(string)
	if !ok {
		return model.User{}, fmt.Errorf("invite_code is required and should be a string")
	}

	isAdmin, ok := user["is_admin"].(bool)
	if !ok {
		return model.User{}, fmt.Errorf("is_admin is required and should be an boolean")
	}

	now := time.Now()
	uuid, err := uuid.NewV7AtTime(now)
	if err != nil {
		return model.User{}, fmt.Errorf("failed generate uuid error: %v", err)
	}

	newUser := model.NewUser(
		uuid.String(),
		email,
		birthday,
		ekycStatus,
		inviteCode,
		isAdmin,
		now,
	)

	return *newUser, nil
}

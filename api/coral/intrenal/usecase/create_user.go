package usecase

import (
	"context"
	"fmt"

	"github.com/tamaco489/async_serverless_application_demo/api/coral/intrenal/model"
	"github.com/tamaco489/async_serverless_application_demo/api/coral/intrenal/repository"
)

type userUseCase struct {
	dynamoRepo *repository.DynamoDBRepository
}

type UserUseCase interface {
	CreateUser(ctx context.Context, user map[string]interface{}) error
}

func NewUserUseCase(dynamoRepo *repository.DynamoDBRepository) *userUseCase {
	return &userUseCase{
		dynamoRepo: dynamoRepo,
	}
}

func (uc *userUseCase) CreateUser(ctx context.Context, user map[string]interface{}) error {

	userModel, err := uc.convertToUserModel(user)
	if err != nil {
		return fmt.Errorf("failed to convert to user model: %v", err)
	}

	if _, err = uc.dynamoRepo.CreateUser(ctx, userModel); err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	return nil
}

var _ UserUseCase = (*userUseCase)(nil)

// convertToUserModel: リクエストボディに指定された内容をUserモデルに変換
func (uc *userUseCase) convertToUserModel(user map[string]interface{}) (model.User, error) {

	userModel := model.User{}

	// 必要なフィールドが存在しない場合や型が不一致の場合にエラーを返す
	if userID, ok := user["UserID"].(string); ok {
		userModel.UserID = userID
	} else {
		return userModel, fmt.Errorf("userID is required and should be a string")
	}

	if email, ok := user["email"].(string); ok {
		userModel.Email = email
	} else {
		return userModel, fmt.Errorf("email is required and should be a string")
	}

	if birthday, ok := user["birthday"].(string); ok {
		userModel.Birthday = birthday
	} else {
		return userModel, fmt.Errorf("birthday is required and should be a string")
	}

	if ekycStatus, ok := user["ekyc_status"].(string); ok {
		userModel.EkycStatus = ekycStatus
	} else {
		return userModel, fmt.Errorf("ekyc_status is required and should be a string")
	}

	if inviteCode, ok := user["invite_code"].(string); ok {
		userModel.InviteCode = inviteCode
	} else {
		return userModel, fmt.Errorf("invite_code is required and should be a string")
	}

	if isAdmin, ok := user["is_admin"].(float64); ok { // DynamoDBでは数値はfloat64として取得されることが多い
		userModel.IsAdmin = int(isAdmin)
	} else {
		return userModel, fmt.Errorf("is_admin is required and should be an integer")
	}

	return userModel, nil
}

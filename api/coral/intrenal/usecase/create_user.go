package usecase

import (
	"context"
	"fmt"

	"github.com/tamaco489/async_serverless_application_demo/api/coral/intrenal/repository"
)

type userUseCase struct {
	dynamoRepo *repository.DynamoDBWrapper
}

type UserUseCase interface {
	CreateUser(ctx context.Context, user map[string]interface{}) error
}

func NewUserUseCase(dynamoRepo *repository.DynamoDBWrapper) *userUseCase {
	return &userUseCase{
		dynamoRepo: dynamoRepo,
	}
}

func (uc *userUseCase) CreateUser(ctx context.Context, user map[string]interface{}) error {
	if _, err := uc.dynamoRepo.CreateUser(ctx, user); err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}
	return nil
}

var _ UserUseCase = (*userUseCase)(nil)

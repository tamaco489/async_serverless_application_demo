package handler

import (
	"context"

	"github.com/tamaco489/async_serverless_application_demo/api/coral/intrenal/configuration"
	"github.com/tamaco489/async_serverless_application_demo/api/coral/intrenal/repository"
	"github.com/tamaco489/async_serverless_application_demo/api/coral/intrenal/usecase"
)

type CoralHandler struct {
	config      configuration.Config
	userUseCase usecase.UserUseCase
}

func NewHandler(ctx context.Context) (*CoralHandler, error) {
	if err := configuration.Load(ctx); err != nil {
		return nil, err
	}

	cnf := configuration.Get()
	dynamoDBRepo := repository.NewDynamoDBRepository(cnf.AWSConfig)
	userUseCase := usecase.NewUserUseCase(dynamoDBRepo)

	return &CoralHandler{
		config:      cnf,
		userUseCase: userUseCase,
	}, nil
}

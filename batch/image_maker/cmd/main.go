package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tamaco489/async_serverless_application_demo/batch/image_maker/internal/handler"
)

func main() {
	ctx := context.Background()
	immHandler, err := handler.NewImageMakerHandler(ctx)
	if err != nil {
		panic(err)
	}
	lambda.Start(immHandler.Handler)
}

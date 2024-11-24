package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/takeuchima0/async_serverless_application_demo/api/coral/intrenal/handler"
)

func main() {
	ctx := context.Background()
	coralHandler, err := handler.NewHandler(ctx)
	if err != nil {
		panic(err)
	}
	lambda.Start(coralHandler.Handler)
}

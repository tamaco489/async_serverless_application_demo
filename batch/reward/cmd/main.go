package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler() error {
	log.Println("Starting Reward Batch Server...")

	return nil
}

func main() {
	lambda.Start(handler)
}
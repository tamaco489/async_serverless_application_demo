package sqs_client

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/tamaco489/async_serverless_application_demo/api/nautilus/models"
)

// SQSClient struct holds the AWS SQS client
type SQSClient struct {
	Client *sqs.Client
}

// NewSQSClient initializes a new SQS client with the given AWS config
func NewSQSClient(ctx context.Context, awsConfig aws.Config) (*SQSClient, error) {
	client := sqs.NewFromConfig(awsConfig)

	// Execute ListQueues as a client initialization test
	_, err := client.ListQueues(ctx, &sqs.ListQueuesInput{})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize SQS client:: %w", err)
	}

	return &SQSClient{Client: client}, nil
}

// SendPurchaseMessage sends a purchase message to the specified SQS queue
func (c *SQSClient) SendPurchaseMessage(ctx context.Context, queueURL string, msg models.PurchaseQueueMessage) error {
	// 構造体を JSON に変換
	msgBody, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal message to JSON: %w", err)
	}

	// SQS にメッセージを送信
	_, err = c.Client.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:    aws.String(queueURL),        // 送信するSQSのURL
		MessageBody: aws.String(string(msgBody)), // メッセージの内容
	})
	if err != nil {
		return fmt.Errorf("failed to send message to SQS: %w", err)
	}

	return nil
}

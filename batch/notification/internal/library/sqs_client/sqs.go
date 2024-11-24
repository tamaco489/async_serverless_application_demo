package sqs_client

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/takeuchima0/async_serverless_application_demo/batch/notification/internal/models"
)

type SQSClient struct {
	Client *sqs.Client
}

func NewSQSClient(ctx context.Context, awsConfig aws.Config) (*SQSClient, error) {
	client := sqs.NewFromConfig(awsConfig)
	return &SQSClient{Client: client}, nil
}

func (c *SQSClient) SendDLQMessage(ctx context.Context, queueURL string, msg models.DLQMessage) error {
	msgBody, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal message to JSON: %w", err)
	}

	slog.InfoContext(ctx, "send DLQ URL", slog.Any("url", string(queueURL)))
	slog.InfoContext(ctx, "send DLQ message", slog.Any("body", string(msgBody)))

	_, err = c.Client.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:    aws.String(queueURL),        // 送信するSQSのURL
		MessageBody: aws.String(string(msgBody)), // メッセージの内容
	})
	if err != nil {
		return fmt.Errorf("failed to send message to SQS: %w", err)
	}

	return nil
}

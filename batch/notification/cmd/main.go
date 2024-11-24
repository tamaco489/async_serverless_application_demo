package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/takeuchima0/async_serverless_application_demo/batch/notification/internal/configuration"
	"github.com/takeuchima0/async_serverless_application_demo/batch/notification/internal/library/sqs_client"
	"github.com/takeuchima0/async_serverless_application_demo/batch/notification/internal/models"
)

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	slog.InfoContext(ctx, "start handler processing SQS event", slog.Any("sqsEvent", sqsEvent))

	if err := configuration.Load(ctx); err != nil {
		return err
	}

	for _, e := range sqsEvent.Records {
		var purchaseMessage models.PurchaseQueueMessage
		if err := json.Unmarshal([]byte(e.Body), &purchaseMessage); err != nil {
			slog.ErrorContext(ctx, "failed to unmarshal SQS message", slog.String("messageBody", e.Body), slog.Any("error", err))
		}

		// SQSから送信されたメッセージキューの購入完了ステータスに応じで処理を分岐
		switch purchaseMessage.Status {
		case models.PurchaseStatusCompleted:
			slog.InfoContext(ctx, "Purchase status is completed", slog.String("purchase status", string(purchaseMessage.Status)))

		case models.PurchaseStatusPending:
			slog.WarnContext(ctx, "purchase status is not completed", slog.String("purchase status", string(purchaseMessage.Status)))

			// NOTE: 意図しないキューメッセージが送信されてきたと見なし、DLQへ送信
			sqsClient, err := sqs_client.NewSQSClient(ctx, configuration.Get().AWSConfig)
			if err != nil {
				slog.ErrorContext(ctx, "failed to create sqs client")
				return err
			}

			// NOTE: DLQメッセージの作成
			dlqMsgBody := models.DLQMessage{
				OriginalMessage: purchaseMessage,
				Error: models.Error{
					Code:    models.ErrorCodeInvalidStatus,
					Message: "The status of the queue message sent had an invalid status.",
				},
				TimeStamp:      time.Now(),
				LambdaFunction: configuration.Get().API.ServiceName,
			}

			// DLQメッセージ送信
			if err := sqsClient.SendDLQMessage(ctx, configuration.Get().SQS.DlqURL, dlqMsgBody); err != nil {
				return err
			}

		default:
			slog.WarnContext(ctx, "purchase status is not completed", slog.String("purchase status", string(purchaseMessage.Status)))
		}
	}

	return nil
}

func main() {
	lambda.Start(handler)
}

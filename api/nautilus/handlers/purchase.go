package handlers

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/takeuchima0/async_serverless_application_demo/api/nautilus/configuration"
	"github.com/takeuchima0/async_serverless_application_demo/api/nautilus/library/sqs_client"
	"github.com/takeuchima0/async_serverless_application_demo/api/nautilus/models"

	constants "github.com/takeuchima0/async_serverless_application_demo/api/nautilus/constants"
)

// 商品購入API
func PurchaseHandler(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()

	if err := configuration.Load(ctx); err != nil {
		http.Error(w, "Error loading configuration file", http.StatusInternalServerError)
		return
	}

	// Validate HTTP Method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Validate API KEY
	apiKey := r.Header.Get("X_API_KEY")
	if apiKey != string(constants.NautilusAPIKey) {
		http.Error(w, "Invalid X-API-KEY", http.StatusUnauthorized)
		return
	}

	// Parse request body
	var request models.PurchaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate UUID
	orderID := uuid.New().String()

	// SQS Clientの初期化
	sqsClient, err := sqs_client.NewSQSClient(ctx, configuration.Get().AWSConfig)
	if err != nil {
		http.Error(w, "Error initializing SQS client", http.StatusInternalServerError)
		return
	}

	// Queue のメッセージ作成
	uid := 10010024
	status := models.PurchaseStatusCompleted
	queueMsgBody := models.PurchaseQueueMessage{
		UserID:  uid,
		OrderID: orderID,
		Status:  status,
	}

	// SQSへメッセージ送信
	if err := sqsClient.SendPurchaseMessage(ctx, configuration.Get().SQS.NotificationsQueueURL, queueMsgBody); err != nil {
		slog.ErrorContext(ctx, "failed to send purchase message to sqs queue", "error", err.Error())
		http.Error(w, "Error sending message to SQS", http.StatusInternalServerError)
		return
	}

	// Mock response
	response := models.PurchaseResponse{
		Status:  models.PurchaseStatusCompleted,
		OrderID: orderID,
	}

	// クライアントにレスポンスを返す
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO: 配送APIをコール、非同期で対象のユーザへ購入した商品の配送準備を行う
}

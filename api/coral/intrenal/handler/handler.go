package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tamaco489/async_serverless_application_demo/api/coral/intrenal/configuration"

	dynamo_db "github.com/tamaco489/async_serverless_application_demo/api/coral/intrenal/library/dynamodb"
)

type CoralHandler struct {
	Config         configuration.Config
	DynamoDBClient dynamo_db.DynamoDBService
}

func NewHandler(ctx context.Context) (*CoralHandler, error) {
	if err := configuration.Load(ctx); err != nil {
		return nil, err
	}

	cnf := configuration.Get()
	dynamodbClient := dynamo_db.NewDynamoDBWrapper(cnf.AWSConfig, "Users")

	return &CoralHandler{
		Config:         cnf,
		DynamoDBClient: dynamodbClient,
	}, nil
}

func (h *CoralHandler) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"status": "ok"}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *CoralHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Body == nil {
		http.Error(w, "Empty request body", http.StatusBadRequest)
		return
	}

	var user map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if _, err := h.DynamoDBClient.CreateUser(r.Context(), user); err != nil {
		http.Error(w, fmt.Sprintf("Failed to create user in DynamoDB: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User created: %v", user)
}

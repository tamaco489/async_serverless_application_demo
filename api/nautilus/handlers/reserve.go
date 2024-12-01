package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/tamaco489/async_serverless_application_demo/api/nautilus/models"
)

// 在庫確保API
func ReserveHandler(w http.ResponseWriter, r *http.Request) {

	// Validate HTTP Method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse request body
	var request models.ReserveRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate UUID
	reservedID := uuid.New().String()

	// Mock response
	response := models.ReserveResponse{
		Status:     models.ReserveStatusConfirmed,
		ReservedID: reservedID,
	}

	// Encode response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

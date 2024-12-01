package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/takeuchima0/async_serverless_application_demo/api/coral/intrenal/configuration"
)

type CoralHandler struct {
	Config configuration.Config
}

func NewHandler(ctx context.Context) (*CoralHandler, error) {
	if err := configuration.Load(ctx); err != nil {
		return nil, err
	}
	return &CoralHandler{
		Config: configuration.Get(),
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

	// Check if the body is empty
	if r.Body == nil {
		http.Error(w, "Empty request body", http.StatusBadRequest)
		return
	}

	// Decode the JSON body
	var user map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Respond with the created user
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User created: %v", user)
}

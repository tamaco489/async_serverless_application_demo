package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *CoralHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Body == nil {
		http.Error(w, "empty request body", http.StatusBadRequest)
		return
	}

	var user map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	responseUser, err := h.userUseCase.CreateUser(r.Context(), user)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create user in DynamoDB: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(responseUser); err != nil {
		http.Error(w, fmt.Sprintf("failed to encode response: %v", err), http.StatusInternalServerError)
		return
	}
}

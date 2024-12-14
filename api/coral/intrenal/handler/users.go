package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (ch *CoralHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
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

	response, err := ch.userUseCase.CreateUser(r.Context(), user)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create user in dynamodb: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("failed to encode response: %v", err), http.StatusInternalServerError)
		return
	}
}

func (ch *CoralHandler) GetMeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not supported", http.StatusInternalServerError)
		return
	}

	// 本来であれば、middlewareでhttp headerに設定されたjwtの正当性検証を行い、その過程を経て以降の処理を行うのが望ましい

	// ここではこのuidは正当なユーザのuidであるものとして扱う
	uid := "0193c504-2b40-79c6-8a35-acfe1e938ac5"

	response, err := ch.userUseCase.GetMeUser(r.Context(), uid)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed fetch user me: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("failed to encode response: %v", err), http.StatusInternalServerError)
		return
	}
}

func (ch *CoralHandler) GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid := vars["userID"]

	if uid == "" {
		http.Error(w, "userID is required", http.StatusBadRequest)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "method not supported", http.StatusMethodNotAllowed)
	}

	response, err := ch.userUseCase.GetUserByID(r.Context(), uid)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed fetch user by id: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("failed to encode response: %v", err), http.StatusInternalServerError)
		return
	}
}

package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type SuccessResponse struct {
	Message string `json:"message,omitempty"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func (h *CoralHandler) JSONResponse(statusCode int, data interface{}) (events.APIGatewayProxyResponse, error) {
	responseBody, err := json.Marshal(data)
	if err != nil {
		slog.Error("error marshalling response body", "error", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers:    map[string]string{"Content-Type": "application/json"},
			Body:       "error marshalling response body",
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(responseBody),
	}, nil
}

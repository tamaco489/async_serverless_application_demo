package handler

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/takeuchima0/async_serverless_application_demo/api/coral/intrenal/configuration"
)

type CoralHandler struct{}

func NewHandler(ctx context.Context) (*CoralHandler, error) {
	if err := configuration.Load(ctx); err != nil {
		return nil, err
	}
	return &CoralHandler{}, nil
}

func (h *CoralHandler) Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	slog.InfoContext(ctx, "start coral handler for request")

	path := request.PathParameters["proxy"]
	if path == "healthcheck" {
		return h.JSONResponse(http.StatusOK, SuccessResponse{Message: "ok"})
	}

	return h.JSONResponse(http.StatusInternalServerError, ErrorResponse{Error: "Internal Server Error"})
}

package handler

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"regexp"

	"github.com/aws/aws-lambda-go/events"
	"github.com/takeuchima0/async_serverless_application_demo/api/ibis/internal/configuration"
	"github.com/takeuchima0/async_serverless_application_demo/api/ibis/internal/usecase"
)

type IbisHandler struct{}

func NewHandler(ctx context.Context) (*IbisHandler, error) {
	if err := configuration.Load(ctx); err != nil {
		return nil, err
	}
	return &IbisHandler{}, nil
}

// 成功レスポンス用の構造体
type SuccessResponse struct {
	Message string `json:"message,omitempty"`
	URL     string `json:"url,omitempty"`
}

// エラーレスポンス用の構造体
type ErrorResponse struct {
	Error string `json:"error"`
}

func (h *IbisHandler) JSONResponse(statusCode int, data interface{}) (events.APIGatewayProxyResponse, error) {
	responseBody, err := json.Marshal(data)
	if err != nil {
		slog.Error("error marshaling response", "error", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers:    map[string]string{"Content-Type": "application/json"},
			Body:       "error marshaling response",
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(responseBody),
	}, nil
}

func (h *IbisHandler) Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	slog.InfoContext(ctx, "start ibis api server...")

	path := request.PathParameters["proxy"]

	if path == "healthcheck" {
		return h.JSONResponse(http.StatusOK, SuccessResponse{Message: "ok"})
	}

	imageDownloadPattern := regexp.MustCompile(`^image/download/(\d+)$`)
	imageUploadPattern := regexp.MustCompile(`^image/upload/(\d+)$`)

	switch {
	case imageDownloadPattern.MatchString(path):
		matches := imageDownloadPattern.FindStringSubmatch(path)
		if len(matches) > 1 {
			uid := matches[1]
			u := usecase.NewGetDownloadImageURL()
			url, err := u.GetDownloadImageURL(ctx, uid)
			if err != nil {
				return h.JSONResponse(http.StatusInternalServerError, ErrorResponse{Error: "error generating download presigned URL"})
			}
			return h.JSONResponse(http.StatusOK, SuccessResponse{URL: url})
		}

	case imageUploadPattern.MatchString(path):
		matches := imageUploadPattern.FindStringSubmatch(path)
		if len(matches) > 1 {
			uid := matches[1]
			u := usecase.NewGetUploadImageURL()
			url, err := u.GetUploadImageURL(ctx, uid)
			if err != nil {
				return h.JSONResponse(http.StatusInternalServerError, ErrorResponse{Error: "error generating upload presigned URL"})
			}
			return h.JSONResponse(http.StatusOK, SuccessResponse{URL: url})
		}

	default:
		return h.JSONResponse(http.StatusNotFound, ErrorResponse{Error: "not found"})
	}

	return h.JSONResponse(http.StatusInternalServerError, ErrorResponse{Error: "internal server error"})
}

package main

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/takeuchima0/async_serverless_application_demo/api/coral/intrenal/configuration"
	"github.com/takeuchima0/async_serverless_application_demo/api/coral/intrenal/handler"
)

func main() {
	ctx := context.Background()
	coralHandler, err := handler.NewHandler(ctx)
	if err != nil {
		panic(err)
	}
	slog.InfoContext(ctx, "Configuration loaded",
		"env", configuration.Get().API.Env,
		"service_name", configuration.Get().API.ServiceName,
		"port", configuration.Get().API.Port,
	)

	// GET: health check API
	http.HandleFunc("/coral/v3/healthcheck", coralHandler.HealthCheckHandler)

	// POST: create new user
	http.HandleFunc("/coral/v3/users", coralHandler.CreateUserHandler)

	// GET: get user by user_id
	// http.HandleFunc("/coral/v3/users/{userID}", handler.HealthCheckHandler)

	// PUT: update user by user_id
	// http.HandleFunc("/coral/v3/users/{userID}", handler.HealthCheckHandler)

	// GET: get users only administrator
	// http.HandleFunc("/coral/v3/users", handler.HealthCheckHandler)

	// DELETE: delete user by user_id only administrator
	// http.HandleFunc("/coral/v3/users/{userID}", handler.HealthCheckHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		slog.ErrorContext(ctx, "ListenAndServe error", "error", err)
	}
}

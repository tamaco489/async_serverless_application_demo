package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/tamaco489/async_serverless_application_demo/api/coral/intrenal/configuration"
	"github.com/tamaco489/async_serverless_application_demo/api/coral/intrenal/handler"
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
	http.HandleFunc("/coral/v3/users/me", coralHandler.GetMeHandler)

	// PUT: update user by user_id
	// http.HandleFunc("/coral/v3/users/{userID}", handler.HealthCheckHandler)

	// GET: get users only administrator
	// http.HandleFunc("/coral/v3/users", handler.HealthCheckHandler)

	// DELETE: delete user by user_id only administrator
	// http.HandleFunc("/coral/v3/users/{userID}", handler.HealthCheckHandler)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", configuration.Get().API.Port), nil); err != nil {
		slog.ErrorContext(ctx, "ListenAndServe error", "error", err)
	}
}

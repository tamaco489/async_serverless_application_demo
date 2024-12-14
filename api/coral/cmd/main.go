package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
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

	r := mux.NewRouter()

	// GET: health check API
	r.HandleFunc("/coral/v3/healthcheck", coralHandler.HealthCheckHandler).Methods(http.MethodGet)

	// POST: create new user API
	r.HandleFunc("/coral/v3/users", coralHandler.CreateUserHandler).Methods(http.MethodPost)

	// GET: get me user API
	r.HandleFunc("/coral/v3/users/me", coralHandler.GetMeHandler).Methods(http.MethodGet)

	// GET/PUT: get or put user by user_id API
	r.HandleFunc("/coral/v3/users/{userID}", coralHandler.GetUserByIDHandler).Methods(http.MethodGet, http.MethodPut)

	// GET: get users only administrator
	// http.HandleFunc("/coral/v3/users", handler.HealthCheckHandler)

	// DELETE: delete user by user_id only administrator
	// http.HandleFunc("/coral/v3/users/{userID}", handler.HealthCheckHandler)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", configuration.Get().API.Port), r); err != nil {
		slog.ErrorContext(ctx, "ListenAndServe error", "error", err)
	}
}

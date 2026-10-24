// Package main
package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/davidaskenazy/devboard/internal/handler"
	"github.com/davidaskenazy/devboard/internal/logger"
	"github.com/davidaskenazy/devboard/internal/middlewares"
	"github.com/davidaskenazy/devboard/internal/server"
)

func main() {

	logger := logger.New(logger.ProductionConfig())

	srv := server.New(":8080")

	srv.Use(middlewares.Recovery(logger))
	srv.Use(middlewares.Logger(logger))

	healthHandler := handler.NewHealthHandler()

	srv.RegisterRoutes("GET /health", healthHandler)

	srv.RegisterRoutes("GET /panic", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("error provocado para probocar Recovery")
	}))

	logger.Info("Servidor iniciando :8080", slog.String("addr", ":8080"))

	if err := srv.Start(); err != nil {
		log.Fatalf("error al iniciar el servidor %v", err)
	}
}

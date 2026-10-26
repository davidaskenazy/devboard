// Package main
package main

import (
	"net/http"
	"os"

	_ "github.com/davidaskenazy/devboard/docs"
	"github.com/davidaskenazy/devboard/internal/handler"
	"github.com/davidaskenazy/devboard/internal/logger"
	"github.com/davidaskenazy/devboard/internal/middlewares"
	"github.com/davidaskenazy/devboard/internal/server"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Devboard API
// @version 1.0
// @description API REST para gestion de proyectos
// @contact.name David Askenazy
// @host localhost:8080
// @BasePath /api/v1
func main() {

	logger := logger.New(logger.DefaultConfig())

	srv := server.New(":8080", logger)

	srv.Use(middlewares.Recovery(logger))
	srv.Use(middlewares.Logger(logger))

	healthHandler := handler.NewHealthHandler()

	srv.RegisterRoutes("GET /docs/", httpSwagger.WrapHandler)
	srv.RegisterRoutes("GET /health", healthHandler)
	srv.RegisterRoutes("GET /panic", http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {
		panic("error provocado para probocar Recovery")
	}))

	if err := srv.Start(); err != nil {
		logger.Error("error al iniciar el servidor", "error", err)
		os.Exit(1)
		//log.Fatalf("error al iniciar el servidor %v", err)
	}
}

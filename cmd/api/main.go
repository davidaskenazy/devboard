// Package main
package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

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
	//1. Configuracion del entorno
	env := os.Getenv("GO_ENV")

	//2. Logger
	var log *slog.Logger
	if env == "production" {
		log = logger.New(logger.ProductionConfig())
	} else {
		log = logger.New(logger.DefaultConfig())
	}

	//3. Dependencias compartidas
	/*
		validate := validator.New()
		notifier := notification.NewLogNotifier(log)
		_ = notifier
	*/
	//4. Servidor con Functional options
	srv := server.New(":8080",
		server.WithLogger(log),
		server.WithReadTimeout(15*time.Second),
		server.WithWriteTimeout(30*time.Second),
	)

	//5. Handlers
	healthHandler := handler.NewHealthHandler()

	//6.Registro de rutas
	srv.RegisterRoutes("GET /docs/", httpSwagger.WrapHandler)
	srv.RegisterRoutes("GET /health", healthHandler)
	srv.RegisterRoutes("GET /panic", http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {
		panic("error provocado para probocar Recovery")
	}))

	//7. Middleware chain
	srv.UseChain(
		middlewares.Recovery(log),
		middlewares.Logger(log),
	)

	//8. Arrancar shutdown limpio
	if err := srv.Start(); err != nil {
		log.Error("error al iniciar el servidor", "error", err)
		os.Exit(1)
		//log.Fatalf("error al iniciar el servidor %v", err)
	}
}

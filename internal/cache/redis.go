// Package cache
package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// Config contiene la configuration del cliente Redis
type Config struct {
	Addr     string //host:puerto
	Password string
	DB       int
	Timeout  time.Duration
}

// DefaultConfig retorna configuration para desarrollo local
func DefaultConfig() Config {
	return Config{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		Timeout:  5 * time.Second,
	}
}

// NewClient crea y verifica un cliente en redis
func NewClient(ctx context.Context, cfg Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         cfg.Addr,
		Password:     cfg.Password,
		DB:           cfg.DB,
		DialTimeout:  cfg.Timeout,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
	})

	pingCtx, cancel := context.WithTimeout(ctx, cfg.Timeout)
	defer cancel()

	if err := client.Ping(pingCtx).Err(); err != nil {
		return nil, fmt.Errorf("verificando conexion a Redis: %w", err)
	}

	return client, nil
}

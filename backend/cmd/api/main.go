// Package main provides the entrypoint for the anti-pattern Echo API.
package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"immortal-architecture-bad-api/backend/internal/controller"
	dbpkg "immortal-architecture-bad-api/backend/internal/db"
	"immortal-architecture-bad-api/backend/internal/db/sqlc"
	"immortal-architecture-bad-api/backend/internal/service"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	pool, err := dbpkg.NewPool(ctx, dbURL)
	if err != nil {
		log.Fatalf("failed to create db pool: %v", err)
	}
	defer pool.Close()

	queries := sqldb.New(pool)
	accountService := service.NewAccountService(queries)

	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.GET("/healthz", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	apiGroup := e.Group("/api")
	accountController := controller.NewAccountController(accountService)
	accountController.Register(apiGroup.Group("/accounts"))

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}

	serverErr := make(chan error, 1)
	go func() {
		if err := e.Start(":" + port); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErr <- err
		}
	}()

	select {
	case err := <-serverErr:
		log.Fatalf("server error: %v", err)
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := e.Shutdown(shutdownCtx); err != nil {
			log.Printf("graceful shutdown failed: %v", err)
		}
	}
}

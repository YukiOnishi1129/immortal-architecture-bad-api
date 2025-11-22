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
	openapi "immortal-architecture-bad-api/backend/internal/generated/openapi"
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

	accountService := service.NewAccountService(pool)
	templateService := service.NewTemplateService(pool)
	noteService := service.NewNoteService(pool)

	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	corsOrigin := os.Getenv("CLIENT_ORIGIN")
	corsConfig := middleware.CORSConfig{
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
		},
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
	}
	if corsOrigin != "" {
		corsConfig.AllowOrigins = []string{corsOrigin}
	}
	e.Use(middleware.CORSWithConfig(corsConfig))

	e.GET("/healthz", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	ctrl := controller.NewController(accountService, templateService, noteService)
	openapi.RegisterHandlers(e, ctrl)

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

// Package main provides the entrypoint for the anti-pattern Echo API.
package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.GET("/healthz", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}

	if err := e.Start(":" + port); err != nil {
		log.Fatal(err)
	}
}

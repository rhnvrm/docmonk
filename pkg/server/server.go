package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Serve() error {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/api", hello)

	// Start server
	return e.Start(":1323")
}

// hello is a Root Handler for the API.
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

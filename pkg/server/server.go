package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var handleRoot func() echo.HandlerFunc = func() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Root!")
	}
}

var handleRootStatic func() echo.HandlerFunc = func() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Root!")
	}
}

func Serve() error {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", handleRoot())
	e.GET("/assets/*", handleRoot())

	e.GET("/api", helloApi)

	// Start server
	return e.Start(":1323")
}

// helloApi is a Root Handler for the API.
func helloApi(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, API!")
}

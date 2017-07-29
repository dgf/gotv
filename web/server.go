package web

import (
	"net/http"

	"github.com/dgf/gotv/api"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Server context
type Server struct {
	echo.Context // extend context
	Name         string
	Repository   api.Repository
}

// New creates handler to serve
func New(name string, repo api.Repository) http.Handler {
	e := echo.New()

	// configure middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${method} ${status} ${uri}\n",
	}))
	e.Use(middleware.Recover())

	// bind context
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return h(&Server{Context: c, Name: name, Repository: repo})
		}
	})

	// route paths
	e.GET("/status", GetStatus)

	return e
}

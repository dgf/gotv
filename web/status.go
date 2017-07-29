package web

import (
	"net/http"

	"github.com/labstack/echo"
)

// Status details STO
type Status struct {
	Name  string `json:"name"`
	Games int    `json:"games"`
}

// GetStatus serves JSON status
func GetStatus(c echo.Context) error {
	server := c.(*Server)

	count, err := server.Repository.Games()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, Status{
		Name:  server.Name,
		Games: count,
	})
}

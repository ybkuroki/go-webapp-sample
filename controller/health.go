package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetHealthCheck returns whether this application is alive or not.
func GetHealthCheck() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "healthy")
	}
}

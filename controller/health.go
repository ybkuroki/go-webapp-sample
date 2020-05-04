package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetHealthCheck is
func GetHealthCheck() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "healthy")
	}
}

package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetLoginStatus is
func GetLoginStatus() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, true)
	}
}

// GetLoginAccount is
func GetLoginAccount() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "{name: test}")
	}
}

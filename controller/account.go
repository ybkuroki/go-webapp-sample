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
		return c.JSON(http.StatusOK, &Account{ID: 1, Name: "test"})
	}
}

// Account is struct (TODO)
type Account struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

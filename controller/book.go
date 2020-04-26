package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/service"
)

// GetBookList is
func GetBookList() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, service.FindAllBooks())
	}
}

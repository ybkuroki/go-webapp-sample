package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/service"
)

// GetCategoryList is
func GetCategoryList() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, service.FindAllCategories())
	}
}

// GetFormatList is
func GetFormatList() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, service.FindAllFormats())
	}
}

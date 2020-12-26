package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/mycontext"
	"github.com/ybkuroki/go-webapp-sample/service"
)

// GetCategoryList returns the list of all categories.
func GetCategoryList(context mycontext.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, service.FindAllCategories(context))
	}
}

// GetFormatList returns the list of all formats.
func GetFormatList(context mycontext.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, service.FindAllFormats(context))
	}
}

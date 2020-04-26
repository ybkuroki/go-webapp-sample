package controller

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/service"
)

// GetBookList is
func GetBookList(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, service.FindAllBooks(db))
	}
}

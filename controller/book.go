package controller

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/model"
)

// BookList is
func BookList(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		book := model.NewBook("test", "123-123-1", model.NewCategory("technical"), model.NewFormat("paper"))
		result, _ := book.FindByID(db, 1)
		return c.JSON(http.StatusOK, result)
	}
}

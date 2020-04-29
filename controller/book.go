package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/model/dto"
	"github.com/ybkuroki/go-webapp-sample/service"
)

// GetBookList is
func GetBookList() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, service.FindAllBooks())
	}
}

// PostBookRegist is
func PostBookRegist() echo.HandlerFunc {
	return func(c echo.Context) error {
		dto := dto.NewRegBookDto()
		if err := c.Bind(dto); err != nil {
			return c.JSON(http.StatusBadRequest, dto)
		}
		book, result := service.RegisterBook(dto)
		if result != nil {
			return c.JSON(http.StatusBadRequest, result)
		}
		return c.JSON(http.StatusOK, book)
	}
}

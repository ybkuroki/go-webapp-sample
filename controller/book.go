package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/model/dto"
	"github.com/ybkuroki/go-webapp-sample/service"
)

// GetBookList is
func GetBookList() echo.HandlerFunc {
	return func(c echo.Context) error {
		page, _ := strconv.Atoi(c.QueryParam("page"))
		size, _ := strconv.Atoi(c.QueryParam("size"))

		return c.JSON(http.StatusOK, service.FindAllBooksByPage(page, size))
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

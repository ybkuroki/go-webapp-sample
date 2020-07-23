package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/model/dto"
	"github.com/ybkuroki/go-webapp-sample/service"
)

// GetBookList returns the list of all books.
func GetBookList() echo.HandlerFunc {
	return func(c echo.Context) error {
		page, _ := strconv.Atoi(c.QueryParam("page"))
		size, _ := strconv.Atoi(c.QueryParam("size"))

		return c.JSON(http.StatusOK, service.FindAllBooksByPage(page, size))
	}
}

// GetBookSearch returns the list of matched books by searching.
func GetBookSearch() echo.HandlerFunc {
	return func(c echo.Context) error {
		title := c.QueryParam("query")
		page, _ := strconv.Atoi(c.QueryParam("page"))
		size, _ := strconv.Atoi(c.QueryParam("size"))

		return c.JSON(http.StatusOK, service.FindBooksByTitle(title, page, size))
	}
}

// PostBookRegist register a new book by http post.
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

// PostBookEdit edit the existing book by http post.
func PostBookEdit() echo.HandlerFunc {
	return func(c echo.Context) error {
		dto := dto.NewChgBookDto()
		if err := c.Bind(dto); err != nil {
			return c.JSON(http.StatusBadRequest, dto)
		}
		book, result := service.EditBook(dto)
		if result != nil {
			return c.JSON(http.StatusBadRequest, result)
		}
		return c.JSON(http.StatusOK, book)
	}
}

// PostBookDelete deletes the existing book by http post.
func PostBookDelete() echo.HandlerFunc {
	return func(c echo.Context) error {
		dto := dto.NewChgBookDto()
		if err := c.Bind(dto); err != nil {
			return c.JSON(http.StatusBadRequest, dto)
		}
		book, result := service.DeleteBook(dto)
		if result != nil {
			return c.JSON(http.StatusBadRequest, result)
		}
		return c.JSON(http.StatusOK, book)
	}
}

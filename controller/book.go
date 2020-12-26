package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/model/dto"
	"github.com/ybkuroki/go-webapp-sample/mycontext"
	"github.com/ybkuroki/go-webapp-sample/service"
)

// GetBookList returns the list of all books.
func GetBookList(context mycontext.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		page, _ := strconv.Atoi(c.QueryParam("page"))
		size, _ := strconv.Atoi(c.QueryParam("size"))

		return c.JSON(http.StatusOK, service.FindAllBooksByPage(context, page, size))
	}
}

// GetBookSearch returns the list of matched books by searching.
func GetBookSearch(context mycontext.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		title := c.QueryParam("query")
		page, _ := strconv.Atoi(c.QueryParam("page"))
		size, _ := strconv.Atoi(c.QueryParam("size"))

		return c.JSON(http.StatusOK, service.FindBooksByTitle(context, title, page, size))
	}
}

// PostBookRegist register a new book by http post.
func PostBookRegist(context mycontext.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		dto := dto.NewRegBookDto()
		if err := c.Bind(dto); err != nil {
			return c.JSON(http.StatusBadRequest, dto)
		}
		book, result := service.RegisterBook(context, dto)
		if result != nil {
			return c.JSON(http.StatusBadRequest, result)
		}
		return c.JSON(http.StatusOK, book)
	}
}

// PostBookEdit edit the existing book by http post.
func PostBookEdit(context mycontext.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		dto := dto.NewChgBookDto()
		if err := c.Bind(dto); err != nil {
			return c.JSON(http.StatusBadRequest, dto)
		}
		book, result := service.EditBook(context, dto)
		if result != nil {
			return c.JSON(http.StatusBadRequest, result)
		}
		return c.JSON(http.StatusOK, book)
	}
}

// PostBookDelete deletes the existing book by http post.
func PostBookDelete(context mycontext.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		dto := dto.NewChgBookDto()
		if err := c.Bind(dto); err != nil {
			return c.JSON(http.StatusBadRequest, dto)
		}
		book, result := service.DeleteBook(context, dto)
		if result != nil {
			return c.JSON(http.StatusBadRequest, result)
		}
		return c.JSON(http.StatusOK, book)
	}
}

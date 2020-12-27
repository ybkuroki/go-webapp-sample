package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/model/dto"
	"github.com/ybkuroki/go-webapp-sample/mycontext"
	"github.com/ybkuroki/go-webapp-sample/service"
)

// BookController is
type BookController struct {
	context mycontext.Context
	service *service.BookService
}

// NewBookController is
func NewBookController(context mycontext.Context) *BookController {
	return &BookController{context: context, service: service.NewBookService(context)}
}

// GetBookList returns the list of all books.
func (controller *BookController) GetBookList(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	size, _ := strconv.Atoi(c.QueryParam("size"))

	return c.JSON(http.StatusOK, controller.service.FindAllBooksByPage(page, size))
}

// GetBookSearch returns the list of matched books by searching.
func (controller *BookController) GetBookSearch(c echo.Context) error {
	title := c.QueryParam("query")
	page, _ := strconv.Atoi(c.QueryParam("page"))
	size, _ := strconv.Atoi(c.QueryParam("size"))

	return c.JSON(http.StatusOK, controller.service.FindBooksByTitle(title, page, size))
}

// PostBookRegist register a new book by http post.
func (controller *BookController) PostBookRegist(c echo.Context) error {
	dto := dto.NewRegBookDto()
	if err := c.Bind(dto); err != nil {
		return c.JSON(http.StatusBadRequest, dto)
	}
	book, result := controller.service.RegisterBook(dto)
	if result != nil {
		return c.JSON(http.StatusBadRequest, result)
	}
	return c.JSON(http.StatusOK, book)
}

// PostBookEdit edit the existing book by http post.
func (controller *BookController) PostBookEdit(c echo.Context) error {
	dto := dto.NewChgBookDto()
	if err := c.Bind(dto); err != nil {
		return c.JSON(http.StatusBadRequest, dto)
	}
	book, result := controller.service.EditBook(dto)
	if result != nil {
		return c.JSON(http.StatusBadRequest, result)
	}
	return c.JSON(http.StatusOK, book)
}

// PostBookDelete deletes the existing book by http post.
func (controller *BookController) PostBookDelete(c echo.Context) error {
	dto := dto.NewChgBookDto()
	if err := c.Bind(dto); err != nil {
		return c.JSON(http.StatusBadRequest, dto)
	}
	book, result := controller.service.DeleteBook(dto)
	if result != nil {
		return c.JSON(http.StatusBadRequest, result)
	}
	return c.JSON(http.StatusOK, book)
}

package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/model/dto"
	"github.com/ybkuroki/go-webapp-sample/mycontext"
	"github.com/ybkuroki/go-webapp-sample/service"
)

// BookController is a controller for managing books.
type BookController struct {
	context mycontext.Context
	service *service.BookService
}

// NewBookController is constructor.
func NewBookController(context mycontext.Context) *BookController {
	return &BookController{context: context, service: service.NewBookService(context)}
}

// GetBook returns one record matched book's id.
func (controller *BookController) GetBook(c echo.Context) error {
	return c.JSON(http.StatusOK, controller.service.FindByID(c.Param("id")))
}

// GetBookList returns the list of matched books by searching.
func (controller *BookController) GetBookList(c echo.Context) error {
	return c.JSON(http.StatusOK, controller.service.FindBooksByTitle(c.QueryParam("query"), c.QueryParam("page"), c.QueryParam("size")))
}

// CreateBook create a new book by http post.
func (controller *BookController) CreateBook(c echo.Context) error {
	dto := dto.NewBookDto()
	if err := c.Bind(dto); err != nil {
		return c.JSON(http.StatusBadRequest, dto)
	}
	book, result := controller.service.RegisterBook(dto)
	if result != nil {
		return c.JSON(http.StatusBadRequest, result)
	}
	return c.JSON(http.StatusOK, book)
}

// UpdateBook update the existing book by http post.
func (controller *BookController) UpdateBook(c echo.Context) error {
	dto := dto.NewBookDto()
	if err := c.Bind(dto); err != nil {
		return c.JSON(http.StatusBadRequest, dto)
	}
	book, result := controller.service.EditBook(dto, c.Param("id"))
	if result != nil {
		return c.JSON(http.StatusBadRequest, result)
	}
	return c.JSON(http.StatusOK, book)
}

// DeleteBook deletes the existing book by http post.
func (controller *BookController) DeleteBook(c echo.Context) error {
	book, result := controller.service.DeleteBook(c.Param("id"))
	if result != nil {
		return c.JSON(http.StatusBadRequest, result)
	}
	return c.JSON(http.StatusOK, book)
}

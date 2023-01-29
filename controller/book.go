package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/container"
	"github.com/ybkuroki/go-webapp-sample/model/dto"
	"github.com/ybkuroki/go-webapp-sample/service"
)

// BookController is a controller for managing books.
type BookController interface {
	GetBook(c echo.Context) error
	GetBookList(c echo.Context) error
	CreateBook(c echo.Context) error
	UpdateBook(c echo.Context) error
	DeleteBook(c echo.Context) error
}

type bookController struct {
	container container.Container
	service   service.BookService
}

// NewBookController is constructor.
func NewBookController(container container.Container) BookController {
	return &bookController{container: container, service: service.NewBookService(container)}
}

// GetBook returns one record matched book's id.
// @Summary Get a book
// @Description Get a book
// @Tags Books
// @Accept  json
// @Produce  json
// @Param book_id path int true "Book ID"
// @Success 200 {object} model.Book "Success to fetch data."
// @Failure 400 {string} message "Failed to fetch data."
// @Failure 401 {boolean} bool "Failed to the authentication. Returns false."
// @Router /books/{book_id} [get]
func (controller *bookController) GetBook(c echo.Context) error {
	book, err := controller.service.FindByID(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, book)
}

// GetBookList returns the list of matched books by searching.
// @Summary Get a book list
// @Description Get the list of matched books by searching
// @Tags Books
// @Accept  json
// @Produce  json
// @Param query query string false "Keyword"
// @Param page query int false "Page number"
// @Param size query int false "Item size per page"
// @Success 200 {object} model.Page "Success to fetch a book list."
// @Failure 400 {string} message "Failed to fetch data."
// @Failure 401 {boolean} bool "Failed to the authentication. Returns false."
// @Router /books [get]
func (controller *bookController) GetBookList(c echo.Context) error {
	book, err := controller.service.FindBooksByTitle(c.QueryParam("query"), c.QueryParam("page"), c.QueryParam("size"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, book)
}

// CreateBook create a new book by http post.
// @Summary Create a new book
// @Description Create a new book
// @Tags Books
// @Accept  json
// @Produce  json
// @Param data body dto.BookDto true "a new book data for creating"
// @Success 200 {object} model.Book "Success to create a new book."
// @Failure 400 {string} message "Failed to the registration."
// @Failure 401 {boolean} bool "Failed to the authentication. Returns false."
// @Router /books [post]
func (controller *bookController) CreateBook(c echo.Context) error {
	dto := dto.NewBookDto(controller.container.GetMessages())
	if err := c.Bind(dto); err != nil {
		return c.JSON(http.StatusBadRequest, dto)
	}
	book, result := controller.service.CreateBook(dto)
	if result != nil {
		return c.JSON(http.StatusBadRequest, result)
	}
	return c.JSON(http.StatusOK, book)
}

// UpdateBook update the existing book by http put.
// @Summary Update the existing book
// @Description Update the existing book
// @Tags Books
// @Accept  json
// @Produce  json
// @Param book_id path int true "Book ID"
// @Param data body dto.BookDto true "the book data for updating"
// @Success 200 {object} model.Book "Success to update the existing book."
// @Failure 400 {string} message "Failed to the update."
// @Failure 401 {boolean} bool "Failed to the authentication. Returns false."
// @Router /books/{book_id} [put]
func (controller *bookController) UpdateBook(c echo.Context) error {
	dto := dto.NewBookDto(controller.container.GetMessages())
	if err := c.Bind(dto); err != nil {
		return c.JSON(http.StatusBadRequest, dto)
	}
	book, result := controller.service.UpdateBook(dto, c.Param("id"))
	if result != nil {
		return c.JSON(http.StatusBadRequest, result)
	}
	return c.JSON(http.StatusOK, book)
}

// DeleteBook deletes the existing book by http delete.
// @Summary Delete the existing book
// @Description Delete the existing book
// @Tags Books
// @Accept  json
// @Produce  json
// @Param book_id path int true "Book ID"
// @Success 200 {object} model.Book "Success to delete the existing book."
// @Failure 400 {string} message "Failed to the delete."
// @Failure 401 {boolean} bool "Failed to the authentication. Returns false."
// @Router /books/{book_id} [delete]
func (controller *bookController) DeleteBook(c echo.Context) error {
	book, result := controller.service.DeleteBook(c.Param("id"))
	if result != nil {
		return c.JSON(http.StatusBadRequest, result)
	}
	return c.JSON(http.StatusOK, book)
}

package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/ybkuroki/go-webapp-sample/container"
	"github.com/ybkuroki/go-webapp-sample/model"
	"github.com/ybkuroki/go-webapp-sample/model/dto"
	"github.com/ybkuroki/go-webapp-sample/test"
)

type BookDtoForBindError struct {
	Title      string
	Isbn       string
	CategoryID string
	FormatID   string
}

func TestGetBook_Success(t *testing.T) {
	router, container := test.Prepare(false)

	book := NewBookController(container)
	router.GET(APIBooksID, func(c echo.Context) error { return book.GetBook(c) })

	setUpTestData(container)

	uri := test.NewRequestBuilder().URL(APIBooks).PathParams("1").Build().GetRequestURL()
	req := httptest.NewRequest("GET", uri, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	entity := &model.Book{}
	data, _ := entity.FindByID(container.GetRepository(), 1)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, test.ConvertToString(data), rec.Body.String())
}

func TestGetBookList_Success(t *testing.T) {
	router, container := test.Prepare(false)

	book := NewBookController(container)
	router.GET(APIBooks, func(c echo.Context) error { return book.GetBookList(c) })

	setUpTestData(container)

	uri := test.NewRequestBuilder().URL(APIBooks).RequestParams("query", "Test").RequestParams("page", "0").RequestParams("size", "5").Build().GetRequestURL()
	req := httptest.NewRequest("GET", uri, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	entity := &model.Book{}
	data, _ := entity.FindByTitle(container.GetRepository(), "Test", "0", "5")

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, test.ConvertToString(data), rec.Body.String())
}

func TestCreateBook_Success(t *testing.T) {
	router, container := test.Prepare(false)

	book := NewBookController(container)
	router.POST(APIBooks, func(c echo.Context) error { return book.CreateBook(c) })

	param := createBookForCreate()
	req := test.NewJsonRequest("POST", APIBooks, param)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	entity := &model.Book{}
	data, _ := entity.FindByID(container.GetRepository(), 1)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, test.ConvertToString(data), rec.Body.String())
}

func TestCreateBook_BindError(t *testing.T) {
	router, container := test.Prepare(false)

	book := NewBookController(container)
	router.POST(APIBooks, func(c echo.Context) error { return book.CreateBook(c) })

	param := createBookForBindError()
	req := test.NewJsonRequest("POST", APIBooks, param)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	result := createResultForBindError()
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.JSONEq(t, test.ConvertToString(result), rec.Body.String())
}

func TestUpdateBook_Success(t *testing.T) {
	router, container := test.Prepare(false)

	book := NewBookController(container)
	router.PUT(APIBooksID, func(c echo.Context) error { return book.UpdateBook(c) })

	setUpTestData(container)

	param := createBookForUpdate()
	uri := test.NewRequestBuilder().URL(APIBooks).PathParams("1").Build().GetRequestURL()
	req := test.NewJsonRequest("PUT", uri, param)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	entity := &model.Book{}
	data, _ := entity.FindByID(container.GetRepository(), 1)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, test.ConvertToString(data), rec.Body.String())
}

func TestUpdateBook_BindError(t *testing.T) {
	router, container := test.Prepare(false)

	book := NewBookController(container)
	router.PUT(APIBooksID, func(c echo.Context) error { return book.UpdateBook(c) })

	setUpTestData(container)

	param := createBookForBindError()
	uri := test.NewRequestBuilder().URL(APIBooks).PathParams("1").Build().GetRequestURL()
	req := test.NewJsonRequest("PUT", uri, param)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	result := createResultForBindError()
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.JSONEq(t, test.ConvertToString(result), rec.Body.String())
}

func TestDeleteBook_Success(t *testing.T) {
	router, container := test.Prepare(false)

	book := NewBookController(container)
	router.DELETE(APIBooksID, func(c echo.Context) error { return book.DeleteBook(c) })

	setUpTestData(container)

	entity := &model.Book{}
	data, _ := entity.FindByID(container.GetRepository(), 1)

	uri := test.NewRequestBuilder().URL(APIBooks).PathParams("1").Build().GetRequestURL()
	req := test.NewJsonRequest("DELETE", uri, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, test.ConvertToString(data), rec.Body.String())
}

func setUpTestData(container container.Container) {
	model := model.NewBook("Test1", "123-123-123-1", 1, 1)
	repo := container.GetRepository()
	_, _ = model.Create(repo)
}

func createBookForCreate() *dto.BookDto {
	return &dto.BookDto{
		Title:      "Test1",
		Isbn:       "123-123-123-1",
		CategoryID: 1,
		FormatID:   1,
	}
}

func createBookForBindError() *BookDtoForBindError {
	return &BookDtoForBindError{
		Title:      "Test1",
		Isbn:       "123-123-123-1",
		CategoryID: "Test",
		FormatID:   "Test",
	}
}

func createResultForBindError() *dto.BookDto {
	return &dto.BookDto{
		Title:      "Test1",
		Isbn:       "123-123-123-1",
		CategoryID: 0,
		FormatID:   0,
	}
}

func createBookForUpdate() *dto.BookDto {
	return &dto.BookDto{
		Title:      "Test2",
		Isbn:       "123-123-123-2",
		CategoryID: 2,
		FormatID:   2,
	}
}

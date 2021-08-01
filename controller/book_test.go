package controller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/ybkuroki/go-webapp-sample/container"
	"github.com/ybkuroki/go-webapp-sample/model"
	"github.com/ybkuroki/go-webapp-sample/model/dto"
	"github.com/ybkuroki/go-webapp-sample/test"
)

func TestGetBookList(t *testing.T) {
	router, container := test.Prepare()

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

func TestCreateBook(t *testing.T) {
	router, container := test.Prepare()

	book := NewBookController(container)
	router.POST(APIBooks, func(c echo.Context) error { return book.CreateBook(c) })

	param := createDto()
	req := httptest.NewRequest("POST", APIBooks, strings.NewReader(test.ConvertToString(param)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	entity := &model.Book{}
	data, _ := entity.FindByID(container.GetRepository(), 1)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, test.ConvertToString(data), rec.Body.String())
}

func TestUpdateBook(t *testing.T) {
	router, container := test.Prepare()

	book := NewBookController(container)
	router.PUT(APIBooksID, func(c echo.Context) error { return book.UpdateBook(c) })

	setUpTestData(container)

	param := changeDto()
	uri := test.NewRequestBuilder().URL(APIBooks).PathParams("1").Build().GetRequestURL()
	req := httptest.NewRequest("PUT", uri, strings.NewReader(test.ConvertToString(param)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	entity := &model.Book{}
	data, _ := entity.FindByID(container.GetRepository(), 1)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, test.ConvertToString(data), rec.Body.String())
}

func TestDeleteBook(t *testing.T) {
	router, container := test.Prepare()

	book := NewBookController(container)
	router.DELETE(APIBooksID, func(c echo.Context) error { return book.DeleteBook(c) })

	setUpTestData(container)

	entity := &model.Book{}
	data, _ := entity.FindByID(container.GetRepository(), 1)

	uri := test.NewRequestBuilder().URL(APIBooks).PathParams("1").Build().GetRequestURL()
	req := httptest.NewRequest("DELETE", uri, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
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

func createDto() *dto.BookDto {
	dto := &dto.BookDto{
		Title:      "Test1",
		Isbn:       "123-123-123-1",
		CategoryID: 1,
		FormatID:   1,
	}
	return dto
}

func changeDto() *dto.BookDto {
	dto := &dto.BookDto{
		Title:      "Test2",
		Isbn:       "123-123-123-2",
		CategoryID: 2,
		FormatID:   2,
	}
	return dto
}

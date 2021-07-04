package controller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/ybkuroki/go-webapp-sample/model"
	"github.com/ybkuroki/go-webapp-sample/model/dto"
	"github.com/ybkuroki/go-webapp-sample/mycontext"
	"github.com/ybkuroki/go-webapp-sample/test"
)

func TestGetBookSearch(t *testing.T) {
	router, context := test.Prepare()

	book := NewBookController(context)
	router.GET(APIBooks, func(c echo.Context) error { return book.GetBookList(c) })

	setUpTestData(context)

	uri := test.NewRequestBuilder().URL(APIBooks).RequestParams("query", "Test").RequestParams("page", "0").RequestParams("size", "5").Build().GetRequestURL()
	req := httptest.NewRequest("GET", uri, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	entity := &model.Book{}
	data, _ := entity.FindByTitle(context.GetRepository(), "Test", "0", "5")

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, test.ConvertToString(data), rec.Body.String())
}

func TestPostBookRegist(t *testing.T) {
	router, context := test.Prepare()

	book := NewBookController(context)
	router.POST(APIBooks, func(c echo.Context) error { return book.CreateBook(c) })

	param := createDto()
	req := httptest.NewRequest("POST", APIBooks, strings.NewReader(test.ConvertToString(param)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	entity := &model.Book{}
	data, _ := entity.FindByID(context.GetRepository(), 1)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, test.ConvertToString(data), rec.Body.String())
}

func TestPostBookEdit(t *testing.T) {
	router, context := test.Prepare()

	book := NewBookController(context)
	router.PUT(APIBooksID, func(c echo.Context) error { return book.UpdateBook(c) })

	setUpTestData(context)

	param := changeDto()
	uri := test.NewRequestBuilder().URL(APIBooks).PathParams("1").Build().GetRequestURL()
	req := httptest.NewRequest("PUT", uri, strings.NewReader(test.ConvertToString(param)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	entity := &model.Book{}
	data, _ := entity.FindByID(context.GetRepository(), 1)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, test.ConvertToString(data), rec.Body.String())
}

func TestPostBookDelete(t *testing.T) {
	router, context := test.Prepare()

	book := NewBookController(context)
	router.DELETE(APIBooksID, func(c echo.Context) error { return book.DeleteBook(c) })

	setUpTestData(context)

	entity := &model.Book{}
	data, _ := entity.FindByID(context.GetRepository(), 1)

	uri := test.NewRequestBuilder().URL(APIBooks).PathParams("1").Build().GetRequestURL()
	req := httptest.NewRequest("DELETE", uri, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, test.ConvertToString(data), rec.Body.String())
}

func setUpTestData(context mycontext.Context) {
	model := model.NewBook("Test1", "123-123-123-1", 1, 1)
	repo := context.GetRepository()
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

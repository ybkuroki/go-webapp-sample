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

func TestGetBookList(t *testing.T) {
	router, context := test.Prepare()

	book := NewBookController(context)
	router.GET(APIBookList, func(c echo.Context) error { return book.GetBookList(c) })

	setUpTestData(context)

	uri := test.NewRequestBuilder().URL(APIBookList).Params("page", "0").Params("size", "5").Build().GetRequestURL()
	req := httptest.NewRequest("GET", uri, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	entity := &model.Book{}
	data, _ := entity.FindAllByPage(context.GetRepository(), 0, 5)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, test.ConvertToString(data), rec.Body.String())
}

func TestGetBookSearch(t *testing.T) {
	router, context := test.Prepare()

	book := NewBookController(context)
	router.GET(APIBookSearch, func(c echo.Context) error { return book.GetBookSearch(c) })

	setUpTestData(context)

	uri := test.NewRequestBuilder().URL(APIBookSearch).Params("query", "Test").Params("page", "0").Params("size", "5").Build().GetRequestURL()
	req := httptest.NewRequest("GET", uri, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	entity := &model.Book{}
	data, _ := entity.FindByTitle(context.GetRepository(), "Test", 0, 5)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, test.ConvertToString(data), rec.Body.String())
}

func TestPostBookRegist(t *testing.T) {
	router, context := test.Prepare()

	book := NewBookController(context)
	router.POST(APIBookRegist, func(c echo.Context) error { return book.PostBookRegist(c) })

	param := createRegDto()
	req := httptest.NewRequest("POST", APIBookRegist, strings.NewReader(test.ConvertToString(param)))
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
	router.POST(APIBookEdit, func(c echo.Context) error { return book.PostBookEdit(c) })

	setUpTestData(context)

	param := createChgDto()
	req := httptest.NewRequest("POST", APIBookEdit, strings.NewReader(test.ConvertToString(param)))
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
	router.POST(APIBookDelete, func(c echo.Context) error { return book.PostBookDelete(c) })

	setUpTestData(context)

	entity := &model.Book{}
	data, _ := entity.FindByID(context.GetRepository(), 1)

	param := createChgDto()
	req := httptest.NewRequest("POST", APIBookDelete, strings.NewReader(test.ConvertToString(param)))
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

func createRegDto() *dto.RegBookDto {
	dto := &dto.RegBookDto{
		Title:      "Test1",
		Isbn:       "123-123-123-1",
		CategoryID: 1,
		FormatID:   1,
	}
	return dto
}

func createChgDto() *dto.ChgBookDto {
	dto := &dto.ChgBookDto{
		ID:         1,
		Title:      "EditedTest1",
		Isbn:       "234-234-234-2",
		CategoryID: 2,
		FormatID:   2,
	}
	return dto
}

package controller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ybkuroki/go-webapp-sample/model"
	"github.com/ybkuroki/go-webapp-sample/model/dto"
	"github.com/ybkuroki/go-webapp-sample/repository"
	"github.com/ybkuroki/go-webapp-sample/test"
)

func TestGetBookList(t *testing.T) {
	router := test.Prepare()
	router.GET(APIBookList, GetBookList())

	setUpTestData()

	uri := test.NewRequestBuilder().URL(APIBookList).Params("page", "0").Params("size", "5").Build().GetRequestURL()
	req := httptest.NewRequest("GET", uri, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	book := &model.Book{}
	data, _ := book.FindAllByPage(repository.GetRepository(), 0, 5)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, test.ConvertToString(data), rec.Body.String())
}

func TestGetBookSearch(t *testing.T) {
	router := test.Prepare()
	router.GET(APIBookSearch, GetBookSearch())

	setUpTestData()

	uri := test.NewRequestBuilder().URL(APIBookSearch).Params("query", "Test").Params("page", "0").Params("size", "5").Build().GetRequestURL()
	req := httptest.NewRequest("GET", uri, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	book := &model.Book{}
	data, _ := book.FindByTitle(repository.GetRepository(), "Test", 0, 5)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, test.ConvertToString(data), rec.Body.String())
}

func TestPostBookRegist(t *testing.T) {
	router := test.Prepare()
	router.POST(APIBookRegist, PostBookRegist())

	param := createRegDto()
	req := httptest.NewRequest("POST", APIBookRegist, strings.NewReader(test.ConvertToString(param)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	book := &model.Book{}
	data, _ := book.FindByID(repository.GetRepository(), 1)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, test.ConvertToString(data), rec.Body.String())
}

func TestPostBookEdit(t *testing.T) {
	router := test.Prepare()
	router.POST(APIBookEdit, PostBookEdit())

	setUpTestData()

	param := createChgDto()
	req := httptest.NewRequest("POST", APIBookEdit, strings.NewReader(test.ConvertToString(param)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	book := &model.Book{}
	data, _ := book.FindByID(repository.GetRepository(), 1)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, test.ConvertToString(data), rec.Body.String())
}

func TestPostBookDelete(t *testing.T) {
	router := test.Prepare()
	router.POST(APIBookDelete, PostBookDelete())

	setUpTestData()

	book := &model.Book{}
	data, _ := book.FindByID(repository.GetRepository(), 1)

	param := createChgDto()
	req := httptest.NewRequest("POST", APIBookDelete, strings.NewReader(test.ConvertToString(param)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, test.ConvertToString(data), rec.Body.String())
}

func setUpTestData() {
	model := &model.Book{
		ID:         1,
		Title:      "Test1",
		Isbn:       "123-123-123-1",
		CategoryID: 1,
		Category:   &model.Category{ID: 1, Name: "技術書"},
		FormatID:   1,
		Format:     &model.Format{ID: 1, Name: "書籍"},
	}
	repo := repository.GetRepository()
	_, _ = model.Save(repo)
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

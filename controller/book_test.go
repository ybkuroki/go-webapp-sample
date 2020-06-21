package controller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ybkuroki/go-webapp-sample/model"
	"github.com/ybkuroki/go-webapp-sample/model/dto"
	"github.com/ybkuroki/go-webapp-sample/test"
)

func TestPostBookRegist(t *testing.T) {
	router := test.Prepare()
	router.POST("/api/book/new", PostBookRegist())

	param := &dto.RegBookDto{
		Title:      "Test1",
		Isbn:       "123-123-123-1",
		CategoryID: 1,
		FormatID:   1,
	}

	req := httptest.NewRequest("POST", "/api/book/new", strings.NewReader(test.ConvertToString(param)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	data := &model.Book{
		ID:         1,
		Title:      "Test1",
		Isbn:       "123-123-123-1",
		CategoryID: 1,
		Category:   &model.Category{1, "技術書"},
		FormatID:   1,
		Format:     &model.Format{1, "書籍"},
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, test.ConvertToString(data), rec.Body.String())
}

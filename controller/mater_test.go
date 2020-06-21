package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ybkuroki/go-webapp-sample/model"
	"github.com/ybkuroki/go-webapp-sample/test"
)

func TestGetCategoryList(t *testing.T) {
	router := test.Prepare()
	router.GET("/api/master/category", GetCategoryList())

	req := httptest.NewRequest("GET", "/api/master/category", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	data := [...]*model.Category{
		&model.Category{ID: 1, Name: "技術書"},
		&model.Category{ID: 2, Name: "雑誌"},
		&model.Category{ID: 3, Name: "小説"},
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, test.ConvertToString(data), rec.Body.String())
}

func TestGetFormatList(t *testing.T) {
	router := test.Prepare()
	router.GET("/api/master/format", GetFormatList())

	req := httptest.NewRequest("GET", "/api/master/format", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	data := [...]*model.Format{
		&model.Format{ID: 1, Name: "書籍"},
		&model.Format{ID: 2, Name: "電子書籍"},
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, test.ConvertToString(data), rec.Body.String())
}

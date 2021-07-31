package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/ybkuroki/go-webapp-sample/model"
	"github.com/ybkuroki/go-webapp-sample/test"
)

func TestGetCategoryList(t *testing.T) {
	router, container := test.Prepare()

	category := NewCategoryController(container)
	router.GET(APICategories, func(c echo.Context) error { return category.GetCategoryList(c) })

	req := httptest.NewRequest("GET", APICategories, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	data := [...]*model.Category{
		{ID: 1, Name: "技術書"},
		{ID: 2, Name: "雑誌"},
		{ID: 3, Name: "小説"},
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, test.ConvertToString(data), rec.Body.String())
}

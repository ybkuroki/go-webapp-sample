package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/maju6406/go-webapp-sample/model"
	"github.com/maju6406/go-webapp-sample/test"
	"github.com/stretchr/testify/assert"
)

func TestGetCategoryList(t *testing.T) {
	router, context := test.Prepare()

	master := NewMasterController(context)
	router.GET(APIMasterCategory, func(c echo.Context) error { return master.GetCategoryList(c) })

	req := httptest.NewRequest("GET", APIMasterCategory, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	data := [...]*model.Category{
		{ID: 1, Name: "Technical book"},
		{ID: 2, Name: "Magazine"},
		{ID: 3, Name: "Novel"},
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, test.ConvertToString(data), rec.Body.String())
}

func TestGetFormatList(t *testing.T) {
	router, context := test.Prepare()

	master := NewMasterController(context)
	router.GET(APIMasterFormat, func(c echo.Context) error { return master.GetFormatList(c) })

	req := httptest.NewRequest("GET", APIMasterFormat, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	data := [...]*model.Format{
		{ID: 1, Name: "Book"},
		{ID: 2, Name: "E-Book"},
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, test.ConvertToString(data), rec.Body.String())
}

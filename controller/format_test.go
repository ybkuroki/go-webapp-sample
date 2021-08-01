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

func TestGetFormatList(t *testing.T) {
	router, container := test.Prepare()

	format := NewFormatController(container)
	router.GET(APIFormats, func(c echo.Context) error { return format.GetFormatList(c) })

	req := httptest.NewRequest("GET", APIFormats, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	data := [...]*model.Format{
		{ID: 1, Name: "Paper Book"},
		{ID: 2, Name: "e-Book"},
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, test.ConvertToString(data), rec.Body.String())
}

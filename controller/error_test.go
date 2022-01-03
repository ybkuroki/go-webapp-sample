package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ybkuroki/go-webapp-sample/test"
)

func TestJSONError(t *testing.T) {
	router, container := test.PrepareForControllerTest(false)

	errorHandler := NewErrorController(container)
	router.HTTPErrorHandler = errorHandler.JSONError

	req := httptest.NewRequest("GET", "/api/movies/1", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
	assert.JSONEq(t, `{"Code":404,"Message":"Not Found"}`, rec.Body.String())
}

package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHealthCheck(t *testing.T) {
	router := Prepare()

	req := httptest.NewRequest("GET", "/api/health", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `"healthy"`, rec.Body.String())
}

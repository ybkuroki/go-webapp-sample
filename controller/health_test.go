package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ybkuroki/go-webapp-sample/test"
)

func TestGetHealthCheck(t *testing.T) {
	router, _ := test.Prepare()
	router.GET(APIHealth, GetHealthCheck())

	req := httptest.NewRequest("GET", APIHealth, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `"healthy"`, rec.Body.String())
}

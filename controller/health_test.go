package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/ybkuroki/go-webapp-sample/config"
	"github.com/ybkuroki/go-webapp-sample/test"
)

func TestGetHealthCheck(t *testing.T) {
	router, container := test.PrepareForControllerTest(false)

	health := NewHealthController(container)
	router.GET(config.APIHealth, func(c echo.Context) error { return health.GetHealthCheck(c) })

	req := httptest.NewRequest("GET", config.APIHealth, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `"healthy"`, rec.Body.String())
}

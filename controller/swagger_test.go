package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/ybkuroki/go-webapp-sample/docs" // for using echo-swagger
	"github.com/ybkuroki/go-webapp-sample/test"
)

func TestSwagger(t *testing.T) {
	router, _ := test.PrepareForControllerTest(false)
	router.GET("/swagger/*", echoSwagger.WrapHandler)

	req := httptest.NewRequest("GET", "/swagger/index.html", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Regexp(t, "Swagger UI", rec.Body.String())
}

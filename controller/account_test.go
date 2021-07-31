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

func TestGetLoginStatus(t *testing.T) {
	router, container := test.Prepare()

	account := NewAccountController(container)
	router.GET(APIAccountLoginStatus, func(c echo.Context) error { return account.GetLoginStatus(c) })

	req := httptest.NewRequest("GET", APIAccountLoginStatus, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, "true", rec.Body.String())
}

func TestGetLoginAccount(t *testing.T) {
	router, container := test.Prepare()

	account := NewAccountController(container)
	router.GET(APIAccountLoginAccount, func(c echo.Context) error { return account.GetLoginAccount(c) })

	req := httptest.NewRequest("GET", APIAccountLoginAccount, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	entity := model.NewAccountWithPlainPassword("test", "test", 1)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, test.ConvertToString(entity), rec.Body.String())
}

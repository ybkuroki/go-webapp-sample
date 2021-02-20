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

func TestGetLoginStatus(t *testing.T) {
	router, context := test.Prepare()

	account := NewAccountController(context)
	router.GET(APIAccountLoginStatus, func(c echo.Context) error { return account.GetLoginStatus(c) })

	req := httptest.NewRequest("GET", APIAccountLoginStatus, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, "true", rec.Body.String())
}

func TestGetLoginAccount(t *testing.T) {
	router, context := test.Prepare()

	account := NewAccountController(context)
	router.GET(APIAccountLoginAccount, func(c echo.Context) error { return account.GetLoginAccount(c) })

	req := httptest.NewRequest("GET", APIAccountLoginAccount, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	entity := model.NewAccountWithPlainPassword("test", "test", 1)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, test.ConvertToString(entity), rec.Body.String())
}

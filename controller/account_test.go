package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/ybkuroki/go-webapp-sample/model"
	"github.com/ybkuroki/go-webapp-sample/model/dto"
	"github.com/ybkuroki/go-webapp-sample/test"
)

func TestGetLoginStatus_Success(t *testing.T) {
	router, container := test.PrepareForControllerTest(false)

	account := NewAccountController(container)
	router.GET(APIAccountLoginStatus, func(c echo.Context) error { return account.GetLoginStatus(c) })

	req := httptest.NewRequest("GET", APIAccountLoginStatus, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, "true", rec.Body.String())
}

func TestGetLoginAccount_Success(t *testing.T) {
	router, container := test.PrepareForControllerTest(false)

	account := NewAccountController(container)
	router.GET(APIAccountLoginAccount, func(c echo.Context) error { return account.GetLoginAccount(c) })

	req := httptest.NewRequest("GET", APIAccountLoginAccount, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	entity := model.NewAccountWithPlainPassword("test", "test", 1)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, test.ConvertToString(entity), rec.Body.String())
}

func TestLogin_Success(t *testing.T) {
	router, container := test.PrepareForControllerTest(true)

	account := NewAccountController(container)
	router.POST(APIAccountLogin, func(c echo.Context) error { return account.Login(c) })

	param := createLoginSuccessAccount()
	req := test.NewJsonRequest("POST", APIAccountLogin, param)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEmpty(t, test.GetCookie(rec, "GSESSION"))
}

func TestLogin_AuthenticationFailure(t *testing.T) {
	router, container := test.PrepareForControllerTest(true)

	account := NewAccountController(container)
	router.POST(APIAccountLogin, func(c echo.Context) error { return account.Login(c) })

	param := createLoginFailureAccount()
	req := test.NewJsonRequest("POST", APIAccountLogin, param)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.Empty(t, test.GetCookie(rec, "GSESSION"))
}

func TestLogout_Success(t *testing.T) {
	router, container := test.PrepareForControllerTest(true)

	account := NewAccountController(container)
	router.POST(APIAccountLogout, func(c echo.Context) error { return account.Logout(c) })

	req := test.NewJsonRequest("POST", APIAccountLogout, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEmpty(t, test.GetCookie(rec, "GSESSION"))
}

func createLoginSuccessAccount() *dto.LoginDto {
	return &dto.LoginDto{
		UserName: "test",
		Password: "test",
	}
}

func createLoginFailureAccount() *dto.LoginDto {
	return &dto.LoginDto{
		UserName: "test",
		Password: "abcde",
	}
}

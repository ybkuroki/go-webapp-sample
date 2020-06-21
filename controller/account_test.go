package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ybkuroki/go-webapp-sample/test"
)

func TestGetLoginStatus(t *testing.T) {
	router := test.Prepare()
	router.GET("/api/account/loginStatus", GetLoginStatus())

	req := httptest.NewRequest("GET", "/api/account/loginStatus", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, "true", rec.Body.String())
}

func TestGetLoginAccount(t *testing.T) {
	router := test.Prepare()
	router.GET("/api/account/loginAccount", GetLoginAccount())

	req := httptest.NewRequest("GET", "/api/account/loginAccount", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, test.ConvertToString(&Account{ID: 1, Name: "test"}), rec.Body.String())
}

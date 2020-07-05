package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/model"
	session "github.com/ybkuroki/go-webapp-sample/session"
)

// GetLoginStatus is
func GetLoginStatus() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, true)
	}
}

// GetLoginAccount is
func GetLoginAccount() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, &model.Account{ID: 1, Name: "test"})
	}
}

// PostLogin is
func PostLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
		account := session.GetAccount(c)
		if session.GetAccount(c) == nil {
			a := &model.Account{ID: 1, Name: "test"}
			_ = session.SetAccount(c, a)
			_ = session.Save(c)
			return c.JSON(http.StatusOK, a)
		}
		return c.JSON(http.StatusOK, account)
	}
}

// PostLogout is
func PostLogout() echo.HandlerFunc {
	return func(c echo.Context) error {
		_ = session.SetAccount(c, nil)
		_ = session.Save(c)
		return c.NoContent(http.StatusOK)
	}
}

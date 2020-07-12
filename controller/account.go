package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/service"
	"github.com/ybkuroki/go-webapp-sample/session"
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
		return c.JSON(http.StatusOK, session.GetAccount(c))
	}
}

// PostLogin is
func PostLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		account := session.GetAccount(c)
		if account == nil {
			authenticate, a := service.AuthenticateByUsernameAndPassword(username, password)
			if authenticate == true {
				_ = session.SetAccount(c, a)
				_ = session.Save(c)
				return c.JSON(http.StatusOK, a)
			}
			return c.NoContent(http.StatusUnauthorized)
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

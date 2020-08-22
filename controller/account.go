package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/config"
	"github.com/ybkuroki/go-webapp-sample/model"
	"github.com/ybkuroki/go-webapp-sample/service"
	"github.com/ybkuroki/go-webapp-sample/session"
)

var dummyAccount = model.NewAccountWithPlainPassword("test", "test", model.NewAuthority("Admin"))

// GetLoginStatus returns the status of login.
func GetLoginStatus() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, true)
	}
}

// GetLoginAccount returns the account data of logged in user.
func GetLoginAccount() echo.HandlerFunc {
	return func(c echo.Context) error {
		if !config.GetConfig().Extension.SecurityEnabled {
			return c.JSON(http.StatusOK, dummyAccount)
		}
		return c.JSON(http.StatusOK, session.GetAccount(c))
	}
}

// PostLogin is the method to login using username and password by http post.
func PostLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		account := session.GetAccount(c)
		if account == nil {
			authenticate, a := service.AuthenticateByUsernameAndPassword(username, password)
			if authenticate {
				_ = session.SetAccount(c, a)
				_ = session.Save(c)
				return c.JSON(http.StatusOK, a)
			}
			return c.NoContent(http.StatusUnauthorized)
		}
		return c.JSON(http.StatusOK, account)
	}
}

// PostLogout is the method to logout by http post.
func PostLogout() echo.HandlerFunc {
	return func(c echo.Context) error {
		_ = session.SetAccount(c, nil)
		_ = session.Delete(c)
		return c.NoContent(http.StatusOK)
	}
}

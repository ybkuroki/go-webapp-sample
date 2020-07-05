package auth

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/config"
	mysession "github.com/ybkuroki/go-webapp-sample/session"
)

// Init initalize session authentication.
// go get -u github.com/ipfans/echo-session@master
func Init(e *echo.Echo, conf *config.Config) {
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.Use(SessionAuthenticationMiddleware)
}

// SessionAuthenticationMiddleware is
func SessionAuthenticationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Path() != "/api/account/login" {
			account := mysession.GetAccount(c)
			if account == nil {
				return c.JSON(http.StatusUnauthorized, false)
			}
			_ = mysession.Save(c)
		}
		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}

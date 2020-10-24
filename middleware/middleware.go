package middleware

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/config"
	"github.com/ybkuroki/go-webapp-sample/logger"
	mySession "github.com/ybkuroki/go-webapp-sample/session"
	"gopkg.in/boj/redistore.v1"
)

// InitLoggerMiddleware initialize a middleware for logger.
func InitLoggerMiddleware(e *echo.Echo) {
	e.Use(RequestLoggerMiddleware)
	e.Use(ActionLoggerMiddleware)
}

// InitSessionMiddleware initialize a middleware for session management.
func InitSessionMiddleware(e *echo.Echo, conf *config.Config) {
	if conf.Extension.SecurityEnabled {
		if conf.Redis.Enabled {
			logger.GetZapLogger().Infof("Try redis connection")
			address := fmt.Sprintf("%s:%s", conf.Redis.Host, conf.Redis.Port)
			store, err := redistore.NewRediStore(conf.Redis.ConnectionPoolSize, "tcp", address, "", []byte("secret"))
			if err != nil {
				logger.GetZapLogger().Errorf("Failure redis connection")
			}
			e.Use(session.Middleware(store))
			logger.GetZapLogger().Infof(fmt.Sprintf("Success redis connection, %s", address))
		} else {
			e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
		}
		e.Use(AuthenticationMiddleware(conf))
	}
}

// RequestLoggerMiddleware is middleware for logging the contents of requests.
func RequestLoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		res := c.Response()
		if err := next(c); err != nil {
			c.Error(err)
		}
		if account := mySession.GetAccount(c); account != nil {
			logger.GetZapLogger().Infof("%s %s %s %d", account.Name, req.RequestURI, req.Method, res.Status)
		} else {
			logger.GetZapLogger().Infof("%s %s %s %d", "None", req.RequestURI, req.Method, res.Status)
		}
		return nil
	}
}

// ActionLoggerMiddleware is middleware for logging the start and end of controller processes.
// ref: https://echo.labstack.com/cookbook/middleware
func ActionLoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		logger.GetZapLogger().Debug(c.Path() + " Action Start")
		if err := next(c); err != nil {
			c.Error(err)
		}
		logger.GetZapLogger().Debugf(c.Path() + " Action End")
		return nil
	}
}

// AuthenticationMiddleware is the middleware of session authentication for echo.
func AuthenticationMiddleware(conf *config.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if !hasAuthorization(c, conf) {
				return c.JSON(http.StatusUnauthorized, false)
			}
			if err := next(c); err != nil {
				c.Error(err)
			}
			return nil
		}
	}
}

// hasAuthorization judges whether the user has the right to access the path.
func hasAuthorization(c echo.Context, conf *config.Config) bool {
	currentPath := c.Path()
	if equalPath(currentPath, conf.Security.ExculdePath) {
		return true
	}
	account := mySession.GetAccount(c)
	if account == nil {
		return false
	}
	if account.Authority.Name == "Admin" && equalPath(currentPath, conf.Security.AdminPath) {
		_ = mySession.Save(c)
		return true
	}
	if account.Authority.Name == "User" && equalPath(currentPath, conf.Security.UserPath) {
		_ = mySession.Save(c)
		return true
	}
	return false
}

// equalPath judges whether a given path contains in the path list.
func equalPath(cpath string, paths []string) bool {
	for i := range paths {
		if regexp.MustCompile(paths[i]).Match([]byte(cpath)) {
			return true
		}
	}
	return false
}

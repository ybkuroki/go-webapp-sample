package logger

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

const (
	logFormat = "${time_rfc3339} [${level}] ${host} ${method} ${uri} ${status}"
)

// InitLogger initialize logger.
func InitLogger(e *echo.Echo) *echo.Echo {
	// logging for each request.
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: strings.Replace(logFormat, "${level}", "INFO", 1) + "\n",
	}))
	// logging for the start and end of controller processes.
	// ref: https://echo.labstack.com/guide/customization
	e.Use(MyLoggerMiddleware)

	// set logformat for echo logger.
	e.Logger.SetHeader(logFormat)
	e.Logger.SetLevel(log.DEBUG)
	return e
}

// MyLoggerMiddleware is middleware for logging the start and end of controller processes.
// ref: https://echo.labstack.com/cookbook/middleware
func MyLoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Logger().Info("Controller Action Start")
		if err := next(c); err != nil {
			c.Error(err)
		}
		c.Logger().Info("Controller Action End")
		return nil
	}
}
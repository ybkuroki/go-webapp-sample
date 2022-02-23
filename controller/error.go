package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/container"
)

// APIError has a error code and a message.
type APIError struct {
	Code    int
	Message string
}

// ErrorController is a controller for handling errors.
type ErrorController interface {
	JSONError(err error, c echo.Context)
}

type errorController struct {
	container container.Container
}

// NewErrorController is constructor.
func NewErrorController(container container.Container) ErrorController {
	return &errorController{container: container}
}

// JSONError is cumstomize error handler
func (controller *errorController) JSONError(err error, c echo.Context) {
	logger := controller.container.GetLogger()
	code := http.StatusInternalServerError
	msg := http.StatusText(code)

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message.(string)
	}

	var apierr APIError
	apierr.Code = code
	apierr.Message = msg

	if !c.Response().Committed {
		if reserr := c.JSON(code, apierr); reserr != nil {
			logger.GetZapLogger().Errorf(reserr.Error())
		}
	}
	logger.GetZapLogger().Debugf(err.Error())
}

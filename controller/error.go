package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/mycontext"
)

// APIError represents
type APIError struct {
	Code    int
	Message string
}

// ErrorController is
type ErrorController struct {
	context mycontext.Context
}

// NewErrorController is
func NewErrorController(context mycontext.Context) *ErrorController {
	return &ErrorController{context: context}
}

// JSONError is cumstomize error handler
func (controller *ErrorController) JSONError(err error, c echo.Context) {
	logger := controller.context.GetLogger()
	code := http.StatusInternalServerError
	msg := http.StatusText(code)

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message.(string)
	}
	if c.Echo().Debug {
		msg = err.Error()
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

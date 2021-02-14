package controller

import (
	"net/http"
	"regexp"

	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/mycontext"
)

// APIError has a error code and a message.
type APIError struct {
	Code    int
	Message string
}

// ErrorController is a controller for handling errors.
type ErrorController struct {
	context mycontext.Context
}

// NewErrorController is constructor.
func NewErrorController(context mycontext.Context) *ErrorController {
	return &ErrorController{context: context}
}

// JSONError is cumstomize error handler
func (controller *ErrorController) JSONError(err error, c echo.Context) {
	logger := controller.context.GetLogger()
	conf := controller.context.GetConfig()
	code := http.StatusInternalServerError
	msg := http.StatusText(code)

	// If a request path is resource path (not api request), redirect to the root path.
	if !regexp.MustCompile(conf.Security.AuthPath[0]).Match([]byte(c.Path())) {
		c.Redirect(http.StatusMovedPermanently, "/")
	}
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

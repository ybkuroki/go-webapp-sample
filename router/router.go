package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ybkuroki/go-webapp-sample/controller"
	"github.com/ybkuroki/go-webapp-sample/mycontext"
)

// Init initialize the routing of this application.
func Init(e *echo.Echo, context mycontext.Context) {
	conf := context.GetConfig()
	if conf.Extension.CorsEnabled {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowCredentials: true,
			AllowOrigins:     []string{"*"},
			AllowHeaders: []string{
				echo.HeaderAccessControlAllowHeaders,
				echo.HeaderContentType,
				echo.HeaderContentLength,
				echo.HeaderAcceptEncoding,
				echo.HeaderXCSRFToken,
				echo.HeaderAuthorization,
				"X-XSRF-TOKEN",
			},
			AllowMethods: []string{
				http.MethodGet,
				http.MethodPost,
			},
			MaxAge: 86400,
		}))
	}

	errorHandler := controller.NewErrorController(context)
	e.HTTPErrorHandler = errorHandler.JSONError
	e.Use(middleware.Recover())

	book := controller.NewBookController(context)
	master := controller.NewMasterController(context)
	account := controller.NewAccountController(context)
	health := controller.NewHealthController(context)

	e.GET(controller.APIBookGet, func(c echo.Context) error { return book.GetBook(c) })
	e.GET(controller.APIBookList, func(c echo.Context) error { return book.GetBookList(c) })
	e.GET(controller.APIBookSearch, func(c echo.Context) error { return book.GetBookSearch(c) })
	e.POST(controller.APIBookRegist, func(c echo.Context) error { return book.PostBookRegist(c) })
	e.POST(controller.APIBookEdit, func(c echo.Context) error { return book.PostBookEdit(c) })
	e.POST(controller.APIBookDelete, func(c echo.Context) error { return book.PostBookDelete(c) })

	e.GET(controller.APIMasterCategory, func(c echo.Context) error { return master.GetCategoryList(c) })
	e.GET(controller.APIMasterFormat, func(c echo.Context) error { return master.GetFormatList(c) })

	e.GET(controller.APIAccountLoginStatus, func(c echo.Context) error { return account.GetLoginStatus(c) })
	e.GET(controller.APIAccountLoginAccount, func(c echo.Context) error { return account.GetLoginAccount(c) })

	if conf.Extension.SecurityEnabled {
		e.POST(controller.APIAccountLogin, func(c echo.Context) error { return account.PostLogin(c) })
		e.POST(controller.APIAccountLogout, func(c echo.Context) error { return account.PostLogout(c) })
	}

	e.GET(controller.APIHealth, func(c echo.Context) error { return health.GetHealthCheck(c) })

}

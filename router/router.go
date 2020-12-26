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

	e.HTTPErrorHandler = controller.JSONErrorHandler
	e.Use(middleware.Recover())

	e.GET(controller.APIBookList, controller.GetBookList(context))
	e.GET(controller.APIBookSearch, controller.GetBookSearch(context))
	e.POST(controller.APIBookRegist, controller.PostBookRegist(context))
	e.POST(controller.APIBookEdit, controller.PostBookEdit(context))
	e.POST(controller.APIBookDelete, controller.PostBookDelete(context))

	e.GET(controller.APIMasterCategory, controller.GetCategoryList(context))
	e.GET(controller.APIMasterFormat, controller.GetFormatList(context))

	e.GET(controller.APIAccountLoginStatus, controller.GetLoginStatus())
	e.GET(controller.APIAccountLoginAccount, controller.GetLoginAccount(context))

	if conf.Extension.SecurityEnabled {
		e.POST(controller.APIAccountLogin, controller.PostLogin(context))
		e.POST(controller.APIAccountLogout, controller.PostLogout())
	}

	e.GET(controller.APIHealth, controller.GetHealthCheck())

}

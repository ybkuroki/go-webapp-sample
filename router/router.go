package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ybkuroki/go-webapp-sample/config"
	"github.com/ybkuroki/go-webapp-sample/controller"
)

// Init initialize the routing of this application.
func Init(e *echo.Echo, conf *config.Config) {
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

	e.GET(controller.APIBookList, controller.GetBookList())
	e.GET(controller.APIBookSearch, controller.GetBookSearch())
	e.POST(controller.APIBookRegist, controller.PostBookRegist())
	e.POST(controller.APIBookEdit, controller.PostBookEdit())
	e.POST(controller.APIBookDelete, controller.PostBookDelete())

	e.GET(controller.APIMasterCategory, controller.GetCategoryList())
	e.GET(controller.APIMasterFormat, controller.GetFormatList())

	e.GET(controller.APIAccountLoginStatus, controller.GetLoginStatus())
	e.GET(controller.APIAccountLoginAccount, controller.GetLoginAccount())

	if conf.Extension.SecurityEnabled {
		e.POST(controller.APIAccountLogin, controller.PostLogin())
		e.POST(controller.APIAccountLogout, controller.PostLogout())
	}

	e.GET(controller.APIHealth, controller.GetHealthCheck())

}

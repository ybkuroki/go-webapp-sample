package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ybkuroki/go-webapp-sample/config"
	"github.com/ybkuroki/go-webapp-sample/controller"
	"github.com/ybkuroki/go-webapp-sample/mycontext"

	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/ybkuroki/go-webapp-sample/docs"
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
			},
			AllowMethods: []string{
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodDelete,
			},
			MaxAge: 86400,
		}))
	}

	errorHandler := controller.NewErrorController(context)
	e.HTTPErrorHandler = errorHandler.JSONError
	e.Use(middleware.Recover())

	book := controller.NewBookController(context)
	category := controller.NewCategoryController(context)
	format := controller.NewFormatController(context)
	account := controller.NewAccountController(context)
	health := controller.NewHealthController(context)

	e.GET(controller.APIBooksID, func(c echo.Context) error { return book.GetBook(c) })
	e.GET(controller.APIBooks, func(c echo.Context) error { return book.GetBookList(c) })
	e.POST(controller.APIBooks, func(c echo.Context) error { return book.CreateBook(c) })
	e.PUT(controller.APIBooksID, func(c echo.Context) error { return book.UpdateBook(c) })
	e.DELETE(controller.APIBooksID, func(c echo.Context) error { return book.DeleteBook(c) })

	e.GET(controller.APICategories, func(c echo.Context) error { return category.GetCategoryList(c) })

	e.GET(controller.APIFormats, func(c echo.Context) error { return format.GetFormatList(c) })

	e.GET(controller.APIAccountLoginStatus, func(c echo.Context) error { return account.GetLoginStatus(c) })
	e.GET(controller.APIAccountLoginAccount, func(c echo.Context) error { return account.GetLoginAccount(c) })

	if conf.Extension.SecurityEnabled {
		e.POST(controller.APIAccountLogin, func(c echo.Context) error { return account.Login(c) })
		e.POST(controller.APIAccountLogout, func(c echo.Context) error { return account.Logout(c) })
	}

	if context.GetEnv() == config.DEV {
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}

	e.GET(controller.APIHealth, func(c echo.Context) error { return health.GetHealthCheck(c) })

}

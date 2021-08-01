package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ybkuroki/go-webapp-sample/config"
	"github.com/ybkuroki/go-webapp-sample/container"
	"github.com/ybkuroki/go-webapp-sample/controller"

	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/ybkuroki/go-webapp-sample/docs" // for using echo-swagger
)

// Init initialize the routing of this application.
func Init(e *echo.Echo, container container.Container) {
	conf := container.GetConfig()
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

	errorHandler := controller.NewErrorController(container)
	e.HTTPErrorHandler = errorHandler.JSONError
	e.Use(middleware.Recover())

	book := controller.NewBookController(container)
	category := controller.NewCategoryController(container)
	format := controller.NewFormatController(container)
	account := controller.NewAccountController(container)
	health := controller.NewHealthController(container)

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

	if container.GetEnv() == config.DEV {
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}

	e.GET(controller.APIHealth, func(c echo.Context) error { return health.GetHealthCheck(c) })

}

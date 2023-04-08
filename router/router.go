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
	setCORSConfig(e, container)

	setErrorController(e, container)
	setBookController(e, container)
	setCategoryController(e, container)
	setFormatController(e, container)
	setAccountController(e, container)
	setHealthController(e, container)

	setSwagger(container, e)
}

func setCORSConfig(e *echo.Echo, container container.Container) {
	if container.GetConfig().Extension.CorsEnabled {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowCredentials:                         true,
			UnsafeWildcardOriginWithAllowCredentials: true,
			AllowOrigins:                             []string{"*"},
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
}

func setErrorController(e *echo.Echo, container container.Container) {
	errorHandler := controller.NewErrorController(container)
	e.HTTPErrorHandler = errorHandler.JSONError
	e.Use(middleware.Recover())
}

func setBookController(e *echo.Echo, container container.Container) {
	book := controller.NewBookController(container)
	e.GET(config.APIBooksID, func(c echo.Context) error { return book.GetBook(c) })
	e.GET(config.APIBooks, func(c echo.Context) error { return book.GetBookList(c) })
	e.POST(config.APIBooks, func(c echo.Context) error { return book.CreateBook(c) })
	e.PUT(config.APIBooksID, func(c echo.Context) error { return book.UpdateBook(c) })
	e.DELETE(config.APIBooksID, func(c echo.Context) error { return book.DeleteBook(c) })
}

func setCategoryController(e *echo.Echo, container container.Container) {
	category := controller.NewCategoryController(container)
	e.GET(config.APICategories, func(c echo.Context) error { return category.GetCategoryList(c) })
}

func setFormatController(e *echo.Echo, container container.Container) {
	format := controller.NewFormatController(container)
	e.GET(config.APIFormats, func(c echo.Context) error { return format.GetFormatList(c) })
}

func setAccountController(e *echo.Echo, container container.Container) {
	account := controller.NewAccountController(container)
	e.GET(config.APIAccountLoginStatus, func(c echo.Context) error { return account.GetLoginStatus(c) })
	e.GET(config.APIAccountLoginAccount, func(c echo.Context) error { return account.GetLoginAccount(c) })

	if container.GetConfig().Extension.SecurityEnabled {
		e.POST(config.APIAccountLogin, func(c echo.Context) error { return account.Login(c) })
		e.POST(config.APIAccountLogout, func(c echo.Context) error { return account.Logout(c) })
	}
}

func setHealthController(e *echo.Echo, container container.Container) {
	health := controller.NewHealthController(container)
	e.GET(config.APIHealth, func(c echo.Context) error { return health.GetHealthCheck(c) })
}

func setSwagger(container container.Container, e *echo.Echo) {
	if container.GetConfig().Swagger.Enabled {
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}
}

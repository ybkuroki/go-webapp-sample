package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ybkuroki/go-webapp-sample/controller"
)

// Init is
func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
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
	e.Use(middleware.Recover())

	api := e.Group("/api")
	{
		book := api.Group("/book")
		{
			book.GET("/list", controller.GetBookList())
			book.POST("/new", controller.PostBookRegist())
		}

		master := api.Group("/master")
		{
			master.GET("/category", controller.GetCategoryList())
			master.GET("/format", controller.GetFormatList())
		}

		account := api.Group("/account")
		{
			account.GET("/loginStatus", controller.GetLoginStatus())
			account.GET("/loginAccount", controller.GetLoginAccount())
		}
	}

	return e
}

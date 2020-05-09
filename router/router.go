package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ybkuroki/go-webapp-sample/config"
	"github.com/ybkuroki/go-webapp-sample/controller"
)

// Init is
func Init(config *config.Config) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())

	if config.Extension.CorsEnabled {
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

	e.Use(middleware.Recover())

	api := e.Group("/api")
	{
		book := api.Group("/book")
		{
			book.GET("/list", controller.GetBookList())
			book.GET("/search", controller.GetBookSearch())
			book.POST("/new", controller.PostBookRegist())
			book.POST("/edit", controller.PostBookEdit())
			book.POST("/delete", controller.PostBookDelete())
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

		api.GET("/health", controller.GetHealthCheck())
	}

	return e
}

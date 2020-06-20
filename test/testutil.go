package test

import (
	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/config"
	"github.com/ybkuroki/go-webapp-sample/logger"
	"github.com/ybkuroki/go-webapp-sample/migration"
	"github.com/ybkuroki/go-webapp-sample/repository"
	"github.com/ybkuroki/go-webapp-sample/router"
)

// Prepare is to prepare for unit test.
func Prepare() *echo.Echo {
	e := echo.New()

	config.Load()
	config.GetConfig().Database.Host = "file::memory:?cache=shared"
	logger.InitLogger(e, config.GetConfig())
	e.Logger.Info("Loaded this configuration : application." + *config.GetEnv() + ".yml")

	repository.InitDB()

	migration.CreateDatabase(config.GetConfig())
	migration.InitMasterData(config.GetConfig())

	router.Init(e, config.GetConfig())

	return e
}

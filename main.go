package main

import (
	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/config"
	"github.com/ybkuroki/go-webapp-sample/logger"
	"github.com/ybkuroki/go-webapp-sample/migration"
	"github.com/ybkuroki/go-webapp-sample/repository"
	"github.com/ybkuroki/go-webapp-sample/router"
)

func main() {
	e := echo.New()

	logger.InitLogger(e)
	config.Load(e.Logger)

	repository.InitDB(e.Logger)
	db := repository.GetDB()

	migration.CreateDatabase(config.GetConfig())
	migration.InitMasterData(config.GetConfig())

	router.Init(e, config.GetConfig())
	e.Start(":8080")

	defer db.Close()
}

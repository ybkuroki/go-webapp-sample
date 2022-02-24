package main

import (
	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/config"
	"github.com/ybkuroki/go-webapp-sample/container"
	"github.com/ybkuroki/go-webapp-sample/logger"
	"github.com/ybkuroki/go-webapp-sample/middleware"
	"github.com/ybkuroki/go-webapp-sample/migration"
	"github.com/ybkuroki/go-webapp-sample/repository"
	"github.com/ybkuroki/go-webapp-sample/router"
)

// @title go-webapp-sample API
// @version 1.5.1
// @description This is API specification for go-webapp-sample project.

// @license.name MIT
// @license.url https://opensource.org/licenses/mit-license.php

// @host localhost:8080
// @BasePath /api
func main() {
	e := echo.New()

	conf, env := config.Load()
	logger := logger.NewLogger(env)
	logger.GetZapLogger().Infof("Loaded this configuration : application." + env + ".yml")

	rep := repository.NewBookRepository(logger, conf)
	container := container.NewContainer(rep, conf, logger, env)

	migration.CreateDatabase(container)
	migration.InitMasterData(container)

	router.Init(e, container)
	middleware.InitLoggerMiddleware(e, container)
	middleware.InitSessionMiddleware(e, container)

	if conf.StaticContents.Path != "" {
		e.Static("/", conf.StaticContents.Path)
		logger.GetZapLogger().Infof("Served the static contents. path: " + conf.StaticContents.Path)
	}

	if err := e.Start(":8000"); err != nil {
		logger.GetZapLogger().Errorf(err.Error())
	}

	defer rep.Close()
}

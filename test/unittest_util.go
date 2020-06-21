package test

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/config"
	"github.com/ybkuroki/go-webapp-sample/logger"
	"github.com/ybkuroki/go-webapp-sample/migration"
	"github.com/ybkuroki/go-webapp-sample/repository"
)

// Prepare func is to prepare for unit test.
func Prepare() *echo.Echo {
	e := echo.New()

	config.Load()
	config.GetConfig().Database.Host = "file::memory:?cache=shared"
	config.GetConfig().Database.Migration = true
	config.GetConfig().Extension.MasterGenerator = true

	logger.InitLogger(e, config.GetConfig())

	repository.InitDB()

	migration.CreateDatabase(config.GetConfig())
	migration.InitMasterData(config.GetConfig())

	return e
}

// ConvertToString func is convert model to string.
func ConvertToString(model interface{}) string {
	bytes, _ := json.Marshal(model)
	return string(bytes)
}

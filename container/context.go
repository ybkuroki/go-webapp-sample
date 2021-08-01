package container

import (
	"github.com/ybkuroki/go-webapp-sample/config"
	"github.com/ybkuroki/go-webapp-sample/logger"
	"github.com/ybkuroki/go-webapp-sample/repository"
)

// Container represents a interface for accessing the data which sharing in overall application.
type Container interface {
	GetRepository() repository.Repository
	GetConfig() *config.Config
	GetLogger() *logger.Logger
	GetEnv() string
}

// container struct is for sharing data which such as database setting, the setting of application and logger in overall this application.
type container struct {
	rep    repository.Repository
	config *config.Config
	logger *logger.Logger
	env    string
}

// NewContainer is constructor.
func NewContainer(rep repository.Repository, config *config.Config, logger *logger.Logger, env string) Container {
	return &container{rep: rep, config: config, logger: logger, env: env}
}

// GetRepository returns the object of repository.
func (c *container) GetRepository() repository.Repository {
	return c.rep
}

// GetConfig returns the object of configuration.
func (c *container) GetConfig() *config.Config {
	return c.config
}

// GetLogger returns the object of logger.
func (c *container) GetLogger() *logger.Logger {
	return c.logger
}

// GetEnv returns the running environment.
func (c *container) GetEnv() string {
	return c.env
}

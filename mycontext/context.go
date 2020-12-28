package mycontext

import (
	"github.com/ybkuroki/go-webapp-sample/config"
	"github.com/ybkuroki/go-webapp-sample/logger"
	"github.com/ybkuroki/go-webapp-sample/repository"
)

// Context represents a interface for accessing the data which sharing in overall application.
type Context interface {
	GetRepository() repository.Repository
	GetConfig() *config.Config
	GetLogger() *logger.Logger
}

// context struct is for sharing data which such as database setting, the setting of application and logger in overall this application.
type context struct {
	rep    repository.Repository
	config *config.Config
	logger *logger.Logger
}

// NewContext is constructor.
func NewContext(rep repository.Repository, config *config.Config, logger *logger.Logger) Context {
	return &context{rep: rep, config: config, logger: logger}
}

// GetRepository returns the object of repository.
func (c *context) GetRepository() repository.Repository {
	return c.rep
}

// GetConfig returns the object of configuration.
func (c *context) GetConfig() *config.Config {
	return c.config
}

// GetLogger returns the object of logger.
func (c *context) GetLogger() *logger.Logger {
	return c.logger
}

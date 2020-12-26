package mycontext

import (
	"github.com/ybkuroki/go-webapp-sample/config"
	"github.com/ybkuroki/go-webapp-sample/logger"
	"github.com/ybkuroki/go-webapp-sample/repository"
)

// Context is
type Context interface {
	GetRepository() repository.Repository
	GetConfig() *config.Config
	GetLogger() *logger.Logger
}

type context struct {
	rep    repository.Repository
	config *config.Config
	logger *logger.Logger
}

// NewContext is
func NewContext(rep repository.Repository, config *config.Config, logger *logger.Logger) Context {
	return &context{rep: rep, config: config, logger: logger}
}

func (c *context) GetRepository() repository.Repository {
	return c.rep
}

func (c *context) GetConfig() *config.Config {
	return c.config
}

func (c *context) GetLogger() *logger.Logger {
	return c.logger
}

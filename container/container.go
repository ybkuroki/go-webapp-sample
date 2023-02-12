package container

import (
	"github.com/ybkuroki/go-webapp-sample/config"
	"github.com/ybkuroki/go-webapp-sample/logger"
	"github.com/ybkuroki/go-webapp-sample/repository"
	"github.com/ybkuroki/go-webapp-sample/session"
)

// Container represents a interface for accessing the data which sharing in overall application.
type Container interface {
	GetRepository() repository.Repository
	GetSession() session.Session
	GetConfig() *config.Config
	GetMessages() map[string]string
	GetLogger() logger.Logger
	GetEnv() string
}

// container struct is for sharing data which such as database setting, the setting of application and logger in overall this application.
type container struct {
	rep      repository.Repository
	session  session.Session
	config   *config.Config
	messages map[string]string
	logger   logger.Logger
	env      string
}

// NewContainer is constructor.
func NewContainer(rep repository.Repository, s session.Session, config *config.Config,
	messages map[string]string, logger logger.Logger, env string) Container {
	return &container{rep: rep, session: s, config: config,
		messages: messages, logger: logger, env: env}
}

// GetRepository returns the object of repository.
func (c *container) GetRepository() repository.Repository {
	return c.rep
}

// GetSession returns the object of session.
func (c *container) GetSession() session.Session {
	return c.session
}

// GetConfig returns the object of configuration.
func (c *container) GetConfig() *config.Config {
	return c.config
}

// GetMessages returns the map has key and message.
func (c *container) GetMessages() map[string]string {
	return c.messages
}

// GetLogger returns the object of logger.
func (c *container) GetLogger() logger.Logger {
	return c.logger
}

// GetEnv returns the running environment.
func (c *container) GetEnv() string {
	return c.env
}

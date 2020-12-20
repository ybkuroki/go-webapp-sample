package logger

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ybkuroki/go-webapp-sample/config"
	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
	"gopkg.in/yaml.v2"
)

var logger *Logger

// Config is
type Config struct {
	ZapConfig zap.Config        `json:"zap_config" yaml:"zap_config"`
	LogRotate lumberjack.Logger `json:"log_rotate" yaml:"log_rotate"`
}

// Logger is an alternative implementation of *gorm.Logger
type Logger struct {
	zap *zap.SugaredLogger
}

// GetLogger is return Logger
func GetLogger() *Logger {
	return logger
}

// SetLogger sets logger
func SetLogger(log *Logger) {
	logger = log
}

// GetZapLogger returns zapSugaredLogger
func GetZapLogger() *zap.SugaredLogger {
	return logger.zap
}

// NewLogger create logger object for *gorm.DB from *echo.Logger
func NewLogger(zap *zap.SugaredLogger) *Logger {
	return &Logger{zap: zap}
}

// InitLogger initialize logger.
func InitLogger() {
	configYaml, err := ioutil.ReadFile("./zaplogger." + *config.GetEnv() + ".yml")
	if err != nil {
		fmt.Printf("Failed to read logger configuration: %s", err)
		os.Exit(2)
	}
	var myConfig *Config
	if err = yaml.Unmarshal(configYaml, &myConfig); err != nil {
		fmt.Printf("Failed to read zap logger configuration: %s", err)
		os.Exit(2)
	}
	var zap *zap.Logger
	zap, err = build(myConfig)
	if err != nil {
		fmt.Printf("Failed to compose zap logger : %s", err)
		os.Exit(2)
	}
	sugar := zap.Sugar()
	// set package varriable logger.
	logger = NewLogger(sugar)
	logger.zap.Infof("Success to read zap logger configuration: zaplogger." + *config.GetEnv() + ".yml")
	_ = zap.Sync()
}

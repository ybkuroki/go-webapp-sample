package logger

import (
	"context"
	"fmt"
	"time"

	gormLogger "gorm.io/gorm/logger"
	gormUtils "gorm.io/gorm/utils"
)

// Customize SQL Logger for gorm library
// ref: https://github.com/wantedly/gorm-zap
// ref: https://github.com/go-gorm/gorm/blob/master/logger/logger.go

const (
	logTitle      = "[gorm] "
	sqlFormat     = logTitle + "%s"
	messageFormat = logTitle + "%s, %s"
	errorFormat   = logTitle + "%s, %s, %s"
	slowThreshold = 200
)

// LogMode The log level of gorm logger is overwrited by the log level of Zap logger.
func (log *logger) LogMode(_ gormLogger.LogLevel) gormLogger.Interface {
	return log
}

// Info prints a information log.
func (log *logger) Info(_ context.Context, msg string, data ...interface{}) {
	log.Zap.Infof(messageFormat, append([]interface{}{msg, gormUtils.FileWithLineNum()}, data...)...)
}

// Warn prints a warning log.
func (log *logger) Warn(_ context.Context, msg string, data ...interface{}) {
	log.Zap.Warnf(messageFormat, append([]interface{}{msg, gormUtils.FileWithLineNum()}, data...)...)
}

// Error prints a error log.
func (log *logger) Error(_ context.Context, msg string, data ...interface{}) {
	log.Zap.Errorf(messageFormat, append([]interface{}{msg, gormUtils.FileWithLineNum()}, data...)...)
}

// Trace prints a trace log such as sql, source file and error.
func (log *logger) Trace(_ context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)

	switch {
	case err != nil:
		sql, _ := fc()
		log.GetZapLogger().Errorf(errorFormat, gormUtils.FileWithLineNum(), err, sql)
	case elapsed > slowThreshold*time.Millisecond && slowThreshold*time.Millisecond != 0:
		sql, _ := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", slowThreshold)
		log.GetZapLogger().Warnf(errorFormat, gormUtils.FileWithLineNum(), slowLog, sql)
	default:
		sql, _ := fc()
		log.GetZapLogger().Debugf(sqlFormat, sql)
	}
}

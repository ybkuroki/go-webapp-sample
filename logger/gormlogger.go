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

// LogMode: The log level of gorm logger is overwrited by the log level of Zap logger.
func (l *logger) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	return l
}

// Info prints a information log.
func (l *logger) Info(ctx context.Context, msg string, data ...interface{}) {
	l.Zap.Infof(messageFormat, append([]interface{}{msg, gormUtils.FileWithLineNum()}, data...)...)
}

// Warn prints a warning log.
func (l *logger) Warn(ctx context.Context, msg string, data ...interface{}) {
	l.Zap.Warnf(messageFormat, append([]interface{}{msg, gormUtils.FileWithLineNum()}, data...)...)
}

// Error prints a error log.
func (l *logger) Error(ctx context.Context, msg string, data ...interface{}) {
	l.Zap.Errorf(messageFormat, append([]interface{}{msg, gormUtils.FileWithLineNum()}, data...)...)
}

// Trace prints a trace log such as sql, source file and error.
func (l *logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)

	switch {
	case err != nil:
		sql, _ := fc()
		l.GetZapLogger().Errorf(errorFormat, gormUtils.FileWithLineNum(), err, sql)
	case elapsed > slowThreshold*time.Millisecond && slowThreshold*time.Millisecond != 0:
		sql, _ := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", slowThreshold)
		l.GetZapLogger().Warnf(errorFormat, gormUtils.FileWithLineNum(), slowLog, sql)
	default:
		sql, _ := fc()
		l.GetZapLogger().Debugf(sqlFormat, sql)
	}
}

package logger

import (
	"database/sql/driver"
	"fmt"
	"io"
	"os"
	"reflect"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ybkuroki/go-webapp-sample/config"
)

// InitLogger initialize logger.
func InitLogger(e *echo.Echo, config *config.Config) {
	// logging for each request.
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: strings.Replace(config.Log.Format, "${level}", "INFO", 1) + "\n",
	}))
	// logging for the start and end of controller processes.
	// ref: https://echo.labstack.com/guide/customization
	e.Use(MyLoggerMiddleware)

	// set logformat for echo logger.
	e.Logger.SetHeader(config.Log.Format)
	e.Logger.SetLevel(config.Log.Level)

	// if the log file exists, write both console and the log file.
	if config.Log.FilePath != "" {
		logfile, err := os.OpenFile(config.Log.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 066)
		if err != nil {
			panic("Cannot open the log file. Please check this file path. Path: " + config.Log.FilePath + ", Error: " + err.Error())
		}
		defer logfile.Close()

		e.Logger.SetOutput(io.MultiWriter(logfile, os.Stdout))
	}
}

// MyLoggerMiddleware is middleware for logging the start and end of controller processes.
// ref: https://echo.labstack.com/cookbook/middleware
func MyLoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Logger().Debug(c.Path() + " Action Start")
		if err := next(c); err != nil {
			c.Error(err)
		}
		c.Logger().Debug(c.Path() + " Action End")
		return nil
	}
}

// ==============================================================
// Customize SQL Logger for gorm library
// ref: https://github.com/wantedly/gorm-zap
// ref: https://github.com/jinzhu/gorm/blob/master/logger.go
// ===============================================================

// Logger is an alternative implementation of *gorm.Logger
type Logger struct {
	logger echo.Logger
}

// NewLogger create logger object for *gorm.DB from *echo.Logger
func NewLogger(elog echo.Logger) *Logger {
	return &Logger{logger: elog}
}

// Print passes arguments to Println
func (l *Logger) Print(values ...interface{}) {
	l.Println(values)
}

// Println format & print log
func (l *Logger) Println(values []interface{}) {
	sql := createLog(values)
	if sql != "" {
		l.logger.Debug(sql)
	}
}

// createLog returns log for output
func createLog(values []interface{}) string {
	ret := ""

	if len(values) > 1 {
		var level = values[0]

		if level == "sql" {
			ret = "[gorm] : " + createSQL(values[3].(string), getFormattedValues(values))
		}
	}

	return ret
}

func isPrintable(s string) bool {
	for _, r := range s {
		if !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}

// getFormattedValues returns values of a SQL statement.
func getFormattedValues(values []interface{}) []string {
	var formattedValues []string
	for _, value := range values[4].([]interface{}) {
		indirectValue := reflect.Indirect(reflect.ValueOf(value))
		if indirectValue.IsValid() {
			value = indirectValue.Interface()
			if t, ok := value.(time.Time); ok {
				if t.IsZero() {
					formattedValues = append(formattedValues, fmt.Sprintf("'%v'", "0000-00-00 00:00:00"))
				} else {
					formattedValues = append(formattedValues, fmt.Sprintf("'%v'", t.Format("2006-01-02 15:04:05")))
				}
			} else if b, ok := value.([]byte); ok {
				if str := string(b); isPrintable(str) {
					formattedValues = append(formattedValues, fmt.Sprintf("'%v'", str))
				} else {
					formattedValues = append(formattedValues, "'<binary>'")
				}
			} else if r, ok := value.(driver.Valuer); ok {
				if value, err := r.Value(); err == nil && value != nil {
					formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
				} else {
					formattedValues = append(formattedValues, "NULL")
				}
			} else {
				switch value.(type) {
				case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, bool:
					formattedValues = append(formattedValues, fmt.Sprintf("%v", value))
				default:
					formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
				}
			}
		} else {
			formattedValues = append(formattedValues, "NULL")
		}
	}
	return formattedValues
}

// createSQL returns complete SQL with values bound to a SQL statement.
func createSQL(sql string, values []string) string {
	var (
		sqlRegexp                = regexp.MustCompile(`\?`)
		numericPlaceHolderRegexp = regexp.MustCompile(`\$\d+`)
		result                   = ""
	)
	// differentiate between $n placeholders or else treat like ?
	if numericPlaceHolderRegexp.MatchString(sql) {
		for index, value := range values {
			placeholder := fmt.Sprintf(`\$%d([^\d]|$)`, index+1)
			result = regexp.MustCompile(placeholder).ReplaceAllString(sql, value+"$1")
		}
	} else {
		formattedValuesLength := len(values)
		for index, value := range sqlRegexp.Split(sql, -1) {
			result += value
			if index < formattedValuesLength {
				result += values[index]
			}
		}
	}
	return result
}

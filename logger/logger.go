package logger

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"regexp"
	"time"
	"unicode"

	"github.com/ybkuroki/go-webapp-sample/config"
	"go.uber.org/zap"
)

var logger *Logger

// Logger is an alternative implementation of *gorm.Logger
type Logger struct {
	zap *zap.SugaredLogger
}

// GetLogger is return Logger
func GetLogger() *Logger {
	return logger
}

// GetZapLogger returns zapSugaredLogger
func GetZapLogger() *zap.SugaredLogger {
	return logger.zap
}

// newLogger create logger object for *gorm.DB from *echo.Logger
func newLogger(zap *zap.SugaredLogger) *Logger {
	return &Logger{zap: zap}
}

// InitLogger initialize logger.
func InitLogger(config *config.Config) {
	zap, err := zap.NewDevelopment()
	if err != nil {
		fmt.Printf("Error")
	}
	defer zap.Sync()
	sugar := zap.Sugar()
	// set package varriable logger.
	logger = newLogger(sugar)
}

// ==============================================================
// Customize SQL Logger for gorm library
// ref: https://github.com/wantedly/gorm-zap
// ref: https://github.com/jinzhu/gorm/blob/master/logger.go
// ===============================================================

// Print passes arguments to Println
func (l *Logger) Print(values ...interface{}) {
	l.Println(values)
}

// Println format & print log
func (l *Logger) Println(values []interface{}) {
	sql := createLog(values)
	if sql != "" {
		l.zap.Debugf(sql)
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

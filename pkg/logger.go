package logger

import (
	"strings"

	"github.com/amoghe/distillog"
)

var appLogger distillog.Logger

func InitLogger(logType, name string) distillog.Logger {
	logType = strings.ToLower(logType)
	switch logType {
	case "stderr":
		appLogger = distillog.NewStderrLogger(name)
	case "syslog":
		appLogger = distillog.NewSyslogLogger(name)
	default:
	case "stdout":
		appLogger = distillog.NewStdoutLogger(name)
	}
	return appLogger
}

func GetLogger() distillog.Logger {
	if appLogger == nil {
		appLogger = distillog.NewStdoutLogger("")
	}
	return appLogger
}

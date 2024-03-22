package log

import (
	"github.com/rs/zerolog"
)

var globalLogger Logging

func LoadLogger() {
	globalLogger = NewLogger()
}

func Info(msg string) {
	globalLogger.Info(msg)
}

func Infof(msg string, args ...interface{}) {
	globalLogger.Infof(msg, args...)
}

func Debug(msg string) {
	globalLogger.Debug(msg)
}

func Debugf(msg string, args ...interface{}) {
	globalLogger.Debugf(msg, args...)
}

func Warn(msg string) {
	globalLogger.Warn(msg)
}

func Warnf(msg string, args ...interface{}) {
	globalLogger.Warnf(msg, args...)
}

func Error(err error, msg string) {
	globalLogger.Error(err, msg)
}

func Errorf(err error, msg string, args ...interface{}) {
	globalLogger.Errorf(err, msg, args...)
}

func Fatal(err error, msg string) {
	globalLogger.Fatal(err, msg)
}

func Fatalf(err error, msg string, args ...interface{}) {
	globalLogger.Fatalf(err, msg, args...)
}

func GetZeroLog() *zerolog.Logger {
	return globalLogger.GetZeroLog()
}

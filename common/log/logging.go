package log

import "github.com/rs/zerolog"

type Logging interface {
	Info(msg string)

	Infof(msgFormat string, args ...interface{})

	Debug(msg string)

	Debugf(msgFormat string, args ...interface{})

	Warn(msg string)

	Warnf(msgFormat string, args ...interface{})

	Error(err error, msg string)

	Errorf(err error, msgFormat string, args ...interface{})

	Fatal(err error, msg string)

	Fatalf(err error, msgFormat string, args ...interface{})

	GetZeroLog() *zerolog.Logger
}

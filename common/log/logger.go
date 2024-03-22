package log

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
)

type logger struct {
	zeroLog *zerolog.Logger
}

func NewLogger() Logging {
	zeroLog := zerolog.New(os.Stderr).With().Timestamp().Logger()
	return &logger{
		zeroLog: &zeroLog,
	}
}

func (l *logger) Info(msg string) {
	l.zeroLog.Info().Msg(msg)
}

func (l *logger) Infof(msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	l.zeroLog.Info().Msg(msg)
}

func (l *logger) Debug(msg string) {
	l.zeroLog.Debug().Msg(msg)
}

func (l *logger) Debugf(msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	l.zeroLog.Debug().Msg(msg)
}

func (l *logger) Warn(msg string) {
	l.zeroLog.Warn().Msg(msg)
}

func (l *logger) Warnf(msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	l.zeroLog.Warn().Msg(msg)
}

func (l *logger) Error(err error, msg string) {
	l.zeroLog.Error().Err(err).Msg(msg)
}

func (l *logger) Errorf(err error, msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	l.zeroLog.Error().Err(err).Msg(msg)
}

func (l *logger) Fatal(err error, msg string) {
	l.zeroLog.Fatal().Err(err).Msg(msg)
}

func (l *logger) Fatalf(err error, msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	l.zeroLog.Fatal().Err(err).Msg(msg)
}

func (l *logger) GetZeroLog() *zerolog.Logger {
	return l.zeroLog
}

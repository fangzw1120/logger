// Package logger provides basic functions.
// @Description:
package logger

import (
	"github.com/rs/zerolog"
	"strings"
)

// Logger ...
type Logger struct {
	Logger       *zerolog.Logger
	removePrefix string
}

// Debug ...
func (l *Logger) Debug(msg string) {
	msg = strings.TrimPrefix(msg, l.removePrefix)
	l.Logger.Debug().Msg(msg)
}

// Debugf ...
func (l *Logger) Debugf(format string, v ...interface{}) {
	format = strings.TrimPrefix(format, l.removePrefix)
	l.Logger.Debug().Msgf(format, v...)
}

// Info ...
func (l *Logger) Info(msg string) {
	msg = strings.TrimPrefix(msg, l.removePrefix)
	l.Logger.Info().Msg(msg)
}

// Infof ...
func (l *Logger) Infof(format string, v ...interface{}) {
	format = strings.TrimPrefix(format, l.removePrefix)
	l.Logger.Info().Msgf(format, v...)
}

// Warn ...
func (l *Logger) Warn(msg string) {
	msg = strings.TrimPrefix(msg, l.removePrefix)
	l.Logger.Warn().Msgf(msg)
}

// Warnf ...
func (l *Logger) Warnf(format string, v ...interface{}) {
	format = strings.TrimPrefix(format, l.removePrefix)
	l.Logger.Warn().Msgf(format, v...)
}

// Error ...
func (l *Logger) Error(msg string) {
	msg = strings.TrimPrefix(msg, l.removePrefix)
	l.Logger.Error().Msg(msg)
}

// Errorf ...
func (l *Logger) Errorf(format string, v ...interface{}) {
	format = strings.TrimPrefix(format, l.removePrefix)
	l.Logger.Error().Msgf(format, v...)
}

// Fatal ...
func (l *Logger) Fatal(msg string) {
	msg = strings.TrimPrefix(msg, l.removePrefix)
	l.Logger.Fatal().Msg(msg)
}

// Fatalf ...
func (l *Logger) Fatalf(format string, v ...interface{}) {
	format = strings.TrimPrefix(format, l.removePrefix)
	l.Logger.Fatal().Msgf(format, v...)
}

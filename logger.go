//go:build linux || darwin || windows
// +build linux darwin windows

// Package logger provides basic functions.
// @Description:
package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path"
	"syscall"
	"time"
)

var (
	logger      *Logger
	stateLogger *Logger
	params      *Params
	stateParams *Params
)

const (
	// SigDebug ...
	SigDebug = syscall.Signal(0x35)
	// SigInfo ...
	SigInfo = syscall.Signal(0x36)
	// SigError ...
	SigError = syscall.Signal(0x37)
)

// Config ...
type Config struct {
	ConsoleLoggingEnabled bool
	EncodeLogsAsJSON      bool
	FileLoggingEnabled    bool
	Directory             string
	Filename              string
	MaxSize               int
	MaxBackups            int
	MaxAge                int
	IfDebug               bool
	IfSimpleLog           bool
}

// Init default 30day, 50MB/file, params to set removePathPrefix
func Init(directory, fileName string, ifDebug bool, ifSimpleLog bool, ifVerbose bool, args ...Params) {
	params = &Params{}
	if len(args) > 0 {
		params = &args[0]
	}
	logger = Configure(Config{
		ConsoleLoggingEnabled: ifVerbose,
		EncodeLogsAsJSON:      true,
		// supervisor会重定向标准输出到文件，可以不需要再次输出文件日志
		FileLoggingEnabled: true,
		Directory:          directory,
		Filename:           fileName,
		IfDebug:            ifDebug,
		IfSimpleLog:        ifSimpleLog,
		MaxSize:            MaxSizeMB,
		MaxBackups:         MaxBackupCnt,
		MaxAge:             MaxAgeDay,
	})
	logger.removePrefix = params.RemovePathPrefix
}

// InitStateLogger ...
func InitStateLogger(directory, fileName string, ifDebug bool, ifSimpleLog bool, ifVerbose bool, args ...Params) {
	stateParams = &Params{}
	if len(args) > 0 {
		stateParams = &args[0]
	}
	stateLogger = Configure(Config{
		ConsoleLoggingEnabled: false,
		EncodeLogsAsJSON:      true,
		// supervisor会重定向标准输出到文件，可以不需要再次输出文件日志
		FileLoggingEnabled: true,
		Directory:          directory,
		Filename:           fileName,
		IfDebug:            ifDebug,
		IfSimpleLog:        ifSimpleLog,
		MaxSize:            MaxSizeMB,
		MaxBackups:         MaxBackupCnt,
		MaxAge:             MaxAgeDay,
	})
	stateLogger.removePrefix = stateParams.RemovePathPrefix
}

// Configure ...
func Configure(config Config) *Logger {
	var writer []io.Writer

	if config.ConsoleLoggingEnabled {
		writer = append(writer, zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	}
	if config.FileLoggingEnabled {
		writer = append(writer, newRollingFile(config))
	}
	if config.IfDebug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		if config.IfSimpleLog {
			zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		} else {
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		}
	}

	mw := io.MultiWriter(writer...)
	l := zerolog.New(mw).With().Timestamp().Logger()

	l.Info().Bool("fileLogging", config.FileLoggingEnabled).
		Bool("jsonLogOutput", config.EncodeLogsAsJSON).
		Str("logDirectory", config.Directory).
		Str("fileName", config.Filename).
		Int("maxSizeMB", config.MaxSize).
		Int("maxBackups", config.MaxBackups).
		Int("maxAgeInDays", config.MaxAge).
		Msg("logging configured")
	return &Logger{Logger: &l}
}

// ChangeLogLevel waiting for signal to change log level
func ChangeLogLevel(sig os.Signal) {
	switch sig {
	case SigDebug:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case SigInfo:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case SigError:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	}
}

func newRollingFile(config Config) io.Writer {
	if err := os.MkdirAll(config.Directory, 0744); err != nil {
		log.Error().Err(err).Str("path", config.Directory).Msg("can't create log directory")
		return nil
	}
	return &lumberjack.Logger{
		Filename:   path.Join(config.Directory, config.Filename),
		MaxSize:    MaxSizeMB,
		MaxAge:     MaxAgeDay,
		MaxBackups: MaxBackupCnt,
		Compress:   true,
	}
}

// Debug ...
func Debug(msg string) {
	msg = FILE() + ":" + LINE() + " " + msg
	logger.Debug(msg)
}

// Debugf ...
func Debugf(format string, v ...interface{}) {
	format = FILE() + ":" + LINE() + " " + format
	logger.Debugf(format, v...)
}

// Info ...
func Info(msg string) {
	msg = FILE() + ":" + LINE() + " " + msg
	logger.Info(msg)
}

// Infof ...
func Infof(format string, v ...interface{}) {
	format = FILE() + ":" + LINE() + " " + format
	logger.Infof(format, v...)
}

// Warn ...
func Warn(msg string) {
	msg = FILE() + ":" + LINE() + " " + msg
	logger.Warn(msg)
}

// Warnf ...
func Warnf(format string, v ...interface{}) {
	format = FILE() + ":" + LINE() + " " + format
	logger.Warnf(format, v...)
}

// Error ...
func Error(msg string) {
	msg = FILE() + ":" + LINE() + " " + msg
	logger.Errorf(msg)
}

// Errorf ...
func Errorf(format string, v ...interface{}) {
	format = FILE() + ":" + LINE() + " " + format
	logger.Errorf(format, v...)
}

// Fatal ...
func Fatal(msg string) {
	msg = FILE() + ":" + LINE() + " " + msg
	logger.Fatal(msg)
}

// Fatalf ...
func Fatalf(format string, v ...interface{}) {
	format = FILE() + ":" + LINE() + " " + format
	logger.Fatalf(format, v...)
}

// CheckErr ...
func CheckErr(err error, msg string) {
	if err != nil {
		Errorf(msg, err)
	}
}

// StatePrint ...
func StatePrint(msg string) {
	msg = FILE() + ":" + LINE() + " " + msg
	stateLogger.Info(msg)
}

// StatePrintf ...
func StatePrintf(format string, v ...interface{}) {
	format = FILE() + ":" + LINE() + " " + format
	stateLogger.Infof(format, v...)
}

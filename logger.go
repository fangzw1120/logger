//go:build linux || darwin || windows
// +build linux darwin windows

package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path"
	"runtime"
	"strconv"
	"syscall"
	"time"
)

var (
	logger      *Logger
	stateLogger *Logger
)

const (
	SigDebug = syscall.Signal(0x35)
	SigInfo  = syscall.Signal(0x36)
	SigError = syscall.Signal(0x37)
)

var (
	MaxAgeDay    = 30
	MaxSizeMB    = 50
	MaxBackupCnt = 0
)

func GetMaxAgeDay() int {
	return MaxAgeDay
}
func SetMaxAgeDay(v int) {
	MaxAgeDay = v
}
func GetMaxSizeMB() int {
	return MaxSizeMB
}
func SetMaxSizeMB(v int) {
	MaxSizeMB = v
}
func GetMaxBackupCnt() int {
	return MaxBackupCnt
}
func SetMaxBackupCnt(v int) {
	MaxBackupCnt = v
}

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

// Init ...
func Init(directory, fileName string, ifDebug bool, ifSimpleLog bool, ifVerbose bool) {
	logger = Configure(Config{
		ConsoleLoggingEnabled: ifVerbose,
		EncodeLogsAsJSON:      false,
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

	logFile := &lumberjack.Logger{
		Filename:   fileName, // 日志文件名
		MaxSize:    50,       // 每个日志文件的最大尺寸，单位 MB
		MaxAge:     7,        // 保留日志文件的最大天数
		MaxBackups: 10,       // 保留日志文件的最大个数
		LocalTime:  true,     // 使用本地时间
		Compress:   true,     // 是否压缩旧日志文件
	}
	// 设置 zerolog 的全局日志级别为 debug
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	// 创建一个 io.Writer 对象用于同时输出到控制台和日志文件
	multiWriter := io.MultiWriter(os.Stdout, logFile)

	// 创建一个自定义的 zerolog.Formatter 对象
	formatter := func() *zerolog.ConsoleWriter {
		return &zerolog.ConsoleWriter{
			TimeFormat: time.RFC3339, // 时间格式
			NoColor:    true,         // 禁用控制台颜色
			Out:        multiWriter,  // 输出目标为 multiWriter
		}
	}

	// 将日志输出格式设置为自定义的格式
	c := log.Output(formatter())
	logger.Logger = &c
	// 输出一条测试日志
	logger.Debug("this is a test log message")
	// 关闭日志文件
	//logFile.Close()
}

// InitStateLogger ...
func InitStateLogger(directory, fileName string, ifDebug bool, ifSimpleLog bool, ifVerbose bool) {
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
	}
}

// Logger ...
type Logger struct {
	Logger *zerolog.Logger
}

// Debug ...
func (l *Logger) Debug(msg string) {
	l.Logger.Debug().Msg(msg)
}

// Debugf ...
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.Logger.Debug().Msgf(format, v...)
}

// Info ...
func (l *Logger) Info(msg string) {
	l.Logger.Info().Msg(msg)
}

// Infof ...
func (l *Logger) Infof(format string, v ...interface{}) {
	l.Logger.Info().Msgf(format, v...)
}

// Warn ...
func (l *Logger) Warn(msg string) {
	l.Logger.Warn().Msgf(msg)
}

// Warnf ...
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.Logger.Warn().Msgf(format, v...)
}

// Error ...
func (l *Logger) Error(msg string) {
	l.Logger.Error().Msg(msg)
}

// Errorf ...
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Logger.Error().Msgf(format, v...)
}

// Fatal ...
func (l *Logger) Fatal(msg string) {
	l.Logger.Fatal().Msg(msg)
}

// Fatalf ...
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Logger.Fatal().Msgf(format, v...)
}

// StringerFunc ...
type StringerFunc func() string

// String ...
func (f StringerFunc) String() string {
	return f()
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

// FILE ...
func FILE() string {
	_, file, _, _ := runtime.Caller(2)
	return file
}

// LINE ...
func LINE() string {
	_, _, line, _ := runtime.Caller(2)
	return strconv.Itoa(line)
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

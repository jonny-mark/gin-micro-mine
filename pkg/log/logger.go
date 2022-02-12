/**
 * @author jiangshangfang
 * @date 2021/7/30 11:53 AM
 **/
package log

import (
	"gin/pkg/config"
	"log"
)

//全局log
var logger Logger

// Fields Type to pass when we want to call WithFields for structured logging
type Fields map[string]interface{}

type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})

	Warn(args ...interface{})
	Warnf(format string, args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})

	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})

	Panic(args ...interface{})
	Panicf(format string, args ...interface{})

	WithFields(keyValues Fields) Logger
}

//初始化
func Init() Logger {
	var cfg Config

	if err := config.Load("logger", &cfg); err != nil {
		log.Panicf("load logger conf err: %v", err)
	}
	logger = newZapLogger(&cfg)
	return logger
}

func GetLogger() Logger {
	return logger
}

// Info log
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Warn log
func Warn(args ...interface{}) {
	logger.Warn(args...)
}

// Error log
func Error(args ...interface{}) {
	logger.Error(args...)
}

// Fatal log
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

// Panic log
func Panic(args ...interface{}) {
	logger.Panic(args...)
}

// Debugf logger
func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

// Infof logger
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

// Warnf logger
func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

// Errorf logger
func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

// Fatalf logger
func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

// Panicf logger
func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}

func WithFields(keyValues Fields) Logger {
	return logger.WithFields(keyValues)
}

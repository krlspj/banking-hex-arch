package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger
var Logg ILogger

func init() {
	var err error
	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//encoderConfig.StacktraceKey = ""
	config.EncoderConfig = encoderConfig

	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}
func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}
func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}

type ILogger interface {
	Info(message string)
	Infof(string, ...interface{})
	Debug(message string)
	Error(message string)
}

type loggerWrapper struct {
	lw *zap.Logger
}

func NewLogger() ILogger {
	var err error
	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//encoderConfig.StacktraceKey = ""
	config.EncoderConfig = encoderConfig

	l, err := config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	Logg = &loggerWrapper{
		lw: l,
	}
	return &loggerWrapper{
		lw: l,
	}
}

func (l *loggerWrapper) Info(message string) {
	l.lw.Info(message)
}
func (l *loggerWrapper) Infof(format string, args ...any) {
	l.lw.Sugar().Infof(format, args)
}
func (l *loggerWrapper) Debug(message string) {
	l.lw.Debug(message)
}
func (l *loggerWrapper) Error(message string) {
	l.lw.Error(message)
}

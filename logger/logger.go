package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	var err error

	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig = encoderConfig
	config.EncoderConfig.StacktraceKey = ""
	log, err = config.Build(zap.AddCallerSkip(1))
	// log, err = zap.NewProduction(zap.AddCallerSkip(1))
	if err != nil {
		panic(err.Error())
	}

}

func Info(msg string, fields ...zap.Field) {
	log.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	log.Error(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	log.Debug(msg, fields...)
}

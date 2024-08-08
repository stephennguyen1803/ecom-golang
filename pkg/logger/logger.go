package logger

import (
	"ecom-project/pkg/settings"
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(loggerConfig settings.LoggerSetting) *LoggerZap {
	// debug -> info-> warn -> error -> dpanic -> panic -> fatal
	// Create a new logger
	var level zapcore.Level
	switch loggerConfig.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	fmt.Println(loggerConfig.FileLogName)
	encoder := getEncodeLog()
	hook := lumberjack.Logger{
		Filename:   loggerConfig.FileLogName,
		MaxSize:    loggerConfig.MaxSize, // megabytes
		MaxAge:     loggerConfig.MaxAge,  // days
		MaxBackups: loggerConfig.MaxBackups,
		Compress:   loggerConfig.Compress,
	}
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		level)
	// logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	return &LoggerZap{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
}

func getEncodeLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()

	// Customize the time format 17723424 -> 2021-09-01 12:00:00
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// change key ts to timestamp
	encodeConfig.TimeKey = "timestamp"

	// change key msg to message
	encodeConfig.MessageKey = "message"

	//change type to UpperCase
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// set shortcaller endcoder
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encodeConfig)
}

package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	encoder := getEncodeLog()
	sync := getWriteSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core)

	//demo using zap logger
	logger.Info("Info Log", zap.Int("line", 1))
	logger.Warn("Warn Log", zap.Int("line", 2))
	logger.Error("Error Log", zap.Int("line", 3))
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

func getWriteSync() zapcore.WriteSyncer {
	// file, _ := os.OpenFile("./log/loginfo.log", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	file, err := os.OpenFile("./log/loginfo.log", os.O_RDWR, os.ModePerm)
	if err != nil {
		file, err = os.Create("./log/loginfo.log")
		if err != nil {
			panic(err)
		}
	}
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncFile, syncConsole)
}

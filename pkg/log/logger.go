package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

var ZLogger *zap.Logger
var Logger *zap.SugaredLogger

func InitializeLogger() {
	config := zap.NewProductionEncoderConfig()
	// Setting time encoder
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	// Setting Log level should be printed in capital letters with level colors
	config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(config)

	var cores []zapcore.Core

	cores = append(cores, zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.InfoLevel))

	core := zapcore.NewTee(cores...)

	if len(cores) == 0 {
		log.Println("[WARN] No logger configured!")
	}

	ZLogger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	defer ZLogger.Sync()
	Logger = ZLogger.Sugar()
	defer Logger.Sync()
}

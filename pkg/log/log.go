package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"server/pkg/config"
)

func GetLogger(cfg *config.Config)*zap.Logger {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	switch cfg.LogConfig.Output {
	case "all":
		writeSyncer=zapcore.NewMultiWriteSyncer(os.Stdout,writeSyncer) //暂时不启用文件
	case "file":
	default:
		writeSyncer=zapcore.NewMultiWriteSyncer(os.Stdout) //暂时不启用文件
	}

	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	return logger
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./data/server.log",
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}


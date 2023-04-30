package log

import (
	"fmt"
	"path"
	"time"

	"github.com/KarlvenK/tgBot/config"
	rotateLog "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.SugaredLogger
)

func init() {
	core := zapcore.NewCore(getEncoder(), getLogWriter(), zapcore.InfoLevel)
	tmpLogger := zap.New(core, zap.AddCaller())
	logger = tmpLogger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoder)
}

func getLogWriter() zapcore.WriteSyncer {
	l, err := rotateLog.New(
		path.Join(config.GetConfig().LogPath, "access_log.%Y%m%d%H%M"),
		rotateLog.WithLinkName(path.Join(config.GetConfig().LogPath, "access_log")),
		rotateLog.WithMaxAge(24*time.Hour),
		rotateLog.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		panic(fmt.Errorf("fail to create rotatelogs: %v", err))
	}
	writeSyncer := zapcore.AddSync(l)
	return writeSyncer
}

// Flush
//
//	@Description: Called when program exits
func Flush() {
	logger.Sync()
}

// Info
//
//	@Description: Wraps zap.Info
//	@param args
func Info(args ...any) {
	logger.Info(args...)
}

// Warn
//
//	@Description: Wraps of zap.Warn
//	@param args
func Warn(args ...any) {
	logger.Warn(args...)
}

// Error
//
//	@Description: Wraps zap.Error
//	@param args
func Error(args ...any) {
	logger.Error(args...)
}

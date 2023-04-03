package utils

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/config"
)

// NewLogger creates a new instance of logger
func NewLogger(logger config.LoggerConfig) (*zap.Logger, error) {

	// Check log level
	level := zap.NewAtomicLevel()
	switch logger.LogLevel {
	case "debug":
		level.SetLevel(zap.DebugLevel)
	case "info":
		level.SetLevel(zap.InfoLevel)
	case "warn":
		level.SetLevel(zap.WarnLevel)
	case "error":
		level.SetLevel(zap.ErrorLevel)
	case "fatal":
		level.SetLevel(zap.FatalLevel)
	case "panic":
		level.SetLevel(zap.PanicLevel)
	default:
		return nil, fmt.Errorf("unknown log level: %s", logger.LogLevel)
	}

	// Initialize instance
	cfg := zap.Config{
		Encoding:         "json",
		Level:            level,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			LevelKey:    "level",
			EncodeLevel: zapcore.LowercaseLevelEncoder,
		},
	}

	return cfg.Build()
}

package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger creates a new instance of logger
func NewLogger() (*zap.Logger, error) {
	cfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
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

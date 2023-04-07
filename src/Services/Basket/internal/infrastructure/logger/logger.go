package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

// InitLogger initializes a new logrus logger instance with specified configuration
func InitLogger() *logrus.Entry {
	// Create a new logger instance
	logger := logrus.New()

	// Set log level to Info
	logger.SetLevel(logrus.InfoLevel)

	// Set formatter to json
	logger.SetFormatter(&logrus.JSONFormatter{})

	// Create a new log file
	logFile, err := os.OpenFile("log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logger.Fatal("Failed to open log file: ", err)
	}

	// Set output to both stdout and log file
	logger.SetOutput(io.MultiWriter(os.Stdout, logFile))

	return logger.WithFields(logrus.Fields{})
}

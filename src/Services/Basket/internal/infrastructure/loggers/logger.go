package loggers

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

// InitLogger initializes a new logrus logger instance with specified configuration
func InitLogger() *logrus.Entry {
	// Create a new logger instance
	logger := logrus.New()

	// Set formatter to json
	logger.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logger.SetOutput(os.Stdout)

	// Set log level to Info
	logger.SetLevel(logrus.TraceLevel)

	// Create a new log file
	logFile, err := os.OpenFile("log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logger.Fatal("Failed to open log file: ", err)
	}

	// Set output to both stdout and log file
	logger.SetOutput(io.MultiWriter(os.Stdout, logFile))

	return logger.WithFields(logrus.Fields{})
}

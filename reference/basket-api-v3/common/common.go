package common

import (
	"encoding/json"
	"os"

	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

// Configuration stores setting values
type Configuration struct {
	EnableZapLog        bool `json:"enableZapLog"`
	EnableGinConsoleLog bool `json:"enableGinConsoleLog"`
	EnableGinFileLog    bool `json:"enableGinFileLog"`

	LogFilename   string `json:"logFilename"`
	LogMaxSize    int    `json:"logMaxSize"`
	LogMaxBackups int    `json:"logMaxBackups"`
	LogMaxAge     int    `json:"logMaxAge"`
}

// Config shares the global configuration
var (
	Config *Configuration
)

var Logger *zap.Logger

// InitializeLogger loads configuration to enable logger the config file
func InitializeLogger() error {
	//Filename is the path to the json config file
	file, err := os.Open("common/config.json")
	if err != nil {
		return err
	}

	// Allocate memory for global configuration
	Config = new(Configuration)
	// decode json info to get configuration
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		return err
	}

	// Setting Service Logger
	log.SetOutput(&lumberjack.Logger{
		Filename:   Config.LogFilename,
		MaxSize:    Config.LogMaxSize,    // megabytes after which new file is created
		MaxBackups: Config.LogMaxBackups, // number of backups
		MaxAge:     Config.LogMaxAge,     // days
	})

	log.SetLevel(log.DebugLevel)

	// log.SetFormatter(&log.TextFormatter{})
	log.SetFormatter(&log.JSONFormatter{})

	return nil
}

package config

import (
	"github.com/spf13/viper"
)

// Configuration stores setting environment values
type Config struct {
	Port         string `mapstructure:"PORT"`
	EnableLogger bool   `mapstructure:"ENABLE_LOGGER"`

	GrpcServerAddress string `mapstructure:"GRPC_SERVER_ADDRESS"`

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
}

// LoadConfig loads configuration from the config file
func LoadConfig(path string) (config Config, err error) {

	// Path to file environment config
	viper.AddConfigPath(path)

	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	// Reading configure from app.env
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	// Return configuration to main
	err = viper.Unmarshal(&config)
	return
}

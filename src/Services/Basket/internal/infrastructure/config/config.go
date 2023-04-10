package config

import (
	"github.com/spf13/viper"
)

type RedisConfig struct {
}

// Configuration stores setting environment values
type Config struct {
	DBUri         string `mapstructure:"MONGODB_LOCAL_URI"`
	RedisUri      string `mapstructure:"REDIS_URL"`
	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisPort     string `mapstructure:"REDIS_PORT"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
	Port          string `mapstructure:"PORT"`
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

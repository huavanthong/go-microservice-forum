package configs

import (
	"github.com/spf13/viper"
)

// Configuration stores setting environment values
type Config struct {
	AppPort string `mapstructure:"APP_PORT"`

	DBLocalUri     string `mapstructure:"MONGODB_LOCAL_URI"`
	DBContainerUri string `mapstructure:"MONGODB_CONTAINER_URI"`
	DBName         string `mapstructure:"MONGODB_NAME"`
	DBCollProduct  string `mapstructure:"MONGODB_COLLECTION_PRODUCT"`
	DBCollCategory string `mapstructure:"MONGODB_COLLECTION_CATEGORY"`

	EnableLogger bool `mapstructure:"ENABLE_LOGGER"`
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

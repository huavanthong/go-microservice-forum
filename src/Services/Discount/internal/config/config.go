package config

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
	} `yaml:"database"`
}

func LoadConfig(configPath string) (*Config, error) {
	// Read config file
	configFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
		return nil, err
	}

	// Parse config file
	config := Config{}
	if err := yaml.Unmarshal(configFile, &config); err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
		return nil, err
	}

	return &config, nil
}

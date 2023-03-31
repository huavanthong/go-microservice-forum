package main

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/models"
)

func Migrate() error {
	// Load configuration
	viper.SetConfigFile("config.yml")
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	// Connect to database
	dsn := viper.GetString("database.dsn")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// Migrate schema
	if err := db.AutoMigrate(&models.Discount{}, &models.Coupon{}); err != nil {
		return fmt.Errorf("failed to migrate schema: %w", err)
	}

	return nil
}

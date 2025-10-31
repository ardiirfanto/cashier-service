package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// Config holds all configuration values for the application
type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	JWT      JWTConfig
}

// DatabaseConfig holds database connection parameters
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port string
}

// JWTConfig holds JWT configuration
type JWTConfig struct {
	Secret string
}

// LoadConfig loads configuration from environment variables using Viper
func LoadConfig() (*Config, error) {
	// Set default configuration file name and type
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./")

	// Enable reading from environment variables
	viper.AutomaticEnv()

	// Set default values
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "3306")
	viper.SetDefault("DB_USER", "root")
	viper.SetDefault("DB_PASS", "")
	viper.SetDefault("DB_NAME", "cashier_db")
	viper.SetDefault("JWT_SECRET", "supersecretkey")
	viper.SetDefault("SERVER_PORT", "8080")

	// Read configuration file (optional, will use env vars if not found)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("No .env file found, using environment variables or defaults")
		} else {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	}

	// Map values to Config struct
	config := &Config{
		Database: DatabaseConfig{
			Host:     viper.GetString("DB_HOST"),
			Port:     viper.GetString("DB_PORT"),
			User:     viper.GetString("DB_USER"),
			Password: viper.GetString("DB_PASS"),
			Name:     viper.GetString("DB_NAME"),
		},
		Server: ServerConfig{
			Port: viper.GetString("SERVER_PORT"),
		},
		JWT: JWTConfig{
			Secret: viper.GetString("JWT_SECRET"),
		},
	}

	return config, nil
}

// GetDSN returns the MySQL Data Source Name for database connection
func (c *Config) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Database.User,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.Name,
	)
}

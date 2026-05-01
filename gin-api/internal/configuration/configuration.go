package configuration

import (
	"fmt"

	"github.com/spf13/viper"
)

// Configuration struct holds the application configuration.
type Configuration struct {
	Server struct {
		Port int
	}

	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		Name     string
		SSLMode  string `mapstructure:"sslmode"`
	}

	Environment string
}

// Load is a function to load the configuration from a file.
func Load(configsPath, environment string) (*Configuration, error) {
	configPath := fmt.Sprintf("%s/%s.yaml", configsPath, environment)

	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	var config Configuration
	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}

	config.Environment = environment
	return &config, nil
}

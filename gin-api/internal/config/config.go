package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var AppConfiguration Configuration

// Configuration is a struct to bundle configurations for the application.
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
	}

	Environment string
}

// Load is a function to load the configuration from a file.
// It uses viper to read the configuration file and unmarshal it into the Configuration struct.
func Load(environment string) error {
	path := fmt.Sprintf("%s/%s.yaml", GetConfigsPath(), environment)
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("read config: %w", err)
	}
	if err := viper.Unmarshal(&AppConfiguration); err != nil {
		return fmt.Errorf("unmarshal config: %w", err)
	}
	return nil
}

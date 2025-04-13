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
	}

	Environment string
}

// Load is a function to load the configuration from a file.
// It uses viper to read the configuration file and unmarshal it into the Configuration struct.
func Load(environment string) (*Configuration, error) {
	path := fmt.Sprintf("%s/%s.yaml", GetConfigsPath(), environment)
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}
	var config Configuration
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}
	return &config, nil
}

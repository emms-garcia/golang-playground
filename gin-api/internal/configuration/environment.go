package configuration

import (
	"fmt"
	"os"
)

const (
	Development string = "development"
	Production  string = "production"
	Test        string = "test"
)

// GetConfigsPath returns the path to the configuration files.
func GetConfigsPath() (string, error) {
	path := os.Getenv("CONFIGS_PATH")
	if path == "" {
		return "", fmt.Errorf("CONFIGS_PATH environment variable is not set")
	}
	return path, nil
}

// GetEnvironment returns the current environment.
func GetEnvironment() (string, error) {
	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		return "", fmt.Errorf("ENVIRONMENT environment variable is not set")
	}
	if environment != string(Development) && environment != string(Production) && environment != string(Test) {
		return "", fmt.Errorf("ENVIRONMENT variable must be one of %s, %s, or %s", Development, Production, Test)
	}
	return environment, nil
}

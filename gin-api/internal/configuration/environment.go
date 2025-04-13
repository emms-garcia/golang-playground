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
// It checks the CONFIGS_PATH environment variable and panics if it's not set.
func GetConfigsPath() string {
	path := os.Getenv("CONFIGS_PATH")
	if path == "" {
		panic("CONFIGS_PATH environment variable is not set")
	}
	return path
}

// GetEnvironment returns the current environment.
// It checks the ENVIRONMENT environment variable and panics if it's not set or not one of the expected values.
func GetEnvironment() string {
	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		panic("ENVIRONMENT environment variable is not set")
	}
	if environment != string(Development) && environment != string(Production) && environment != string(Test) {
		panic(fmt.Sprintf("ENVIRONMENT variable must be one of %s, %s, or %s", Development, Production, Test))
	}
	return environment
}

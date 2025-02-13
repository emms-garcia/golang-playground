package main

import "fmt"

// Configuration is a struct to hold the configuration of the application
type Configuration struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
}

// BuildDBDSN is a method to build the DringSN st for the database connection
func (c *Configuration) BuildDBDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", c.DBHost, c.DBUser, c.DBPassword, c.DBName)
}

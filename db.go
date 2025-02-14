package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConfigureDB is a function to configure the database connection
func ConfigureDB(configuration *Configuration) *gorm.DB {
	db, err := gorm.Open(postgres.Open(configuration.BuildDBDSN()))
	if err != nil {
		panic("failed to connect with db")
	}
	// TODO: this shoudl be in a migration file with something like golang-migrate
	db.AutoMigrate(&Todo{})
	return db
}

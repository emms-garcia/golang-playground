package config

import (
	"fmt"

	"github.com/emms-garcia/golang-playground/gin-api/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConfigureDB is a function to configure the database connection
func ConfigureDB(config *Configuration) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", config.Database.Host, config.Database.User, config.Database.Password, config.Database.Name)
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic("failed to connect with db")
	}
	// TODO: this should be in a migration file with something like golang-migrate
	db.AutoMigrate(&model.Todo{})
	return db
}

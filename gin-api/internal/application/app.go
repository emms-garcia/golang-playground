package application

import (
	"fmt"
	"os"

	"github.com/emms-garcia/golang-playground/gin-api/internal/configuration"
	"github.com/emms-garcia/golang-playground/gin-api/internal/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Application struct holds the application configuration, database connection, and logger instance.
type Application struct {
	Config *configuration.Configuration
	DB     *gorm.DB
	Logger *zap.Logger
}

func Load() *Application {
	environment := os.Getenv("ENVIRONMENT")
	// Load configuration
	config, err := configuration.Load(environment)
	if err != nil {
		panic(fmt.Errorf("error loading configuration: %w", err))
	}
	// Initialize logger
	logger, err := logger.Load(environment)
	if err != nil {
		panic(fmt.Errorf("error loading logger: %w", err))
	}
	// Initialize database connection
	db, err := configuration.ConfigureDB(config)
	if err != nil {
		panic(fmt.Errorf("error connecting to database: %W", err))
	}
	return &Application{
		Config: config,
		DB:     db,
		Logger: logger,
	}
}

package application

import (
	"context"
	"errors"
	"fmt"

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

// Close gracefully closes the database connection if it exists.
func (a *Application) Close() error {
	if a == nil || a.DB == nil {
		return nil
	}

	sqlDB, err := a.DB.DB()
	if err != nil {
		return fmt.Errorf("get sql db: %w", err)
	}

	if err := sqlDB.Close(); err != nil && !errors.Is(err, gorm.ErrInvalidDB) {
		return fmt.Errorf("close database: %w", err)
	}

	return nil
}

func Load(ctx context.Context, configsPath, environment string) (*Application, error) {
	// Load configuration
	config, err := configuration.Load(configsPath, environment)
	if err != nil {
		return nil, fmt.Errorf("load configuration: %w", err)
	}

	// Initialize logger
	logger, err := logger.Load(environment)
	if err != nil {
		return nil, fmt.Errorf("load logger: %w", err)
	}

	// Initialize database connection
	db, err := configuration.ConfigureDB(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("connect database: %w", err)
	}

	return &Application{
		Config: config,
		DB:     db,
		Logger: logger,
	}, nil
}

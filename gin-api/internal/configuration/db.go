package configuration

import (
	"context"
	"fmt"

	"github.com/emms-garcia/golang-playground/gin-api/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConfigureDB is a function to configure the database connection
func ConfigureDB(ctx context.Context, config *Configuration) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		config.Database.Host,
		config.Database.User,
		config.Database.Password,
		config.Database.Name,
		config.Database.Port,
		config.Database.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("get sql db: %w", err)
	}

	if err := sqlDB.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping database: %w", err)
	}

	if err := db.WithContext(ctx).AutoMigrate(&model.Todo{}, &model.Url{}); err != nil {
		return nil, fmt.Errorf("auto migrate: %w", err)
	}

	return db, nil
}

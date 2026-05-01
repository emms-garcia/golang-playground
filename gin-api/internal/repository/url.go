package repository

import (
	"context"

	"github.com/emms-garcia/golang-playground/gin-api/internal/model"
	"gorm.io/gorm"
)

type UrlRepository struct {
	*baseRepository[model.Url]
	db *gorm.DB
}

// NewUrlRepository is a function to create a new URL repository
func NewUrlRepository(db *gorm.DB) *UrlRepository {
	return &UrlRepository{
		baseRepository: newBaseRepository[model.Url](db),
		db:             db,
	}
}

// CreateUrl is a function to create a new URL entry with the original url and the generated short code
func (r *UrlRepository) CreateUrl(ctx context.Context, original string, shortCode string) (*model.Url, error) {
	url := &model.Url{Original: original, ShortCode: shortCode}
	if err := r.db.WithContext(ctx).Create(url).Error; err != nil {
		return nil, err
	}
	return url, nil
}

// GetUrlByShortCode is a function to get a URL entry by its short code
func (r *UrlRepository) GetUrlByShortCode(ctx context.Context, shortCode string) (*model.Url, error) {
	var url model.Url
	if err := r.db.WithContext(ctx).Where("short_code = ?", shortCode).First(&url).Error; err != nil {
		return nil, err
	}
	return &url, nil
}

// GetUrlByOriginal is a function to get a URL entry by its original URL
func (r *UrlRepository) GetUrlByOriginal(ctx context.Context, original string) (*model.Url, error) {
	var url model.Url
	if err := r.db.WithContext(ctx).Where("original = ?", original).First(&url).Error; err != nil {
		return nil, err
	}
	return &url, nil
}

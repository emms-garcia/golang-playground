package repository

import (
	"github.com/emms-garcia/golang-playground/gin-api/internal/model"
	"gorm.io/gorm"
)

type UrlRepository interface {
	// CreateUrl creates a new URL in the database
	CreateUrl(original string, shortCode string) (*model.Url, error)
	GetUrlByShort(short string) (*model.Url, error)
}

type urlRepository struct {
	Db *gorm.DB
}

func NewUrlRepository(db *gorm.DB) UrlRepository {
	return &urlRepository{Db: db}
}

func (r *urlRepository) CreateUrl(original string, shortCode string) (*model.Url, error) {
	url := &model.Url{Original: original, ShortCode: shortCode}
	if err := r.Db.Create(url).Error; err != nil {
		return nil, err
	}
	return url, nil
}

func (r *urlRepository) GetUrlByShort(shortCode string) (*model.Url, error) {
	var url model.Url
	if err := r.Db.Where("short_code = ?", shortCode).First(&url).Error; err != nil {
		return nil, err
	}
	return &url, nil
}

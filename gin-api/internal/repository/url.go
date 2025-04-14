package repository

import (
	"github.com/emms-garcia/golang-playground/gin-api/internal/model"
	"gorm.io/gorm"
)

// URLRepository is an interface that defines the methods for the URL repository
type UrlRepository interface {
	BaseRepository[model.Url]
	CreateUrl(original string, shortCode string) (*model.Url, error)
	GetUrlByShortCode(short string) (*model.Url, error)
	GetUrlByOriginal(original string) (*model.Url, error)
}

type urlRepository struct {
	BaseRepository[model.Url]
	Db *gorm.DB
}

// NewUrlRepository is a function to create a new URL repository
func NewUrlRepository(db *gorm.DB) UrlRepository {
	return &urlRepository{
		BaseRepository: NewBaseRepository[model.Url](db),
		Db:             db,
	}
}

// CreateUrl is a function to create a new URL entry with the original url and the generated short code
func (r *urlRepository) CreateUrl(original string, shortCode string) (*model.Url, error) {
	url := &model.Url{Original: original, ShortCode: shortCode}
	if err := r.Db.Create(url).Error; err != nil {
		return nil, err
	}
	return url, nil
}

// GetUrlByShortCode is a function to get a URL entry by its short code
func (r *urlRepository) GetUrlByShortCode(shortCode string) (*model.Url, error) {
	var url model.Url
	if err := r.Db.Where("short_code = ?", shortCode).First(&url).Error; err != nil {
		return nil, err
	}
	return &url, nil
}

// GetUrlByOriginal is a function to get a URL entry by its original URL
func (r *urlRepository) GetUrlByOriginal(original string) (*model.Url, error) {
	var url model.Url
	if err := r.Db.Where("original = ?", original).First(&url).Error; err != nil {
		return nil, err
	}
	return &url, nil
}

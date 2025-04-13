package service

import (
	"fmt"
	"math/big"

	"github.com/google/uuid"

	"github.com/emms-garcia/golang-playground/gin-api/internal/model"
	"github.com/emms-garcia/golang-playground/gin-api/internal/repository"
)

var base62 = big.NewInt(62)

const base62Charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

type UrlService interface {
	// CreateUrl creates a new URL
	CreateUrl(original string) (*model.Url, error)
	// GenerateShortCode generates a short code
	GenerateShortCode() string
	// GetUrlByShort retrieves a URL by its short code
	GetUrlByShort(short string) (*model.Url, error)
}

type urlService struct {
	repo repository.UrlRepository
}

func NewUrlService(repo repository.UrlRepository) UrlService {
	return &urlService{repo: repo}
}

func (s *urlService) GenerateShortCode() string {
	u := uuid.New()
	uInt := new(big.Int).SetBytes(u[:])

	var encoded []byte
	zero := big.NewInt(0)
	mod := new(big.Int)

	for uInt.Cmp(zero) > 0 {
		uInt.DivMod(uInt, base62, mod)
		index := mod.Int64()

		// Safety check
		if index < 0 || index >= int64(len(base62Charset)) {
			panic(fmt.Sprintf("Invalid base62 index: %d", index))
		}

		encoded = append([]byte{base62Charset[index]}, encoded...)
	}

	// Optionally truncate to 8 characters for shortness
	if len(encoded) > 8 {
		encoded = encoded[:8]
	}

	return string(encoded)
}

func (s *urlService) CreateUrl(original string) (*model.Url, error) {
	shortCode := s.GenerateShortCode()
	url, err := s.repo.CreateUrl(original, shortCode)
	if err != nil {
		return nil, err
	}
	return url, nil
}

func (s *urlService) GetUrlByShort(shortCode string) (*model.Url, error) {
	url, err := s.repo.GetUrlByShort(shortCode)
	if err != nil {
		return nil, err
	}
	return url, nil
}

package service

import (
	"context"
	"fmt"
	"math/big"

	"github.com/google/uuid"

	"github.com/emms-garcia/golang-playground/gin-api/internal/model"
	"github.com/emms-garcia/golang-playground/gin-api/internal/repository"
)

var base62 = big.NewInt(62)

const base62Charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

type UrlService struct {
	repo *repository.UrlRepository
}

func NewUrlService(repo *repository.UrlRepository) *UrlService {
	return &UrlService{repo: repo}
}

func (s *UrlService) GenerateShortCode() string {
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

func (s *UrlService) CreateUrl(ctx context.Context, original string) (*model.Url, error) {
	shortCode := s.GenerateShortCode()
	url, err := s.repo.CreateUrl(ctx, original, shortCode)
	if err != nil {
		return nil, err
	}

	return url, nil
}

func (s *UrlService) GetUrlByShortCode(ctx context.Context, shortCode string) (*model.Url, error) {
	url, err := s.repo.GetUrlByShortCode(ctx, shortCode)
	if err != nil {
		return nil, err
	}

	url.Usages += 1
	if err := s.repo.Update(ctx, url); err != nil {
		return nil, err
	}

	return url, nil
}

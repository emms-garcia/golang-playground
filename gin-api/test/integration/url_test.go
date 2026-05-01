package integration

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/emms-garcia/golang-playground/gin-api/internal/repository"
	"github.com/stretchr/testify/assert"
)

// TestShortenHandler tests the POST /u/shorten endpoint
func TestShortenHandler(t *testing.T) {
	app := NewTestApplication(t)
	longUrl := "https://github.com"
	response := app.Request("POST", "/u/shorten", fmt.Sprintf("{\"url\": \"%s\"}", longUrl))
	assert.Equal(t, http.StatusOK, response.Code)
	url, err := repository.NewUrlRepository(app.DB).GetUrlByOriginal(context.Background(), longUrl)
	if !assert.NoError(t, err) {
		return
	}

	shortUrl := "http://localhost/u/" + url.ShortCode
	assert.Equal(t, fmt.Sprintf("{\"short\":\"%s\"}", shortUrl), response.Body.String())
}

// TestRedirectHandler tests the GET /u/:short endpoint
func TestRedirectHandler(t *testing.T) {
	app := NewTestApplication(t)
	longUrl := "https://github.com"
	repo := repository.NewUrlRepository(app.DB)
	url, err := repo.CreateUrl(context.Background(), longUrl, "abc123")
	if !assert.NoError(t, err) {
		return
	}

	response := app.Request("GET", "/u/"+url.ShortCode, "")
	url, err = repo.Get(context.Background(), url.ID)
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, http.StatusFound, response.Code)
	assert.Equal(t, longUrl, response.Header().Get("Location"))
	assert.Equal(t, 1, url.Usages)
}

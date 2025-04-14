package integration

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/emms-garcia/golang-playground/gin-api/internal/repository"
	"github.com/stretchr/testify/assert"
)

// TestShortenHandler tests the GET /u/shorten endpoint
func TestShortenHandler(t *testing.T) {
	app := NewTestApplication()
	defer app.Teardown()

	longUrl := "https://github.com"
	response := app.Request("POST", "/u/shorten", fmt.Sprintf("{\"url\": \"%s\"}", longUrl))
	assert.Equal(t, http.StatusOK, response.Code)

	url, err := repository.NewUrlRepository(app.DB).GetUrlByOriginal(longUrl)
	if err != nil {
		t.Error("failed to retrieve url")
		return
	}

	shortUrl := "http://localhost/u/" + url.ShortCode
	assert.Equal(t, fmt.Sprintf("{\"short\":\"%s\"}", shortUrl), response.Body.String())
}

// TestRedirectHandler tests the GET /u/:short endpoint
func TestRedirectHandler(t *testing.T) {
	app := NewTestApplication()
	defer app.Teardown()

	longUrl := "https://github.com"
	repo := repository.NewUrlRepository(app.DB)
	url, err := repo.CreateUrl(longUrl, "abc123")
	if err != nil {
		t.Error("failed to create url")
		return
	}

	response := app.Request("GET", "/u/"+url.ShortCode, "")
	url, err = repo.Get(url.ID)
	if err != nil {
		t.Error("failed to retrieve url")
		return
	}

	assert.Equal(t, http.StatusFound, response.Code)
	assert.Equal(t, longUrl, response.Header().Get("Location"))
	assert.Equal(t, 1, url.Usages)
}

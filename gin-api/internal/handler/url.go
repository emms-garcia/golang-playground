package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/emms-garcia/golang-playground/gin-api/internal/service"
	"github.com/gin-gonic/gin"
)

// UrlHandler struct holds the URL service instance.
type UrlHandler struct {
	service *service.UrlService
}

// NewUrlHandler is a function to create a new URL handler
func NewUrlHandler(service *service.UrlService) *UrlHandler {
	return &UrlHandler{service: service}
}

// ShortenPayload is a struct to hold the request body (JSON) for the ShortenHandler
type ShortenPayload struct {
	Url string `json:"url" binding:"required,url"`
}

// Shorten is a handler function to shorten a URL
func (h *UrlHandler) Shorten(c *gin.Context) {
	var payload ShortenPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	url, err := h.service.CreateUrl(c.Request.Context(), payload.Url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unexpected error while shortening url. please try again later"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"short": buildShortURL(c, url.ShortCode)})
}

// Redirect is a handler function to redirect to the original URL
func (h *UrlHandler) Redirect(c *gin.Context) {
	shortCode := c.Param("short")
	url, err := h.service.GetUrlByShortCode(c.Request.Context(), shortCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.Redirect(http.StatusFound, url.Original)
}

func buildShortURL(c *gin.Context, shortCode string) string {
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}

	if forwardedProto := c.GetHeader("X-Forwarded-Proto"); forwardedProto != "" {
		scheme = strings.TrimSpace(strings.Split(forwardedProto, ",")[0])
	}

	host := c.GetHeader("X-Forwarded-Host")
	if host == "" {
		host = c.Request.Host
	}

	if host == "" {
		host = "localhost"
	}

	return fmt.Sprintf("%s://%s/u/%s", scheme, host, shortCode)
}

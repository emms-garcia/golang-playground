package handler

import (
	"net/http"

	"github.com/emms-garcia/golang-playground/gin-api/internal/service"
	"github.com/gin-gonic/gin"
)

// UrlHandler is an interface that defines the methods for the URL handler
type UrlHandler interface {
	Shorten(c *gin.Context)
	Redirect(c *gin.Context)
}

type urlHandler struct {
	service service.UrlService
}

// NewUrlHandler is a function to create a new URL handler
func NewUrlHandler(service service.UrlService) UrlHandler {
	return &urlHandler{service: service}
}

type ShortenPayload struct {
	Url string `json:"url" binding:"required"`
}

// Shorten is a handler function to shorten a URL
func (h *urlHandler) Shorten(c *gin.Context) {
	var payload ShortenPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url, err := h.service.CreateUrl(payload.Url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"short": url.GetShortUrl()})
}

// Redirect is a handler function to redirect to the original URL
func (h *urlHandler) Redirect(c *gin.Context) {
	shortCode := c.Param("short")
	url, err := h.service.GetUrlByShort(shortCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.Redirect(http.StatusFound, url.Original)
}

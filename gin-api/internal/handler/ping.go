package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingHandler is an interface that defines the methods for the ping handler
type PingHandler interface {
	Ping(c *gin.Context)
}

type pingHandler struct{}

// NewPingHandler is a function to create a new ping handler
func NewPingHandler() PingHandler {
	return &pingHandler{}
}

// Ping is a handler function to respond to the ping endpoint
func (h *pingHandler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

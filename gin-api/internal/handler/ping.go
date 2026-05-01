package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingHandler struct{}

// NewPingHandler is a function to create a new ping handler
func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

// Ping is a handler function to respond to the ping endpoint
func (h *PingHandler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

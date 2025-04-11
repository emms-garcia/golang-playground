package integration

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPingHandler tests the GET /ping endpoint
func TestPingHandler(t *testing.T) {
	ctx := setup()
	defer teardown(ctx)

	response := request(ctx, "GET", "/ping", "")
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "{\"message\":\"pong\"}", response.Body.String())
}

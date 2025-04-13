package integration

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPingHandler tests the GET /ping endpoint
func TestPingHandler(t *testing.T) {
	app := NewTestApplication()
	defer app.Teardown()

	response := app.Request("GET", "/ping", "")
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "{\"message\":\"pong\"}", response.Body.String())
}

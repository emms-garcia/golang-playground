package integration

import (
	"testing"

	"github.com/emms-garcia/golang-playground/gin-api/internal/configuration"
	"github.com/stretchr/testify/assert"
)

// TestConfigureDB tests the database connection is established
func TestConfigureDB(t *testing.T) {
	config, err := configuration.Load(configuration.Test)
	if err != nil {
		t.Errorf("failed to load config: %v", err)
	}
	db, err := configuration.ConfigureDB(config)
	if err != nil {
		t.Errorf("failed to load db: %v", err)
	}
	var one int
	result := db.Raw("SELECT 1").Scan(&one)
	if result.Error != nil {
		t.Error("failed to connect with db")
	}
	assert.Equal(t, 1, one)
}

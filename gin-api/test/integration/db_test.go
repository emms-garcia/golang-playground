package integration

import (
	"testing"

	"github.com/emms-garcia/golang-playground/gin-api/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestConfigureDB(t *testing.T) {
	err := config.Load(config.Test)
	if err != nil {
		t.Errorf("failed to load config: %v", err)
	}
	db := config.ConfigureDB(&config.AppConfiguration)
	var one int
	result := db.Raw("SELECT 1").Scan(&one)
	if result.Error != nil {
		t.Error("failed to connect with db")
	}
	assert.Equal(t, 1, one)
}

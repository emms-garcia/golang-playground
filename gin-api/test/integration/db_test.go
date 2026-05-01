package integration

import (
	"context"
	"testing"

	"github.com/emms-garcia/golang-playground/gin-api/internal/configuration"
	"github.com/stretchr/testify/assert"
)

// TestConfigureDB tests the database connection is established
func TestConfigureDB(t *testing.T) {
	config, err := configuration.Load(testConfigsPath(t), configuration.Test)
	if !assert.NoError(t, err) {
		return
	}

	db, err := configuration.ConfigureDB(context.Background(), config)
	if err != nil {
		t.Skipf("skipping integration test: %v", err)
	}

	var one int
	result := db.WithContext(context.Background()).Raw("SELECT 1").Scan(&one)
	if !assert.NoError(t, result.Error) {
		return
	}

	assert.Equal(t, 1, one)
}

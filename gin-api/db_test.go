package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigureDB(t *testing.T) {
	db := ConfigureDB(&TestConfiguration)
	var one int
	result := db.Raw("SELECT 1").Scan(&one)
	if result.Error != nil {
		t.Error("failed to connect with db")
	}
	assert.Equal(t, 1, one)
}

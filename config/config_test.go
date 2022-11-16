package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	err := LoadConfig()
	assert.Nil(t, err)
	assert.NotNil(t, config)
	assert.NotNil(t, Database())
}

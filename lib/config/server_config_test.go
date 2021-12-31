package config_test

import (
	. "github.com/pulkit-tanwar/omh-users-management/lib/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig(t *testing.T) {
	config := NewConfig("test", "local", 3040, "someApiPath")
	assert.Equal(t, "test", config.Env)
	assert.Equal(t, "local", config.Host)
	assert.Equal(t, 3040, config.Port)
	assert.Equal(t, "someApiPath", config.APIPath)
}

func TestDefaultConfig(t *testing.T) {
	config := DefaultConfig()
	assert.Equal(t, "dev", config.Env)
	assert.Equal(t, "localhost", config.Host)
	assert.Equal(t, 3000, config.Port)
	assert.Equal(t, "/", config.APIPath)
}

package server

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestConfig_Default(t *testing.T) {
	err, cancel := setEnv("APPLICATION_NAME", "todo")
	require.NoError(t, err)
	defer cancel()

	err, cancel = setEnv("DATABASE_NAME", "todo_db")
	require.NoError(t, err)
	defer cancel()

	cfg, err := ReadConfig()
	assert.NoError(t, err)
	assert.NotZero(t, cfg)
	assert.Equal(t, 3030, cfg.Port)
	assert.Equal(t, "todo", cfg.Name)
	assert.Equal(t, "todo_db", cfg.DB.Name)
}

func setEnv(key string, value string) (error, func()) {
	return os.Setenv(key, value), func() {
		_ = os.Unsetenv(key)
	}
}

package config

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetVaultPath_FromConfig(t *testing.T) {
	t.Parallel()

	v := viper.New()
	v.Set("vaultPath", "/tmp/from-config")

	path, err := GetVaultPath(v)
	require.NoError(t, err)
	assert.Equal(t, "/tmp/from-config", path)
}

func TestGetVaultPath_FromEnv(t *testing.T) {
	t.Setenv("OBS_VAULT_PATH", "/tmp/from-env")

	v := viper.New()
	v.Set("vaultPath", "/tmp/from-config")

	path, err := GetVaultPath(v)
	require.NoError(t, err)
	assert.Equal(t, "/tmp/from-env", path)
}

func TestGetVaultPath_Missing(t *testing.T) {
	t.Parallel()

	v := viper.New()

	_, err := GetVaultPath(v)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "vault path not configured")
}

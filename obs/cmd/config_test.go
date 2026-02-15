package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestNewConfigCmd(t *testing.T) {
	t.Parallel()
	tempDir := t.TempDir()
	configDir := filepath.Join(tempDir, ".config", "obsidian-cli")
	configFile := filepath.Join(configDir, "config.yaml")

	// create the config directory
	err := os.MkdirAll(configDir, 0750)
	assert.NoError(t, err)

	// Clean up any existing config file after the test
	t.Cleanup(func() {
		err = os.RemoveAll(configDir)
		assert.NoError(t, err)
	})

	cmd := NewConfigCmd(tempDir)
	cmd.SetArgs([]string{"--vault", "testVault", "--targetFolder", "testFolder"})

	err = cmd.Execute()
	assert.NoError(t, err)

	vault := viper.GetString("vault")
	targetFolder := viper.GetString("targetFolder")

	assert.Equal(t, "testVault", vault)
	assert.Equal(t, "testFolder", targetFolder)

	// Check if the config file is created
	_, err = os.Stat(configFile)
	assert.NoError(t, err)
}

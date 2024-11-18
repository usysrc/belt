package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestNewConfigCmd(t *testing.T) {
	homeDir := os.Getenv("HOME")
	configDir := filepath.Join(homeDir, ".config", "obsidian-cli")
	configFile := filepath.Join(configDir, "config.yaml")

	// Clean up before and after the test
	defer os.RemoveAll(configDir)

	cmd := NewConfigCmd()
	cmd.SetArgs([]string{"--vault", "testVault", "--targetFolder", "testFolder"})

	err := cmd.Execute()
	assert.NoError(t, err)

	vault := viper.GetString("vault")
	targetFolder := viper.GetString("targetFolder")

	assert.Equal(t, "testVault", vault)
	assert.Equal(t, "testFolder", targetFolder)

	// Check if the config file is created
	_, err = os.Stat(configFile)
	assert.NoError(t, err)
}

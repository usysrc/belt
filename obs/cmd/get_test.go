package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewGetCmd_Metadata(t *testing.T) {
	t.Parallel()

	cfg := viper.New()
	cmd := NewGetCmd(cfg)

	assert.Equal(t, "get", cmd.Use)
	assert.Equal(t, "Print note content to stdout", cmd.Short)
	assert.Equal(t, "Read a note from your Obsidian vault and print its content to stdout.", cmd.Long)
	assert.Equal(t, "stdio", cmd.Annotations["mode"])
}

func TestNewGetCmd_MissingVaultPath(t *testing.T) {
	t.Parallel()

	cfg := viper.New()
	cmd := NewGetCmd(cfg)
	cmd.SetArgs([]string{"note"})

	err := cmd.Execute()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "vault path not configured")
}

func TestNewGetCmd_NotFound(t *testing.T) {
	t.Parallel()

	cfg := viper.New()
	cfg.Set("vaultPath", t.TempDir())

	cmd := NewGetCmd(cfg)
	cmd.SetArgs([]string{"missing"})

	err := cmd.Execute()
	require.Error(t, err)
	assert.Contains(t, err.Error(), `note "missing" not found`)
}

func TestNewGetCmd_PrintsContentByBaseName(t *testing.T) {
	t.Parallel()

	vault := t.TempDir()
	notePath := filepath.Join(vault, "topic.md")
	require.NoError(t, os.WriteFile(notePath, []byte("hello world"), 0600))

	cfg := viper.New()
	cfg.Set("vaultPath", vault)
	cmd := NewGetCmd(cfg)
	cmd.SetArgs([]string{"topic"})

	out := &bytes.Buffer{}
	cmd.SetOut(out)

	err := cmd.Execute()
	require.NoError(t, err)
	assert.Equal(t, "hello world", out.String())
}

func TestNewGetCmd_PrintsContentByRelativePath(t *testing.T) {
	t.Parallel()

	vault := t.TempDir()
	require.NoError(t, os.MkdirAll(filepath.Join(vault, "notes"), 0750))
	notePath := filepath.Join(vault, "notes", "nested.md")
	require.NoError(t, os.WriteFile(notePath, []byte("nested content"), 0600))

	cfg := viper.New()
	cfg.Set("vaultPath", vault)
	cmd := NewGetCmd(cfg)
	cmd.SetArgs([]string{"notes/nested"})

	out := &bytes.Buffer{}
	cmd.SetOut(out)

	err := cmd.Execute()
	require.NoError(t, err)
	assert.Equal(t, "nested content", out.String())
}

package cmd

import (
	"bytes"
	"fmt"
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

//nolint:paralleltest // Mutates package-level fzfLookPath stub for deterministic fallback behavior.
func TestNewGetCmd_DuplicateUsesAskFallback(t *testing.T) {
	oldLookPath := fzfLookPath

	t.Cleanup(func() {
		fzfLookPath = oldLookPath
	})

	fzfLookPath = func(file string) (string, error) {
		return "", fmt.Errorf("%s not found", file)
	}

	vault := t.TempDir()
	require.NoError(t, os.MkdirAll(filepath.Join(vault, "a"), 0750))
	require.NoError(t, os.MkdirAll(filepath.Join(vault, "b"), 0750))
	require.NoError(t, os.WriteFile(filepath.Join(vault, "a", "dup.md"), []byte("first"), 0600))
	require.NoError(t, os.WriteFile(filepath.Join(vault, "b", "dup.md"), []byte("second"), 0600))

	cfg := viper.New()
	cfg.Set("vaultPath", vault)
	cmd := NewGetCmd(cfg)
	cmd.SetArgs([]string{"dup"})

	in := bytes.NewBufferString("2\n")
	out := &bytes.Buffer{}
	errOut := &bytes.Buffer{}

	cmd.SetIn(in)
	cmd.SetOut(out)
	cmd.SetErr(errOut)

	err := cmd.Execute()
	require.NoError(t, err)
	assert.Equal(t, "second", out.String())
	assert.Contains(t, errOut.String(), "Multiple notes matched:")
}

//nolint:paralleltest // Mutates package-level fzfLookPath stub for deterministic fallback behavior.
func TestNewGetCmd_DuplicateAskFallbackInvalidSelection(t *testing.T) {
	oldLookPath := fzfLookPath

	t.Cleanup(func() {
		fzfLookPath = oldLookPath
	})

	fzfLookPath = func(file string) (string, error) {
		return "", fmt.Errorf("%s not found", file)
	}

	vault := t.TempDir()
	require.NoError(t, os.MkdirAll(filepath.Join(vault, "a"), 0750))
	require.NoError(t, os.MkdirAll(filepath.Join(vault, "b"), 0750))
	require.NoError(t, os.WriteFile(filepath.Join(vault, "a", "dup.md"), []byte("first"), 0600))
	require.NoError(t, os.WriteFile(filepath.Join(vault, "b", "dup.md"), []byte("second"), 0600))

	cfg := viper.New()
	cfg.Set("vaultPath", vault)
	cmd := NewGetCmd(cfg)
	cmd.SetArgs([]string{"dup"})
	cmd.SetIn(bytes.NewBufferString("3\n"))

	err := cmd.Execute()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "selection out of range")
}

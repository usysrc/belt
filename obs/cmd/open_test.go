package cmd

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestNewOpenCmd(t *testing.T) {
	t.Parallel()

	cfg := viper.New()
	cfg.Set("vault", "testVault")
	cfg.Set("targetFolder", "./testFolder")
	cmd := NewOpenCmd(cfg)

	assert.Equal(t, "open", cmd.Use)
	assert.Equal(t, "Open an existing note in Obsidian (GUI required)", cmd.Short)
	assert.Equal(t, `Open an existing note in your Obsidian vault using the specified file name (GUI required).`, cmd.Long)
	assert.Equal(t, "gui", cmd.Annotations["mode"])

	cmd.SetArgs([]string{""})

	err := cmd.Execute()
	if err == nil {
		t.Fatalf("expected err got nil")
	}

	if err.Error() != "note name cannot be empty" {
		t.Fatalf("expected \"%s\" got \"%s\"", "note name cannot be empty", err.Error())
	}
}

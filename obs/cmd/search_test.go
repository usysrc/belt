package cmd

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewSearchCmd(t *testing.T) {
	t.Parallel()

	cfg := viper.New()
	cfg.Set("vault", "testVault")
	cfg.Set("targetFolder", "./testFolder")
	cmd := NewSearchCmd(cfg)

	assert.Equal(t, "search", cmd.Use)
	assert.Equal(t, "Search in vault", cmd.Short)
	assert.Equal(t, `Search for content in your Obsidian vault using the specified query.`, cmd.Long)

	cmd.SetArgs([]string{""})

	err := cmd.Execute()
	if err == nil {
		t.Fatalf("expected err got nil")
	}

	if err.Error() != "query can not be empty" {
		t.Fatalf("expected \"%s\" got \"%s\"", "query can not be empty", err.Error())
	}
}

func TestNewSearchCmd_NoArgs(t *testing.T) {
	t.Parallel()

	cfg := viper.New()
	cfg.Set("vault", "testVault")
	cmd := NewSearchCmd(cfg)
	cmd.SetArgs([]string{})

	err := cmd.Execute()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "accepts 1 arg(s), received 0")
}

package cmd

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestNewOpenCmd(t *testing.T) {
	viper.Set("vault", "testVault")
	viper.Set("targetFolder", "./testFolder")
	cmd := NewOpenCmd()

	assert.Equal(t, "open", cmd.Use)
	assert.Equal(t, "Open an existing note", cmd.Short)
	assert.Equal(t, `Open an existing note in your Obsidian vault using the specified file name.`, cmd.Long)

	cmd.SetArgs([]string{""})
	err := cmd.Execute()
	if err == nil {
		t.Fatalf("expected err got nil")
	}
	if err != nil && err.Error() != "note name cannot be empty" {
		t.Fatalf("expected \"%s\" got \"%s\"", "note name cannot be empty", err.Error())
	}
}

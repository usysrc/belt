package cmd

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestNewCreateCmd(t *testing.T) {
	viper.Set("vault", "testVault")
	viper.Set("targetFolder", "./testFolder")
	cmd := NewCreateCmd()

	assert.Equal(t, "create", cmd.Use)
	assert.ElementsMatch(t, []string{"new", "add"}, cmd.Aliases)
	assert.Equal(t, "Create a new note", cmd.Short)
	assert.Equal(t, `Create a new note in your Obsidian vault using the specified file name.`, cmd.Long)

	cmd.SetArgs([]string{""})
	err := cmd.Execute()
	if err == nil {
		t.Fatalf("expected err got nil")
	}
	if err != nil && err.Error() != "note name cannot be empty" {
		t.Fatalf("expected \"%s\" got \"%s\"", "note name cannot be empty", err.Error())
	}
}

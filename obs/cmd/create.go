package cmd

import (
	"fmt"
	"io"
	"os"

	"codeberg.org/usysrc/belt/obs/internal/config"
	"codeberg.org/usysrc/belt/obs/internal/uri"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewCreateCmd(cfg *viper.Viper) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create",
		Aliases: []string{"new", "add"},
		Short:   "Create a new note",
		Long:    `Create a new note in your Obsidian vault using the specified file name.`,
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			vault, err := config.GetVault(cfg)
			if err != nil {
				return fmt.Errorf("failed to get vault from config: %w", err)
			}

			targetFolder, err := config.GetTargetFolder(cfg)
			if err != nil {
				return fmt.Errorf("failed to get target folder from config: %w", err)
			}

			noteName := args[0]
			if noteName == "" {
				return fmt.Errorf("note name cannot be empty")
			}

			// Arguments take precedence over stdin
			if len(args) > 1 {
				noteContent := args[1]

				return uri.Execute("new", vault, noteName, targetFolder, noteContent)
			}

			note, err := io.ReadAll(os.Stdin)
			if err != nil {
				return fmt.Errorf("failed to read note content from stdin: %w", err)
			}

			noteContent := string(note)

			return uri.Execute("new", vault, noteName, targetFolder, noteContent)
		},
	}

	return cmd
}

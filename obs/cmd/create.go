package cmd

import (
	"fmt"
	"io"
	"os"

	"codeberg.org/usysrc/belt/obs/internal/config"
	"codeberg.org/usysrc/belt/obs/internal/uri"
	"github.com/spf13/cobra"
)

func NewCreateCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "create",
		Aliases: []string{"new", "add"},
		Short:   "Create a new note",
		Long:    `Create a new note in your Obsidian vault using the specified file name.`,
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			vault, err := config.GetVault()
			if err != nil {
				return err
			}
			targetFolder, err := config.GetTargetFolder()
			if err != nil {
				return err
			}
			noteName := args[0]
			if noteName == "" {
				return fmt.Errorf("note name cannot be empty")
			}

			// Arguments take precedence over stdin
			noteContent := ""
			if len(args) > 1 {
				noteContent = args[1]
			} else {
				noteContents, err := io.ReadAll(os.Stdin)
				if err != nil {
					return err
				}
				noteContent = string(noteContents)
			}
			return uri.Execute("new", vault, noteName, targetFolder, noteContent)
		},
	}

	return cmd
}

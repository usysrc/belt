package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/usysrc/obs/internal/config"
	"github.com/usysrc/obs/internal/uri"
)

func NewCreateCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new note",
		Long:  `Create a new note in your Obsidian vault using the specified file name.`,
		Args:  cobra.ExactArgs(1),
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
			return uri.Execute("new", vault, noteName, targetFolder)
		},
	}

	return cmd
}

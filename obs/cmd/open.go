package cmd

import (
	"fmt"

	"codeberg.org/usysrc/belt/obs/internal/config"
	"codeberg.org/usysrc/belt/obs/internal/uri"
	"github.com/spf13/cobra"
)

func NewOpenCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "open",
		Short: "Open an existing note",
		Long:  `Open an existing note in your Obsidian vault using the specified file name.`,
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
			return uri.Execute("open", vault, noteName, targetFolder, "")
		},
	}

	return cmd
}

package cmd

import (
	"fmt"

	"codeberg.org/usysrc/belt/obs/internal/config"
	"codeberg.org/usysrc/belt/obs/internal/uri"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewOpenCmd(cfg *viper.Viper) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "open",
		Short: "Open an existing note in Obsidian (GUI required)",
		Long:  `Open an existing note in your Obsidian vault using the specified file name (GUI required).`,
		Args:  cobra.ExactArgs(1),
		Annotations: map[string]string{
			"mode": "gui",
		},
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

			return uri.Open(vault, noteName, targetFolder)
		},
	}

	return cmd
}

package cmd

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewConfigCmd(basePath string, cfg *viper.Viper) *cobra.Command {
	var vault string

	var targetFolder string

	var vaultPath string

	cmd := &cobra.Command{
		Use:   "config",
		Short: "Configure the CLI",
		Long:  `Configure the CLI with your Obsidian vault name and other settings.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg.Set("vault", vault)
			cfg.Set("targetFolder", targetFolder)

			if cmd.Flags().Changed("vaultPath") {
				cfg.Set("vaultPath", vaultPath)
			}

			configFile := filepath.Join(
				basePath,
				".config",
				"obsidian-cli",
				"config.yaml",
			)
			if err := cfg.WriteConfigAs(configFile); err != nil {
				return fmt.Errorf("failed to write config: %w", err)
			}

			fmt.Printf(
				"Configuration saved: vault = %s, targetFolder = %s, vaultPath = %s\n",
				vault,
				targetFolder,
				cfg.GetString("vaultPath"),
			)

			return nil
		},
	}

	cmd.Flags().StringVarP(&vault, "vault", "v", "", "Name of the Obsidian vault")

	if err := cmd.MarkFlagRequired("vault"); err != nil {
		log.Fatal("Error marking flag as required:", err)
	}

	cmd.Flags().StringVarP(&targetFolder, "targetFolder", "t", "", "Folder where new notes are created")

	if err := cmd.MarkFlagRequired("targetFolder"); err != nil {
		log.Fatal("Error marking flag as required:", err)
	}

	cmd.Flags().StringVar(
		&vaultPath,
		"vaultPath",
		"",
		"Absolute path to the Obsidian vault on disk (used by non-GUI commands like get)",
	)

	return cmd
}

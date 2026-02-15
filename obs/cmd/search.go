package cmd

import (
	"fmt"

	"codeberg.org/usysrc/belt/obs/internal/config"
	"codeberg.org/usysrc/belt/obs/internal/uri"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewSearchCmd(cfg *viper.Viper) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "search",
		Short: "Search in vault",
		Long:  `Search for content in your Obsidian vault using the specified query.`,
		Args:  cobra.ExactArgs(1),

		RunE: func(cmd *cobra.Command, args []string) error {
			vault, err := config.GetVault(cfg)
			if err != nil {
				return fmt.Errorf("failed to get vault from config: %w", err)
			}

			query := args[0]
			if query == "" {
				return fmt.Errorf("query can not be empty")
			}

			return uri.Execute("search", vault, query, "", "")
		},
	}

	return cmd
}

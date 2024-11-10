package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/usysrc/obs/internal/config"
	"github.com/usysrc/obs/internal/uri"
)

func NewSearchCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "search",
		Short: "Search in vault",
		Long:  `Search for content in your Obsidian vault using the specified query.`,

		RunE: func(cmd *cobra.Command, args []string) error {
			vault, err := config.GetVault()
			if err != nil {
				return err
			}
			query := args[0]
			if query == "" {
				return fmt.Errorf("query can not be empty")
			}
			return uri.Execute("search", vault, query, "")
		},
	}

	return cmd
}

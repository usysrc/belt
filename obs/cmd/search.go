package cmd

import (
	"fmt"

	"codeberg.org/usysrc/belt/obs/internal/config"
	"codeberg.org/usysrc/belt/obs/internal/uri"
	"github.com/spf13/cobra"
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
			return uri.Execute("search", vault, query, "", "")
		},
	}

	return cmd
}

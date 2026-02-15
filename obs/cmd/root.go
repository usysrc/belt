package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewRootCmd(basePath string, cfg *viper.Viper) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "obs",
		Short: "A CLI for interacting with Obsidian",
		Long: `A Command Line Interface for Obsidian that uses the Obsidian URI scheme 
to perform various operations like creating notes, opening notes, and searching.`,
	}

	cmd.AddCommand(
		NewConfigCmd(basePath, cfg),
		NewCreateCmd(cfg),
		NewOpenCmd(cfg),
		NewSearchCmd(cfg),
	)

	return cmd
}

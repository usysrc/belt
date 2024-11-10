package cmd

import (
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "obs",
		Short: "A CLI for interacting with Obsidian",
		Long: `A Command Line Interface for Obsidian that uses the Obsidian URI scheme 
to perform various operations like creating notes, opening notes, and searching.`,
	}

	cmd.AddCommand(
		NewConfigCmd(),
		NewCreateCmd(),
		NewOpenCmd(),
		NewSearchCmd(),
	)

	return cmd
}

package cmd

import (
	"os"

	view "codeberg.org/usysrc/belt/hex/internal/viewer"
	"github.com/spf13/cobra"
)

var width int

func run(cmd *cobra.Command, args []string) {
	if len(args) < 1 || args[0] == "" {
		if err := cmd.Usage(); err != nil {
			panic(err)
		}

		os.Exit(1)

		return
	}

	view.CreateView(args[0], width)
}

var rootCmd = &cobra.Command{
	Use:   "hex <filename>",
	Short: "View a file in hex format in a TUI.",
	Long:  `This is a simple hex viewer that displays a file in hex format in a TUI.`,
	Run:   run,
}

// Execute executes the root command.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// init initializes the root command.
func init() {
	// The number of bytes per line
	rootCmd.Flags().IntVar(&width, "width", 16, "bytes per line")
}

package cmd

import (
	"os"

	"github.com/spf13/cobra"
	view "github.com/usysrc/hex/internal/viewer"
)

var width int

var rootCmd = &cobra.Command{
	Use:   "hex",
	Short: "View a file in hex format in a TUI.",
	Long:  `This is a simple hex viewer that displays a file in hex format in a TUI.`,
	Run: func(cmd *cobra.Command, args []string) {
		view.CreateView(args[0], width)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
func init() {
	// The number of bytes per line
	rootCmd.Flags().IntVar(&width, "width", 16, "bytes per line")
}

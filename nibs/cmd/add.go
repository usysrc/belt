package cmd

import (
	"log"

	"codeberg.org/usysrc/belt/nibs/battery"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [file or directory]",
	Short: "Add a file or directory to the project",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		whatToAdd := args[0]
		switch whatToAdd {
		case "hump":
			battery.Hump()
		case "pico":
			battery.Pico()
		default:
			log.Fatalf("Unknown item to add: %s", whatToAdd)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [file or directory]",
	Short: "Add a file or directory to the project",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		whatToAdd := args[0]
		if whatToAdd == "hump" {
			// Repository URL
			// Repository URL
			repoURL := "https://github.com/vrld/hump.git"

			// Current directory
			currentDir, err := os.Getwd()
			if err != nil {
				log.Fatalf("Failed to get current directory: %v", err)
			}

			// Directory to clone the repository to
			humpDir := filepath.Join(currentDir, "hump")

			// Clone the repository
			_, err = git.PlainClone(humpDir, false, &git.CloneOptions{
				URL:      repoURL,
				Depth:    1, // Only fetch the latest commit
				Progress: os.Stdout,
			})

			if err != nil {
				log.Fatalf("Failed to clone repository: %v", err)
			}

			fmt.Println("Successfully cloned vrld/hump repository")

			// Add your logic here
		} else {
			log.Fatalf("Unknown item to add: %s", whatToAdd)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

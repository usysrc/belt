package battery

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
)

func Hump() {
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
}

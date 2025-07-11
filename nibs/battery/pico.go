package battery

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func Pico() {
	fileURL := "https://codeberg.org/usysrc/labs/raw/branch/main/pico/pico.lua"
	fileName := "pico.lua"

	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current directory: %v", err)
	}

	// Create the full path for the new file
	filePath := filepath.Join(currentDir, fileName)

	// Create the file
	out, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Failed to create file %s: %v", filePath, err)
	}
	defer out.Close() // Ensure the file is closed

	// Get the data from the URL
	resp, err := http.Get(fileURL)
	if err != nil {
		log.Fatalf("Failed to download %s: %v", fileURL, err)
	}
	defer resp.Body.Close() // Ensure the response body is closed

	// Check server response
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Bad status: %s", resp.Status)
	}

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatalf("Failed to write content to file %s: %v", filePath, err)
	}

	fmt.Printf("Successfully downloaded %s to %s\n", fileURL, filePath)
}

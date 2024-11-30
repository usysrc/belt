package cmd

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var bundleCmd = &cobra.Command{
	Use:   "bundle",
	Short: "bundle the project into a .love file",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dir := "./"
		outputFile := getOutputFile(cmd)
		bundleProject(dir, outputFile)
	},
}

func init() {
	// add -o flag to specify output file
	bundleCmd.Flags().StringP("output", "o", "game.love", "output file")
	// add bundle command to root command
	rootCmd.AddCommand(bundleCmd)
}

func getOutputFile(cmd *cobra.Command) string {
	// default set name of output file to directory name
	currentDirectory, err := os.Getwd()
	if err != nil {
		currentDirectory = "game"
	}
	outputFile := filepath.Base(currentDirectory) + ".love"

	// read the output flag if it was set
	if cmd.Flags().Changed("output") {
		outputFile, _ = cmd.Flags().GetString("output")
	}
	return outputFile
}

func bundleProject(dir, outputFile string) {
	log.Println("Bundling project...")
	out, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Failed to create .love file: %v", err)
	}
	defer out.Close()

	archive := zip.NewWriter(out)
	defer archive.Close()

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories and unnecessary files
		if info.IsDir() || shouldIgnore(path) {
			return nil
		}

		// Add file to the archive
		relPath, err := filepath.Rel(dir, path)
		if err != nil {
			return err
		}
		return addFileToZip(archive, path, relPath)
	})

	if err != nil {
		log.Fatalf("Failed to bundle project: %v", err)
	}
	log.Printf("Project bundled as %s", outputFile)
}

func shouldIgnore(path string) bool {
	// Example: Ignore git and temporary files
	ignorePatterns := []string{".git", ".DS_Store", "~", ".swp"}
	for _, pattern := range ignorePatterns {
		if strings.Contains(path, pattern) {
			return true
		}
	}
	return false
}

func addFileToZip(archive *zip.Writer, path, relPath string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer, err := archive.Create(relPath)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, file)
	return err
}

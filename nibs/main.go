package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)

var (
	cmd        *exec.Cmd
	mu         sync.Mutex
	done       = make(chan bool)
	running    = false
	lastBundle time.Time
)

const debounceDuration = 500 * time.Millisecond

func main() {
	dirToWatch := "./" // Change to your project directory
	lovePath := "love"
	outputFile := "game.love"

	startLove(lovePath, outputFile)

	// Initialize watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("Failed to create watcher: %v", err)
	}
	defer watcher.Close()

	go watchFiles(watcher, dirToWatch, lovePath, outputFile)

	if err := addSubdirectories(watcher, dirToWatch); err != nil {
		log.Fatalf("Failed to add directories: %v", err)
	}

	<-done
}

func watchFiles(watcher *fsnotify.Watcher, dir, lovePath, outputFile string) {
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if isRelevantChange(event) {
				log.Printf("Change detected: %s %s", event.Name, event.Op)
				// only bundle and start LÖVE if there hasn't been a change in the last 500ms
				// and the change is a WRITE
				if time.Since(lastBundle) > debounceDuration && event.Op&fsnotify.Write == fsnotify.Write {
					lastBundle = time.Now()
					bundleProject(dir, outputFile)
					startLove(lovePath, outputFile)
				}
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Printf("Watcher error: %v", err)
		}
	}
}

func addSubdirectories(watcher *fsnotify.Watcher, root string) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		// ignore directories that start with dot but not the root directory
		if strings.HasPrefix(info.Name(), ".") && path != root {
			return filepath.SkipDir
		}
		if err != nil {
			return err
		}
		if info.IsDir() {
			return watcher.Add(path)
		}
		return nil
	})
}

func isRelevantChange(event fsnotify.Event) bool {
	extensions := []string{".lua", ".png", ".jpg", ".ogg", ".wav", ".frag", ".vert"}
	for _, ext := range extensions {
		if strings.HasSuffix(event.Name, ext) {
			return true
		}
	}
	return false
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

func startLove(lovePath, outputFile string) {
	mu.Lock()
	defer mu.Unlock()

	log.Println("Attempting to start LÖVE...")

	for running {
		log.Println("Stopping previous LÖVE...")
		if err := cmd.Process.Kill(); err != nil {
			log.Printf("Failed to stop LÖVE: %v", err)
		} else if err := cmd.Wait(); err != nil {
			log.Println("Previous LÖVE stopped successfully")
		}
	}

	log.Println("Starting LÖVE with bundled project...")
	cmd = exec.Command(lovePath, outputFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		log.Printf("Failed to start LÖVE: %v", err)
		return
	}

	running = true
	log.Println("LÖVE started successfully")

	go func() {
		if err := cmd.Wait(); err != nil && !strings.Contains(err.Error(), "signal: killed") {
			log.Printf("LÖVE exited: %v", err)
		}
		running = false
		log.Println("LÖVE stopped")
	}()
}

package cmd

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
)

var (
	cmd        *exec.Cmd
	done       = make(chan bool)
	mu         sync.Mutex
	running    = false
	lastBundle time.Time
)

const debounceDuration = 500 * time.Millisecond

func isRelevantChange(event fsnotify.Event) bool {
	extensions := []string{".lua", ".png", ".jpg", ".ogg", ".wav", ".frag", ".vert"}
	for _, ext := range extensions {
		if strings.HasSuffix(event.Name, ext) {
			return true
		}
	}
	return false
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

var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watch the project directory, bundle and run LÖVE when changes are detected",
	Run: func(cmd *cobra.Command, args []string) {
		dirToWatch := "./" // Change to your project directory
		lovePath := "love"
		outputFile := getOutputFile(cmd)

		// Bundle project and start LÖVE
		bundleProject(dirToWatch, outputFile)

		// Start LÖVE
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
	},
}

func init() {
	// add -o flag to specify output file
	watchCmd.Flags().StringP("output", "o", "game.love", "output file")
	// add watch command
	rootCmd.AddCommand(watchCmd)
}

func startLove(lovePath, outputFile string) {
	mu.Lock()
	defer mu.Unlock()

	log.Println("Attempting to start LÖVE2D...")

	if running {
		log.Println("Stopping LÖVE2D...")
		if err := cmd.Process.Kill(); err != nil {
			log.Printf("Failed to stop LÖVE2D: %v", err)
		} else {
			cmd.Wait()
			log.Println("LÖVE2D stopped successfully")
		}
		running = false
	}
	// wait until the process is killed
	for running {
		time.Sleep(100 * time.Millisecond)
	}

	log.Println("Starting LÖVE2D with bundled project...")
	cmd = exec.Command(lovePath, outputFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		log.Printf("Failed to start LÖVE2D: %v", err)
		return
	}

	running = true
	log.Println("LÖVE2D started successfully")

	go func() {
		if err := cmd.Wait(); err != nil && !strings.Contains(err.Error(), "signal: killed") {
			log.Printf("LÖVE2D exited: %v", err)
		}
		mu.Lock()
		running = false
		mu.Unlock()
		log.Println("LÖVE2D stopped")
	}()
}

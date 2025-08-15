package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func isStdinFromTerminal() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) != 0
}

func slowPrint(lines []string, delay time.Duration) error {
	for _, line := range lines {
		if _, err := fmt.Println(line); err != nil {
			if err == syscall.EPIPE {
				return nil // Exit gracefully on broken pipe
			}
			return err
		}
		time.Sleep(delay)
	}
	return nil
}

func main() {
	// Define and parse the delay flag
	delay := flag.Duration("delay", 750*time.Millisecond, "delay between lines in milliseconds")
	flag.Parse()

	// Set up signal handling
	signal.Ignore(syscall.SIGPIPE)

	var lines []string
	reader := bufio.NewReader(os.Stdin)

	if isStdinFromTerminal() {
		fmt.Println("Enter lines (press Ctrl+D to finish):")
	}

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
				os.Exit(1)
			}
			break
		}
		lines = append(lines, line[:len(line)-1])
	}

	if err := slowPrint(lines, *delay); err != nil {
		fmt.Fprintf(os.Stderr, "Error printing: %v\n", err)
		os.Exit(1)
	}
}

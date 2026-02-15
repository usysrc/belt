package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	log.Printf("Do you want to proceed? (y/N): ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if strings.ToLower(input) == "y" {
		// Return exit code 0 for "yes"
		os.Exit(0)
	} else {
		// Return exit code 1 for "no"
		os.Exit(1)
	}
}

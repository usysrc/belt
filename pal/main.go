package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Read hex values from stdin (each on a new line)
	scanner := bufio.NewScanner(os.Stdin)
	var hexValues []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		hexValues = append(hexValues, line)
	}

	// Create the palette string
	paletteString := "000:"
	for _, hexValue := range hexValues {
		paletteString += strings.TrimSpace(hexValue)
	}

	fmt.Println(paletteString)
}

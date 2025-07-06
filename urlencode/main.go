package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

func urlEncode(input string) string {
	return url.QueryEscape(input)
}

func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("Usage: urlencode <string to encode>...")
	}

	input := strings.Join(args[1:], " ")
	encoded := urlEncode(input)
	fmt.Println(encoded)
	return nil
}

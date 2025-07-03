//go:build !darwin && !linux && !windows

package main

import "fmt"

func openURL(url string) error {
	return fmt.Errorf("unsupported operating system")
}

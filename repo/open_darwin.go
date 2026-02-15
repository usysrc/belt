//go:build darwin

package main

import (
	"context"
	"fmt"
	"os/exec"
)

func openURL(url string) error {
	cmd := exec.CommandContext(context.Background(), "open", url)

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to open URL: %w", err)
	}

	return nil
}

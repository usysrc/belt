//go:build linux

package main

import "os/exec"

func openURL(url string) error {
	cmd := exec.Command("xdg-open", url)
	return cmd.Run()
}

//go:build darwin

package main

import "os/exec"

func openURL(url string) error {
	cmd := exec.Command("open", url)
	return cmd.Run()
}

//go:build windows

package main

import "os/exec"

func openURL(url string) error {
	cmd := exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	return cmd.Run()
}

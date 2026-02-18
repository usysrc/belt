package uri

import (
	"context"
	"fmt"
	"net/url"
	"os/exec"
	"runtime"
)

func New(vault, file, targetFolder, content string) error {
	uri := buildNewURI(vault, file, targetFolder, content)

	return open(uri)
}

func Open(vault, file, targetFolder string) error {
	uri := buildOpenURI(vault, file, targetFolder)

	return open(uri)
}

func Search(vault, query string) error {
	uri := buildSearchURI(vault, query)

	return open(uri)
}

func buildNewURI(vault, file, targetFolder, content string) string {
	encodedVault := url.PathEscape(vault)
	encodedFile := url.PathEscape(file)
	encodedContent := url.PathEscape(content)
	paramValue := encodedFile

	if targetFolder != "" {
		encodedFolder := url.PathEscape(targetFolder)
		paramValue = fmt.Sprintf("%s/%s", encodedFolder, encodedFile)
	}

	return fmt.Sprintf(
		"obsidian://new?vault=%s&file=%s&content=%s",
		encodedVault,
		paramValue,
		encodedContent,
	)
}

func buildOpenURI(vault, file, targetFolder string) string {
	encodedVault := url.PathEscape(vault)
	encodedFile := url.PathEscape(file)
	paramValue := encodedFile

	if targetFolder != "" {
		encodedFolder := url.PathEscape(targetFolder)
		paramValue = fmt.Sprintf("%s/%s", encodedFolder, encodedFile)
	}

	return fmt.Sprintf(
		"obsidian://open?vault=%s&file=%s&content=",
		encodedVault,
		paramValue,
	)
}

func buildSearchURI(vault, query string) string {
	encodedVault := url.PathEscape(vault)
	encodedQuery := url.PathEscape(query)

	return fmt.Sprintf(
		"obsidian://search?vault=%s&query=%s&content=",
		encodedVault,
		encodedQuery,
	)
}

func open(uri string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.CommandContext(context.Background(), "open", uri)
	case "windows":
		cmd = exec.CommandContext(context.Background(), "cmd", "/c", "start", uri)
	default: // Linux and others
		cmd = exec.CommandContext(context.Background(), "xdg-open", uri)
	}

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to open URI: %w", err)
	}

	return nil
}

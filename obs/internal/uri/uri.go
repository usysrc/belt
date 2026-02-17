package uri

import (
	"context"
	"fmt"
	"net/url"
	"os/exec"
	"runtime"
)

// Request represents the parameters needed to build an Obsidian URI.
type Request struct {
	Vault        string
	Param        string
	Content      string
	Action       string
	TargetFolder string
}

// Execute builds the Obsidian URI based on the provided parameters and opens it.
func Execute(action, vault, param, targetFolder, content string) error {
	req := Request{
		Vault:        vault,
		Param:        param,
		Action:       action,
		TargetFolder: targetFolder,
		Content:      content,
	}
	uri := build(req)

	return open(uri)
}

// build constructs the Obsidian URI based on the request parameters.
func build(params Request) string {
	encodedVault := url.PathEscape(params.Vault)
	encodedParam := url.PathEscape(params.Param)
	encodedContent := url.PathEscape(params.Content)

	var paramName string

	switch params.Action {
	case "search":
		paramName = "query"
	default:
		paramName = "file"

		if params.TargetFolder != "" {
			encodedFolder := url.PathEscape(params.TargetFolder)
			encodedParam = fmt.Sprintf("%s/", encodedFolder) + encodedParam
		}
	}

	uri := fmt.Sprintf("obsidian://%s?vault=%s&%s=%s&content=%s",
		params.Action, encodedVault, paramName, encodedParam, encodedContent)

	return uri
}

// open attempts to open the given URI using the appropriate command based on the operating system.
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

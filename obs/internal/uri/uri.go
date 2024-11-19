package uri

import (
	"fmt"
	"net/url"
	"os/exec"
	"runtime"
)

type URIParams struct {
	Vault        string
	Param        string
	Action       string
	TargetFolder string
}

func Execute(action, vault, param, targetFolder string) error {
	params := URIParams{
		Vault:        vault,
		Param:        param,
		Action:       action,
		TargetFolder: targetFolder,
	}
	uri := buildURI(params)

	return openURI(uri)
}

func buildURI(params URIParams) string {
	encodedVault := url.PathEscape(params.Vault)
	encodedParam := url.PathEscape(params.Param)

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

	uri := fmt.Sprintf("obsidian://%s?vault=%s&%s=%s",
		params.Action, encodedVault, paramName, encodedParam)

	fmt.Println(uri)
	return uri
}

func openURI(uri string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", uri)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", uri)
	default: // Linux and others
		cmd = exec.Command("xdg-open", uri)
	}

	return cmd.Run()
}

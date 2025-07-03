package main

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

// Forge represents a git forge.
type Forge struct {
	Name    string
	TreeURL string // URL to the tree view, with placeholders for repo, branch, and path
}

var forges = map[string]Forge{
	"github.com": {
		Name:    "GitHub",
		TreeURL: "https://{repo}/tree/{branch}/{path}",
	},
	"gitlab.com": {
		Name:    "GitLab",
		TreeURL: "https://{repo}/-/tree/{branch}/{path}",
	},
	"codeberg.org": {
		Name:    "Codeberg",
		TreeURL: "https://{repo}/src/branch/{branch}/{path}",
	},
	"gitea.com": {
		Name:    "Gitea",
		TreeURL: "https://{repo}/src/branch/{branch}/{path}",
	},
}

func main() {
	remoteURL, err := getRemoteURL()
	if err != nil {
		fmt.Println("directory not a repository")
		os.Exit(1)
	}

	toplevelDir, err := getGitToplevel()
	if err != nil {
		fmt.Println("Error getting git toplevel:", err)
		os.Exit(1)
	}

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		os.Exit(1)
	}

	subDir := strings.TrimPrefix(currentDir, toplevelDir)
	subDir = strings.TrimPrefix(subDir, "/")

	forge, repoPath := detectForge(remoteURL)
	var webURL string

	if subDir != "" {
		defaultBranch, err := getDefaultBranch()
		if err != nil {
			fmt.Println("Error getting default branch:", err)
			os.Exit(1)
		}
		treeURL := strings.Replace(forge.TreeURL, "{repo}", repoPath, 1)
		treeURL = strings.Replace(treeURL, "{branch}", defaultBranch, 1)
		webURL = strings.Replace(treeURL, "{path}", subDir, 1)
	} else {
		webURL = "https://" + repoPath
	}

	err = openURL(webURL)
	if err != nil {
		fmt.Println("Error opening URL:", err)
		os.Exit(1)
	}

	fmt.Println("Opening:", webURL)
}

func getRemoteURL() (string, error) {
	cmd := exec.Command("git", "remote", "get-url", "origin")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("directory not a repository")
	}
	return strings.TrimSpace(string(output)), nil
}

func getGitToplevel() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func getDefaultBranch() (string, error) {
	cmd := exec.Command("git", "symbolic-ref", "refs/remotes/origin/HEAD")
	output, err := cmd.Output()
	if err != nil {
		// Fallback for older git versions or different remote names
		cmd = exec.Command("git", "rev-parse", "--abbrev-ref", "origin/HEAD")
		output, err = cmd.Output()
		if err != nil {
			return "", err
		}
	}
	branchName := strings.TrimSpace(string(output))
	branchName = strings.TrimPrefix(branchName, "refs/remotes/origin/")
	return branchName, nil
}

func detectForge(remoteURL string) (Forge, string) {
	var host, path string

	if strings.HasPrefix(remoteURL, "git@") {
		parts := strings.Split(remoteURL, ":")
		hostParts := strings.Split(parts[0], "@")
		host = hostParts[1]
		path = strings.TrimSuffix(parts[1], ".git")
	} else if strings.HasPrefix(remoteURL, "https://") {
		u, err := url.Parse(remoteURL)
		if err == nil {
			host = u.Host
			path = strings.TrimSuffix(u.Path, ".git")
			path = strings.TrimPrefix(path, "/")
		}
	}

	if forge, ok := forges[host]; ok {
		return forge, host + "/" + path
	}

	// Default to Gitea/Codeberg style
	return forges["codeberg.org"], host + "/" + path
}

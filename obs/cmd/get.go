package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"codeberg.org/usysrc/belt/obs/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewGetCmd(cfg *viper.Viper) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Print note content to stdout",
		Long:  "Read a note from your Obsidian vault and print its content to stdout.",
		Args:  cobra.ExactArgs(1),
		Annotations: map[string]string{
			"mode": "stdio",
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			vaultPath, err := config.GetVaultPath(cfg)
			if err != nil {
				return fmt.Errorf("failed to get vault path from config: %w", err)
			}

			noteName := strings.TrimSpace(args[0])
			if noteName == "" {
				return fmt.Errorf("note name cannot be empty")
			}

			candidates, err := resolveNoteCandidates(vaultPath, noteName)
			if err != nil {
				return err
			}

			selected, err := selectCandidate(candidates)
			if err != nil {
				return err
			}

			// #nosec G304 -- selected path is constrained to the vault root by resolver helpers.
			content, err := os.ReadFile(selected)
			if err != nil {
				return fmt.Errorf("failed to read note %q: %w", selected, err)
			}

			if _, err := cmd.OutOrStdout().Write(content); err != nil {
				return fmt.Errorf("failed to write note content: %w", err)
			}

			return nil
		},
	}

	return cmd
}

func resolveNoteCandidates(vaultPath, noteName string) ([]string, error) {
	root := filepath.Clean(vaultPath)
	hasPathSeparator := strings.ContainsRune(noteName, '/') || strings.ContainsRune(noteName, '\\')

	if hasPathSeparator {
		path, err := resolvePathCandidate(root, noteName)
		if err != nil {
			return nil, err
		}

		return []string{path}, nil
	}

	targetName := noteName
	if filepath.Ext(targetName) == "" {
		targetName += ".md"
	}

	matches := make([]string, 0, 1)

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}

		if d.IsDir() {
			return nil
		}

		if filepath.Base(path) == targetName {
			matches = append(matches, path)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to search vault: %w", err)
	}

	if len(matches) == 0 {
		return nil, fmt.Errorf("note %q not found", noteName)
	}

	sort.Strings(matches)

	return matches, nil
}

func resolvePathCandidate(vaultPath, notePath string) (string, error) {
	candidate := filepath.Clean(filepath.Join(vaultPath, notePath))
	if filepath.Ext(candidate) == "" {
		candidate += ".md"
	}

	rootAbs, err := filepath.Abs(vaultPath)
	if err != nil {
		return "", fmt.Errorf("failed to resolve vault path: %w", err)
	}

	candidateAbs, err := filepath.Abs(candidate)
	if err != nil {
		return "", fmt.Errorf("failed to resolve note path: %w", err)
	}

	rel, err := filepath.Rel(rootAbs, candidateAbs)
	if err != nil {
		return "", fmt.Errorf("failed to resolve note path: %w", err)
	}

	if rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
		return "", fmt.Errorf("note path %q escapes vault root", notePath)
	}

	stat, err := os.Stat(candidateAbs)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("note %q not found", notePath)
		}

		return "", fmt.Errorf("failed to read note %q: %w", notePath, err)
	}

	if stat.IsDir() {
		return "", fmt.Errorf("note %q is a directory", notePath)
	}

	return candidateAbs, nil
}

func selectCandidate(candidates []string) (string, error) {
	switch len(candidates) {
	case 0:
		return "", fmt.Errorf("no note candidates found")
	case 1:
		return candidates[0], nil
	}

	return "", fmt.Errorf("multiple note candidates found")
}

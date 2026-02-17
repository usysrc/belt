package cmd

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"codeberg.org/usysrc/belt/obs/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	fzfLookPath = exec.LookPath
	fzfCommand  = exec.CommandContext
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

			selected, err := selectCandidate(
				candidates,
				cmd.ErrOrStderr(),
				cmd.InOrStdin(),
			)
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

func selectCandidate(candidates []string, errOut io.Writer, in io.Reader) (string, error) {
	switch len(candidates) {
	case 0:
		return "", fmt.Errorf("no note candidates found")
	case 1:
		return candidates[0], nil
	}

	if _, err := fzfLookPath("fzf"); err == nil {
		selected, selectErr := selectWithFZF(candidates)
		if selectErr != nil {
			return "", selectErr
		}

		return selected, nil
	}

	return selectWithAsk(candidates, errOut, in)
}

func selectWithFZF(candidates []string) (string, error) {
	cmd := fzfCommand(context.Background(), "fzf")
	cmd.Stdin = strings.NewReader(strings.Join(candidates, "\n") + "\n")

	out, err := cmd.Output()
	if err != nil {
		var exitErr *exec.ExitError

		if errors.As(err, &exitErr) {
			return "", fmt.Errorf("note selection canceled")
		}

		return "", fmt.Errorf("failed to run fzf: %w", err)
	}

	selected := strings.TrimSpace(string(out))
	if selected == "" {
		return "", fmt.Errorf("note selection canceled")
	}

	return selected, nil
}

func selectWithAsk(candidates []string, errOut io.Writer, in io.Reader) (string, error) {
	if _, err := fmt.Fprintln(errOut, "Multiple notes matched:"); err != nil {
		return "", fmt.Errorf("failed to write selection prompt: %w", err)
	}

	for idx, candidate := range candidates {
		if _, err := fmt.Fprintf(errOut, "%d) %s\n", idx+1, candidate); err != nil {
			return "", fmt.Errorf("failed to write selection prompt: %w", err)
		}
	}

	if _, err := fmt.Fprintf(errOut, "Select note [1-%d]: ", len(candidates)); err != nil {
		return "", fmt.Errorf("failed to write selection prompt: %w", err)
	}

	reader := bufio.NewReader(in)

	line, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read selection: %w", err)
	}

	choice, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		return "", fmt.Errorf("invalid selection")
	}

	if choice < 1 || choice > len(candidates) {
		return "", fmt.Errorf("selection out of range")
	}

	return candidates[choice-1], nil
}

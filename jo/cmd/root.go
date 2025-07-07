package cmd

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"
)

// ProcessArgs processes key-value arguments and returns a map
func ProcessArgs(args []string) (map[string]any, error) {
	output := make(map[string]any)

	for _, arg := range args {
		// Split each argument by the first occurrence of '='.
		// `strings.SplitN` is used to ensure only the first '=' acts as a delimiter,
		// allowing values to contain '=' characters (e.g., "url=http://example.com?a=b").
		parts := strings.SplitN(arg, "=", 2)

		// Check if the argument is in a valid "key=value" format.
		// If not, return an error for this specific argument.
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid argument format '%s'. Expected 'key=value' or 'key[subkey]=value'", arg)
		}

		keyPart := parts[0] // The part before the first '='
		value := parts[1]   // The part after the first '='

		// Parse nested keys (e.g., "user[name]" or "users[123][name]")
		if strings.Contains(keyPart, "[") && strings.Contains(keyPart, "]") {
			err := setNestedValue(output, keyPart, value)
			if err != nil {
				return nil, err
			}
		} else {
			// If no valid nested key format is found, treat it as a simple key-value pair.
			output[keyPart] = value
		}
	}

	return output, nil
}

// setNestedValue sets a value at a nested key path (e.g., "user[name]" or "users[123][name]")
func setNestedValue(output map[string]any, keyPath string, value string) error {
	// Parse the key path to extract all keys
	keys, err := parseKeyPath(keyPath)
	if err != nil {
		return err
	}

	// Navigate through the nested structure, creating maps as needed
	current := output
	for i, key := range keys {
		if i == len(keys)-1 {
			// Last key, set the value
			current[key] = value
		} else {
			// Not the last key, ensure we have a map to navigate into
			if existing, exists := current[key]; exists {
				if nestedMap, ok := existing.(map[string]any); ok {
					current = nestedMap
				} else {
					// Override non-map value with a new map
					newMap := make(map[string]any)
					current[key] = newMap
					current = newMap
				}
			} else {
				// Create new map
				newMap := make(map[string]any)
				current[key] = newMap
				current = newMap
			}
		}
	}

	return nil
}

// parseKeyPath parses a key path like "user[name]" or "users[123][name]" into individual keys
func parseKeyPath(keyPath string) ([]string, error) {
	var keys []string
	current := ""
	inBracket := false

	for _, char := range keyPath {
		switch char {
		case '[':
			if inBracket {
				return nil, fmt.Errorf("invalid key path '%s': nested brackets not properly closed", keyPath)
			}
			if current != "" {
				keys = append(keys, current)
				current = ""
			}
			inBracket = true
		case ']':
			if !inBracket {
				return nil, fmt.Errorf("invalid key path '%s': closing bracket without opening bracket", keyPath)
			}
			if current == "" {
				return nil, fmt.Errorf("invalid key path '%s': empty key in brackets", keyPath)
			}
			keys = append(keys, current)
			current = ""
			inBracket = false
		default:
			current += string(char)
		}
	}

	if inBracket {
		return nil, fmt.Errorf("invalid key path '%s': unclosed bracket", keyPath)
	}

	if current != "" {
		keys = append(keys, current)
	}

	if len(keys) == 0 {
		return nil, fmt.Errorf("invalid key path '%s': no keys found", keyPath)
	}

	return keys, nil
}

// ReadStdinArgs reads key-value pairs from stdin
func ReadStdinArgs(reader io.Reader) ([]string, error) {
	args := make([]string, 0)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			args = append(args, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading from stdin: %w", err)
	}

	return args, nil
}

// ConvertToJSON converts a map to pretty-printed JSON
func ConvertToJSON(data map[string]any) (string, error) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error marshalling JSON: %w", err)
	}
	return string(jsonData), nil
}

var rootCmd = &cobra.Command{
	Use:   "jo [key=value...]",
	Short: "A command-line tool that converts key-value arguments to JSON output",
	Long: `jo is a simple command-line tool that converts key-value arguments to JSON output.
It supports both command-line arguments and stdin input, as well as nested objects
using bracket notation.`,
	Example: `  # Simple key-value pairs
  jo name=John age=30 city=Boston

  # Nested objects using bracket notation
  jo user[name]=John user[age]=30 config[debug]=true

  # Reading from stdin
  echo -e "name=John\nage=30" | jo

  # Combined stdin and command-line arguments
  echo "database[host]=localhost" | jo database[port]=5432 debug=true`,
	Args: cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		var allArgs []string

		// Check if we have stdin input
		stat, err := os.Stdin.Stat()
		if err != nil {
			return fmt.Errorf("error checking stdin: %w", err)
		}

		hasStdin := (stat.Mode() & os.ModeCharDevice) == 0
		if hasStdin {
			// Read from stdin
			stdinArgs, err := ReadStdinArgs(os.Stdin)
			if err != nil {
				return err
			}
			allArgs = append(allArgs, stdinArgs...)
		}

		// Add command-line arguments
		allArgs = append(allArgs, args...) // Exclude the command name

		// If no arguments provided from either source, show usage
		if len(allArgs) == 0 {
			return cmd.Help()
		}

		// Process arguments with error handling for invalid format
		var validArgs []string
		for _, arg := range allArgs {
			// Quick validation check
			if !strings.Contains(arg, "=") {
				fmt.Fprintf(os.Stderr, "Warning: Skipping invalid argument format '%s'. Expected 'key=value' or 'key[subkey]=value'.\n", arg)
				continue
			}
			validArgs = append(validArgs, arg)
		}

		// Process all valid arguments
		output, err := ProcessArgs(validArgs)
		if err != nil {
			return err
		}

		// Convert to JSON
		jsonStr, err := ConvertToJSON(output)
		if err != nil {
			return err
		}

		// Print the resulting JSON string to standard output
		fmt.Println(jsonStr)
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	//rootCmd.ExecuteContext()
	return fang.Execute(context.Background(), rootCmd)
}

package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

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

		// Check for the presence of square brackets to identify nested keys
		// (e.g., "some_key[sub_key]").
		openBracketIndex := strings.Index(keyPart, "[")
		closeBracketIndex := strings.Index(keyPart, "]")

		// If both an opening and closing bracket are found, and the opening bracket
		// comes before the closing one, treat it as a nested key.
		if openBracketIndex != -1 && closeBracketIndex != -1 && openBracketIndex < closeBracketIndex {
			// Extract the main key (e.g., "some_key") and the subkey (e.g., "sub_key").
			mainKey := keyPart[:openBracketIndex]
			subKey := keyPart[openBracketIndex+1 : closeBracketIndex]

			// Ensure that the value associated with the `mainKey` in our `output` map
			// is itself a map. If it doesn't exist, or if it's currently a non-map type
			// (e.g., if "mainKey=simple_value" was processed earlier),
			// initialize or re-initialize it as a new map.
			// This handles cases where a simple key might be later turned into a parent
			// for a nested key, or vice-versa, adhering to JSON's last-assignment-wins rule.
			nestedMap, ok := output[mainKey].(map[string]any)
			if !ok {
				nestedMap = make(map[string]any)
				output[mainKey] = nestedMap
			}

			// Add the subKey and its value to the nested map.
			nestedMap[subKey] = value
		} else {
			// If no valid nested key format is found, treat it as a simple key-value pair.
			output[keyPart] = value
		}
	}

	return output, nil
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
		allArgs = append(allArgs, args...)

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
	return rootCmd.Execute()
}
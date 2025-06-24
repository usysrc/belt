#!/bin/bash

# Define the base directory where your tools are located within the repo.
# This assumes the script is run from the 'belt' directory, or the 'belt' directory
# is the parent of the directories containing your Go tools.
# TOOLS_BASE_DIR="." means tools are direct subdirectories of 'belt'.
TOOLS_BASE_DIR="."

echo "Starting Go CLI tools installation..."
echo "" # Add a blank line for better separation

# Iterate through each subdirectory in the TOOLS_BASE_DIR
for tool_dir in "$TOOLS_BASE_DIR"/*/; do
    if [ -d "$tool_dir" ]; then # Check if it's actually a directory
        tool_name=$(basename "$tool_dir")

        echo -n "  $tool_name... " # -n prevents newline

        # Navigate into the tool's directory and run go install.
        # This will install the binary to the directory specified by GOBIN.
        # @latest ensures it builds the latest version of the module.
        (cd "$tool_dir" && go install "./...") > /dev/null 2>&1 # Redirect stdout/stderr to hide verbose output

        if [ $? -eq 0 ]; then
            echo "✓" # Print checkmark
        else
            echo "✗" # Print cross mark
            echo "    Failed to install $tool_name. Please check for errors above if you re-run without /dev/null."
        fi
    fi
done

echo "" # Add a blank line for better separation
echo "Installation process completed."

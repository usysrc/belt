#!/bin/bash

# List of tool repositories to install
tools=(
  github.com/usysrc/uuid
  github.com/usysrc/ssl-expiry
  github.com/usysrc/timezone
  github.com/usysrc/xls-format
  github.com/usysrc/serve
  github.com/usysrc/jenv
  github.com/usysrc/hex
  # Add more tool repositories here
)

# Check if Go is installed
if ! command -v go &>/dev/null; then
  echo "Go is not installed. Please install Go to proceed."
  exit 1
fi

# Function to install tools
install_tool() {
  local tool=$1
  echo "Installing $tool..."
  if go install "$tool@latest"; then
    echo "$tool installed successfully."
  else
    echo "Failed to install $tool." >&2
  fi
}

# Install each tool
for tool in "${tools[@]}"; do
  install_tool "$tool"
done

echo "All installations completed."


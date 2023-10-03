#!/bin/bash

# List of tool repositories to install
tools=(
   github.com/usysrc/uuid
   github.com/usysrc/ssl-expiry
   github.com/usysrc/timezone
   github.com/usysrc/xls-format
   # Add more tool repositories as needed
)

# Install each tool
for tool in "${tools[@]}"; do
   go install "$tool@latest"
done

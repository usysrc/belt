package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// Decode JSON from stdin
	var data map[string]interface{}
	err := json.NewDecoder(os.Stdin).Decode(&data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error decoding JSON: %v\n", err)
		os.Exit(1)
	}

	// Export each key-value pair as an environment variable
	for key, value := range data {
		strValue := fmt.Sprintf("%v", value)
		fmt.Printf("export %s=%s\n", key, strValue)
	}
}

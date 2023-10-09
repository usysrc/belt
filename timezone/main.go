package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var timeFormat string

func getCurrentTime(city string) (string, error) {
	// LoadLocation loads the timezone data for the given city.
	loc, err := time.LoadLocation(city)
	if err != nil {
		return "", err
	}
	// Get the current time in the specified city.
	currentTime := time.Now().In(loc)
	// Format the time according to the user-specified format.
	return currentTime.Format(timeFormat), nil
}

func findCurrentTimeByRegion(city string, regions []string) (string, error) {
	for _, region := range regions {
		// Fetch the current time for the specified city.
		currentTime, err := getCurrentTime(region + "/" + city)
		if err == nil {
			return currentTime, nil
		}
	}
	return "", fmt.Errorf("no valid timezone found for %s", city)
}

var rootCmd = &cobra.Command{
	Use:   "timezone [city]",
	Short: "Get current time in a city",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Combine the input arguments (city name) into a single string with spaces if needed.
		city := strings.Join(args, " ")

		// Define the list of regions to search for the city's timezone.
		regions := []string{"Asia", "America", "Europe"}

		// Find the current time for the specified city within the available regions.
		currentTime, err := findCurrentTimeByRegion(city, regions)
		if err != nil {
			// If an error occurs (e.g., invalid city or no valid timezone found), display an error message and exit with status code 1.
			fmt.Printf("Error getting time for %s: %v\n", city, err)
			os.Exit(1)
		}

		// Print the current time in the specified city.
		fmt.Printf("Current time in %s: %s\n", city, currentTime)
	},
}

func main() {
	// Add a flag to the root command to allow users to specify the time format.
	rootCmd.Flags().StringVarP(&timeFormat, "format", "f", "15:04", "Time format (e.g., 15:04 or 03:04 PM)")

	// Execute the root command, which will run the CLI app.
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

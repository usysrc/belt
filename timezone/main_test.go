package main

import (
	"testing"
)

func TestGetCurrentTime(t *testing.T) {

	timeFormat = "15:04"

	tests := []struct {
		city        string
		expectError bool
	}{
		{"America/New_York", false},
		{"Europe/London", false},
		{"Invalid_City", true},
	}

	for _, tt := range tests {
		t.Run(tt.city, func(t *testing.T) {
			currentTime, err := getCurrentTime(tt.city)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error for city %s, but got nil", tt.city)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error for city %s: %v", tt.city, err)
				}

				// Check if currentTime is in HH:MM format
				if len(currentTime) != 5 {
					t.Errorf("Invalid time format for city %s: %s", tt.city, currentTime)
				}
			}
		})
	}
}

package cmd

import (
	"reflect"
	"strings"
	"testing"
)

func TestProcessArgs(t *testing.T) {
	t.Parallel()
	//nolint:govet
	tests := []struct {
		name     string
		args     []string
		expected map[string]any
		wantErr  bool
	}{
		{
			name:     "simple key-value",
			args:     []string{"name=value"},
			expected: map[string]any{"name": "value"},
			wantErr:  false,
		},
		{
			name:     "multiple simple key-values",
			args:     []string{"name=value", "age=30"},
			expected: map[string]any{"name": "value", "age": "30"},
			wantErr:  false,
		},
		{
			name:     "nested key",
			args:     []string{"user[name]=John"},
			expected: map[string]any{"user": map[string]any{"name": "John"}},
			wantErr:  false,
		},
		{
			name:     "multiple nested keys",
			args:     []string{"user[name]=John", "user[age]=30"},
			expected: map[string]any{"user": map[string]any{"name": "John", "age": "30"}},
			wantErr:  false,
		},
		{
			name:     "value with equals sign",
			args:     []string{"url=http://example.com?param=value"},
			expected: map[string]any{"url": "http://example.com?param=value"},
			wantErr:  false,
		},
		{
			name: "mixed simple and nested",
			args: []string{"name=John", "address[city]=New York", "address[zip]=10001"},
			expected: map[string]any{
				"name": "John",
				"address": map[string]any{
					"city": "New York",
					"zip":  "10001",
				},
			},
			wantErr: false,
		},
		{
			name:     "invalid format returns error",
			args:     []string{"invalid"},
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result, err := ProcessArgs(tt.args)

			if tt.wantErr {
				if err == nil {
					t.Errorf("ProcessArgs() expected error but got none")
				}

				return
			}

			if err != nil {
				t.Errorf("ProcessArgs() unexpected error: %v", err)

				return
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ProcessArgs() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestNestedKeyOverrides(t *testing.T) {
	t.Parallel()

	//nolint:govet
	tests := []struct {
		name     string
		args     []string
		expected map[string]any
	}{
		{
			name: "simple key becomes nested map",
			args: []string{"user=simple", "user[name]=John"},
			expected: map[string]any{
				"user": map[string]any{
					"name": "John",
				},
			},
		},
		{
			name: "nested map becomes simple key",
			args: []string{"user[name]=John", "user=simple"},
			expected: map[string]any{
				"user": "simple",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result, err := ProcessArgs(tt.args)
			if err != nil {
				t.Errorf("ProcessArgs() unexpected error: %v", err)

				return
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ProcessArgs() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestReadStdinArgs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected []string
		wantErr  bool
	}{
		{
			name:     "simple input",
			input:    "name=John\nage=30\n",
			expected: []string{"name=John", "age=30"},
			wantErr:  false,
		},
		{
			name:     "nested keys",
			input:    "user[name]=John\nuser[age]=30\n",
			expected: []string{"user[name]=John", "user[age]=30"},
			wantErr:  false,
		},
		{
			name:     "empty lines ignored",
			input:    "name=John\n\nage=30\n\n",
			expected: []string{"name=John", "age=30"},
			wantErr:  false,
		},
		{
			name:     "empty input",
			input:    "",
			expected: []string{},
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			reader := strings.NewReader(tt.input)
			result, err := ReadStdinArgs(reader)

			if tt.wantErr {
				if err == nil {
					t.Errorf("ReadStdinArgs() expected error but got none")
				}

				return
			}

			if err != nil {
				t.Errorf("ReadStdinArgs() unexpected error: %v", err)

				return
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ReadStdinArgs() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestConvertToJSON(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    map[string]any
		expected string
		wantErr  bool
	}{
		{
			name:     "simple map",
			input:    map[string]any{"name": "John", "age": "30"},
			expected: "{\n  \"age\": \"30\",\n  \"name\": \"John\"\n}",
			wantErr:  false,
		},
		{
			name: "nested map",
			input: map[string]any{
				"user": map[string]any{
					"name": "John",
					"age":  "30",
				},
			},
			expected: "{\n  \"user\": {\n    \"age\": \"30\",\n    \"name\": \"John\"\n  }\n}",
			wantErr:  false,
		},
		{
			name:     "empty map",
			input:    map[string]any{},
			expected: "{}",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result, err := ConvertToJSON(tt.input)

			if tt.wantErr {
				if err == nil {
					t.Errorf("ConvertToJSON() expected error but got none")
				}

				return
			}

			if err != nil {
				t.Errorf("ConvertToJSON() unexpected error: %v", err)

				return
			}

			if result != tt.expected {
				t.Errorf("ConvertToJSON() = %q, want %q", result, tt.expected)
			}
		})
	}
}

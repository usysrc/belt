package uri

import "testing"

func TestBuildSearchURI(t *testing.T) {
	t.Parallel()

	got := buildSearchURI("my Vault", "my Query")
	want := "obsidian://search?vault=my%20Vault&query=my%20Query&content="

	if got != want {
		t.Errorf("buildSearchURI() = %v, want %v", got, want)
	}
}

func TestBuildOpenURI(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		vault        string
		file         string
		targetFolder string
		expected     string
	}{
		{
			name:     "open without folder",
			vault:    "myVault",
			file:     "myParam",
			expected: "obsidian://open?vault=myVault&file=myParam&content=",
		},
		{
			name:         "open with folder",
			vault:        "myVault",
			file:         "myParam",
			targetFolder: "myFolder",
			expected:     "obsidian://open?vault=myVault&file=myFolder/myParam&content=",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := buildOpenURI(tt.vault, tt.file, tt.targetFolder)
			if got != tt.expected {
				t.Errorf("buildOpenURI() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestBuildNewURI(t *testing.T) {
	t.Parallel()

	got := buildNewURI("my Vault", "my Param", "my Folder", "hello world")
	want := "obsidian://new?vault=my%20Vault&file=my%20Folder/my%20Param&content=hello%20world"

	if got != want {
		t.Errorf("buildNewURI() = %v, want %v", got, want)
	}
}

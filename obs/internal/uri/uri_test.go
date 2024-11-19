package uri

import (
	"testing"
)

func TestBuildURI(t *testing.T) {
	tests := []struct {
		vault        string
		param        string
		action       string
		targetFolder string
		expected     string
	}{
		{
			vault:    "myVault",
			param:    "myParam",
			action:   "search",
			expected: "obsidian://search?vault=myVault&query=myParam",
		},
		{
			vault:    "myVault",
			param:    "myParam",
			action:   "open",
			expected: "obsidian://open?vault=myVault&file=myParam",
		},
		{
			vault:        "myVault",
			param:        "myParam",
			action:       "open",
			targetFolder: "myFolder",
			expected:     "obsidian://open?vault=myVault&file=myFolder/myParam",
		},
		{
			vault:        "my Vault",
			param:        "my Param",
			action:       "open",
			targetFolder: "my Folder",
			expected:     "obsidian://open?vault=my%20Vault&file=my%20Folder/my%20Param",
		},
	}

	for _, tt := range tests {
		t.Run(tt.action, func(t *testing.T) {
			params := URIParams{
				Vault:        tt.vault,
				Param:        tt.param,
				Action:       tt.action,
				TargetFolder: tt.targetFolder,
			}
			got := buildURI(params)
			if got != tt.expected {
				t.Errorf("buildURI() = %v, want %v", got, tt.expected)
			}
		})
	}
}

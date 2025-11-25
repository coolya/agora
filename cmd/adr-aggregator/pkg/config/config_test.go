package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Create a config.yaml file in the current directory for the test
	content := `
sources:
  - type: github
    url: https://github.com/test-org/test-repo
    path: docs/adrs
    auth:
      token: test-token
`
	configFileName := "config.yaml"
	if err := os.WriteFile(configFileName, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to write temp config file: %v", err)
	}
	defer os.Remove(configFileName) // clean up

	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig() failed: %v", err)
	}

	if len(cfg.Sources) != 1 {
		t.Fatalf("expected 1 source, got %d", len(cfg.Sources))
	}

	source := cfg.Sources[0]
	if source.Type != "github" {
		t.Errorf("expected source type 'github', got %q", source.Type)
	}
	if source.URL != "https://github.com/test-org/test-repo" {
		t.Errorf("expected source URL 'https://github.com/test-org/test-repo', got %q", source.URL)
	}
}

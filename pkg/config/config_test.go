package config_test

import (
	"helm-airgap-bundler/pkg/config"
	"os"
	"testing"
)

// createTempConfigFile creates a temporary YAML config file for testing purposes.
func createTempConfigFile(t *testing.T, content string) string {
	t.Helper()
	tmpFile, err := os.CreateTemp("", "config-*.yaml")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer tmpFile.Close()

	_, err = tmpFile.WriteString(content)
	if err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}

	return tmpFile.Name()
}

func TestLoadConfig(t *testing.T) {
	configContent := `
charts:
  - url: "https://example.com/charts/chart1"
    username: "user1"
    password: "pass1"
  - url: "https://example.com/charts/chart2"
`
	configPath := createTempConfigFile(t, configContent)
	defer os.Remove(configPath) // Clean up the temp file after the test

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		t.Fatalf("LoadConfig returned an error: %v", err)
	}

	if len(cfg.Charts) != 2 {
		t.Errorf("Expected 2 charts, got %d", len(cfg.Charts))
	}

	expectedURL := "https://example.com/charts/chart1"
	if cfg.Charts[0].URL != expectedURL {
		t.Errorf("Expected URL %s, got %s", expectedURL, cfg.Charts[0].URL)
	}
}

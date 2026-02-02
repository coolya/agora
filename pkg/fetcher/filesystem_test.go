package fetcher

import (
	"agora/pkg/config"
	"os"
	"path/filepath"
	"testing"
)

func TestFileSystemFetcher_Fetch(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "adr-test")
	if err != nil {
		t.Fatalf("failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create test markdown files
	testFiles := map[string]string{
		"adr-001-test.md": "# ADR 001: Test Decision\n\nThis is a test ADR.",
		"adr-002-another.md": "# ADR 002: Another Decision\n\nThis is another test ADR.",
	}

	for filename, content := range testFiles {
		filePath := filepath.Join(tmpDir, filename)
		err := os.WriteFile(filePath, []byte(content), 0644)
		if err != nil {
			t.Fatalf("failed to write test file %s: %v", filename, err)
		}
	}

	// Create a non-markdown file (should be ignored)
	err = os.WriteFile(filepath.Join(tmpDir, "readme.txt"), []byte("This is a readme"), 0644)
	if err != nil {
		t.Fatalf("failed to write readme.txt: %v", err)
	}

	// Create a subdirectory (should be ignored)
	subDir := filepath.Join(tmpDir, "subdir")
	err = os.Mkdir(subDir, 0755)
	if err != nil {
		t.Fatalf("failed to create subdirectory: %v", err)
	}

	fetcher := NewFileSystemFetcher()
	source := config.Source{
		Name:   "Test Local",
		Type:   "filesystem",
		URL:    "", // Not used for filesystem
		Path:   tmpDir,
	}

	adrs, err := fetcher.Fetch(source)
	if err != nil {
		t.Fatalf("Fetch failed: %v", err)
	}

	if len(adrs) != 2 {
		t.Errorf("expected 2 ADRs, got %d", len(adrs))
	}

	// Verify the ADRs contain expected data
	for _, adr := range adrs {
		if adr.SourceType != "filesystem" {
			t.Errorf("expected source type 'filesystem', got %s", adr.SourceType)
		}
		if adr.SourceName != "Test Local" {
			t.Errorf("expected source name 'Test Local', got %s", adr.SourceName)
		}
		if adr.ID == "" {
			t.Error("expected non-empty ID")
		}
		if adr.Content == "" {
			t.Error("expected non-empty content")
		}
	}
}

func TestFileSystemFetcher_InvalidPath(t *testing.T) {
	fetcher := NewFileSystemFetcher()
	source := config.Source{
		Name:   "Invalid",
		Type:   "filesystem",
		Path:   "/nonexistent/path",
	}

	_, err := fetcher.Fetch(source)
	if err == nil {
		t.Error("expected error for non-existent path")
	}
}

func TestFileSystemFetcher_NotADirectory(t *testing.T) {
	// Create a temporary file
	tmpFile, err := os.CreateTemp("", "test-file-*.txt")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	fetcher := NewFileSystemFetcher()
	source := config.Source{
		Name:   "NotDir",
		Type:   "filesystem",
		Path:   tmpFile.Name(),
	}

	_, err = fetcher.Fetch(source)
	if err == nil {
		t.Error("expected error when path is not a directory")
	}
}

func TestFileSystemFetcher_PathNormalization(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "adr-test-path-norm")
	if err != nil {
		t.Fatalf("failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create a test markdown file
	testFile := "adr-001.md"
	testPath := filepath.Join(tmpDir, testFile)
	err = os.WriteFile(testPath, []byte("# ADR 001\n\nTest ADR"), 0644)
	if err != nil {
		t.Fatalf("failed to write test file: %v", err)
	}

	// Save current working directory
	originalWD, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get working directory: %v", err)
	}
	defer os.Chdir(originalWD)

	// Change to temp directory to test relative paths
	err = os.Chdir(tmpDir)
	if err != nil {
		t.Fatalf("failed to change working directory: %v", err)
	}

	fetcher := NewFileSystemFetcher()

	// Fetch using relative path
	relativeSource := config.Source{
		Name:   "Relative",
		Type:   "filesystem",
		Path:   ".",
	}
	relativeADRs, err := fetcher.Fetch(relativeSource)
	if err != nil {
		t.Fatalf("Fetch with relative path failed: %v", err)
	}

	if len(relativeADRs) != 1 {
		t.Fatalf("expected 1 ADR from relative path, got %d", len(relativeADRs))
	}

	// Fetch using absolute path
	absoluteSource := config.Source{
		Name:   "Absolute",
		Type:   "filesystem",
		Path:   tmpDir,
	}
	absoluteADRs, err := fetcher.Fetch(absoluteSource)
	if err != nil {
		t.Fatalf("Fetch with absolute path failed: %v", err)
	}

	if len(absoluteADRs) != 1 {
		t.Fatalf("expected 1 ADR from absolute path, got %d", len(absoluteADRs))
	}

	// Verify that both relative and absolute paths produce the same URL
	// This ensures no collision issues if the same path is added with different formats
	if relativeADRs[0].URL != absoluteADRs[0].URL {
		t.Errorf("URL mismatch: relative path produced %q, absolute path produced %q",
			relativeADRs[0].URL, absoluteADRs[0].URL)
	}

	// Verify that the IDs are identical (which ensures uniqueness across different path formats)
	if relativeADRs[0].ID != absoluteADRs[0].ID {
		t.Errorf("ID mismatch: relative path produced %q, absolute path produced %q",
			relativeADRs[0].ID, absoluteADRs[0].ID)
	}
}

func TestFileSystemFetcher_EmptyDirectory(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "adr-test-empty")
	if err != nil {
		t.Fatalf("failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create some non-markdown files in the directory
	err = os.WriteFile(filepath.Join(tmpDir, "readme.txt"), []byte("This is a readme"), 0644)
	if err != nil {
		t.Fatalf("failed to write readme.txt: %v", err)
	}

	err = os.WriteFile(filepath.Join(tmpDir, "config.yaml"), []byte("key: value"), 0644)
	if err != nil {
		t.Fatalf("failed to write config.yaml: %v", err)
	}

	fetcher := NewFileSystemFetcher()
	source := config.Source{
		Name: "Empty Test",
		Type: "filesystem",
		Path: tmpDir,
	}

	adrs, err := fetcher.Fetch(source)
	if err != nil {
		t.Fatalf("Fetch failed: %v", err)
	}

	if len(adrs) != 0 {
		t.Errorf("expected 0 ADRs from directory with no markdown files, got %d", len(adrs))
	}

	if adrs == nil {
		t.Error("expected non-nil slice, got nil")
	}
}


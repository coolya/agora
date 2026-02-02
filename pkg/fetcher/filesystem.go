package fetcher

import (
	"agora/pkg/config"
	"agora/pkg/domain"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// FileSystemFetcher fetches ADRs from the local file system.
type FileSystemFetcher struct{}

// NewFileSystemFetcher creates a new FileSystemFetcher.
func NewFileSystemFetcher() *FileSystemFetcher {
	return &FileSystemFetcher{}
}

// Fetch retrieves ADRs from the configured local directory.
func (f *FileSystemFetcher) Fetch(source config.Source) ([]domain.ADR, error) {
	// Resolve to absolute path and evaluate symlinks to avoid collisions from path variations
	absPath, err := filepath.Abs(source.Path)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve absolute path for %s: %w", source.Path, err)
	}

	// Evaluate symlinks to ensure consistent paths across different reference methods
	absPath, err = filepath.EvalSymlinks(absPath)
	if err != nil {
		return nil, fmt.Errorf("failed to evaluate symlinks for %s: %w", absPath, err)
	}

	// Check if the path exists and is a directory
	fileInfo, err := os.Stat(absPath)
	if err != nil {
		return nil, fmt.Errorf("failed to access path %s: %w", absPath, err)
	}

	if !fileInfo.IsDir() {
		return nil, fmt.Errorf("path %s is not a directory", absPath)
	}

	// Read all files in the directory
	entries, err := os.ReadDir(absPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory %s: %w", absPath, err)
	}

	adrs := make([]domain.ADR, 0)

	for _, entry := range entries {
		// Only process files, skip directories
		if entry.IsDir() {
			continue
		}

		// Only process markdown files
		if !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}

		filePath := filepath.Join(absPath, entry.Name())

		// Read the file content
		content, err := os.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
		}

		// Create ADR record with normalized absolute path in URL to ensure uniqueness
		normalizedPath := filepath.ToSlash(filePath)
		fileURL := "file:///" + normalizedPath
		if strings.HasPrefix(normalizedPath, "/") {
			// Unix-style absolute path already starts with '/', so "file://" yields "file:///..."
			fileURL = "file://" + normalizedPath
		}
		adrs = append(adrs, domain.ADR{
			ID:         domain.GenerateID(fileURL),
			Title:      entry.Name(),
			Content:    string(content),
			URL:        fileURL,
			SourceType: source.Type,
			SourceURL:  source.URL,
			SourceName: source.Name,
		})
	}

	return adrs, nil
}

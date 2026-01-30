package storage

import (
	"agora/pkg/domain"
	"path/filepath"
	"testing"
)

func TestInitDB(t *testing.T) {
	// Create a temporary directory for the test
	tmpDir := t.TempDir()
	dbPath := filepath.Join(tmpDir, "test_agora.db")

	err := InitDB(dbPath)
	if err != nil {
		t.Fatalf("InitDB() failed: %v", err)
	}

	if DB == nil {
		t.Fatal("DB is nil after InitDB")
	}
}

func TestSaveADRs(t *testing.T) {
	// Create a temporary directory for the test
	tmpDir := t.TempDir()
	dbPath := filepath.Join(tmpDir, "test_agora_save.db")

	err := InitDB(dbPath)
	if err != nil {
		t.Fatalf("InitDB() failed: %v", err)
	}

	// Test saving ADRs
	adrs := []domain.ADR{
		{
			ID:         domain.GenerateID("https://example.com/adr-1"),
			Title:      "ADR 1: Use Go",
			Status:     "ACCEPTED",
			Content:    "We will use Go for this project.",
			URL:        "https://example.com/adr-1",
			SourceType: "github",
			SourceURL:  "https://github.com/example/repo",
			SourceName: "github-0",
		},
		{
			ID:         domain.GenerateID("https://example.com/adr-2"),
			Title:      "ADR 2: Use SQLite",
			Status:     "PROPOSED",
			Content:    "We will use SQLite for persistence.",
			URL:        "https://example.com/adr-2",
			SourceType: "github",
			SourceURL:  "https://github.com/example/repo",
			SourceName: "github-0",
		},
	}

	err = SaveADRs(adrs)
	if err != nil {
		t.Fatalf("SaveADRs() failed: %v", err)
	}

	// Verify ADRs were saved
	var count int64
	DB.Model(&domain.ADR{}).Count(&count)
	if count != 2 {
		t.Errorf("expected 2 ADRs in database, got %d", count)
	}
}

func TestSaveADRsUpsert(t *testing.T) {
	// Create a temporary directory for the test
	tmpDir := t.TempDir()
	dbPath := filepath.Join(tmpDir, "test_agora_upsert.db")

	err := InitDB(dbPath)
	if err != nil {
		t.Fatalf("InitDB() failed: %v", err)
	}

	// Save initial ADR
	adrs := []domain.ADR{
		{
			ID:         domain.GenerateID("https://example.com/adr-1"),
			Title:      "ADR 1: Use Go",
			Status:     "PROPOSED",
			Content:    "We will use Go for this project.",
			URL:        "https://example.com/adr-1",
			SourceType: "github",
			SourceURL:  "https://github.com/example/repo",
			SourceName: "github-0",
		},
	}

	err = SaveADRs(adrs)
	if err != nil {
		t.Fatalf("SaveADRs() failed on first save: %v", err)
	}

	// Update the same ADR (same URL, different status)
	updatedADRs := []domain.ADR{
		{
			ID:         domain.GenerateID("https://example.com/adr-1"),
			Title:      "ADR 1: Use Go (Updated)",
			Status:     "ACCEPTED",
			Content:    "We will use Go for this project. Updated content.",
			URL:        "https://example.com/adr-1",
			SourceType: "github",
			SourceURL:  "https://github.com/example/repo",
			SourceName: "github-0",
		},
	}

	err = SaveADRs(updatedADRs)
	if err != nil {
		t.Fatalf("SaveADRs() failed on upsert: %v", err)
	}

	// Verify only one ADR exists
	var count int64
	DB.Model(&domain.ADR{}).Count(&count)
	if count != 1 {
		t.Errorf("expected 1 ADR in database after upsert, got %d", count)
	}

	// Verify the ADR was updated
	var adr domain.ADR
	result := DB.Where("url = ?", "https://example.com/adr-1").First(&adr)
	if result.Error != nil {
		t.Fatalf("Failed to query ADR: %v", result.Error)
	}
	if adr.Status != "ACCEPTED" {
		t.Errorf("expected status 'ACCEPTED', got %q", adr.Status)
	}
	if adr.Title != "ADR 1: Use Go (Updated)" {
		t.Errorf("expected title 'ADR 1: Use Go (Updated)', got %q", adr.Title)
	}
}

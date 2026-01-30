package main

import (
	"agora/pkg/fetcher"
	"agora/pkg/parser"
	"agora/pkg/storage"
	"path/filepath"
	"testing"
)

func TestIntegrationWithDatabase(t *testing.T) {
	// Create a temporary directory for the test
	tmpDir := t.TempDir()
	dbPath := filepath.Join(tmpDir, "test_integration.db")

	// Initialize the database
	err := storage.InitDB(dbPath)
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}

	// Create some mock ADRs
	adrs := []fetcher.ADR{
		{
			Title:   "ADR 1: Use Go",
			Content: "# ADR 1: Use Go\nStatus: ACCEPTED\nWe will use Go.",
			URL:     "https://example.com/adr-1",
		},
		{
			Title:   "ADR 2: Use SQLite",
			Content: "# ADR 2: Use SQLite\nStatus: PROPOSED\nWe will use SQLite.",
			URL:     "https://example.com/adr-2",
		},
	}

	// Parse the ADRs
	for i := range adrs {
		parser.Parse(&adrs[i])
	}

	// Verify status was parsed
	if adrs[0].Status != "ACCEPTED" {
		t.Errorf("expected status 'ACCEPTED', got %q", adrs[0].Status)
	}
	if adrs[1].Status != "PROPOSED" {
		t.Errorf("expected status 'PROPOSED', got %q", adrs[1].Status)
	}

	// Save to database
	err = storage.SaveADRs(adrs)
	if err != nil {
		t.Fatalf("Failed to save ADRs: %v", err)
	}

	// Verify count
	var count int64
	storage.DB.Model(&fetcher.ADR{}).Count(&count)
	if count != 2 {
		t.Errorf("expected 2 ADRs in database, got %d", count)
	}

	// Verify data
	var savedADRs []fetcher.ADR
	storage.DB.Find(&savedADRs)

	// Check that we have the right number of records
	if len(savedADRs) != 2 {
		t.Fatalf("expected 2 saved ADRs, got %d", len(savedADRs))
	}
}

package storage

import (
	"agora/pkg/domain"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// DB holds the database connection
var DB *gorm.DB

// InitDB initializes the SQLite database connection and migrates the schema
func InitDB(dbPath string) error {
	var err error
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// Auto-migrate the schema
	err = DB.AutoMigrate(&domain.ADR{})
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	return nil
}

// SaveADRs saves ADRs to the database using upsert (insert or update on conflict)
func SaveADRs(adrs []domain.ADR) error {
	if DB == nil {
		return fmt.Errorf("database not initialized")
	}

	// Use GORM's Clauses to handle upsert based on URL (unique index)
	// This will update all fields if a conflict occurs on the URL
	result := DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "url"}},
		DoUpdates: clause.AssignmentColumns([]string{"id", "title", "status", "content", "source_type", "source_url", "source_name", "updated_at"}),
	}).Create(&adrs)

	if result.Error != nil {
		return fmt.Errorf("failed to save ADRs: %w", result.Error)
	}

	return nil
}

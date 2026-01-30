package fetcher

import (
	"agora/pkg/config"

	"gorm.io/gorm"
)

// ADR holds the data of a single Architecture Decision Record.
type ADR struct {
	gorm.Model
	Title   string `gorm:"index"`
	Status  string
	Content string `gorm:"type:text"`
	URL     string `gorm:"uniqueIndex"`
}

// Fetcher is the interface for fetching ADRs from a source.
type Fetcher interface {
	Fetch(source config.Source) ([]ADR, error)
}

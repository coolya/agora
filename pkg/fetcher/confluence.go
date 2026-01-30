package fetcher

import (
	"agora/pkg/config"
	"agora/pkg/domain"
	"fmt"
)

// ConfluenceFetcher fetches ADRs from a Confluence space.
type ConfluenceFetcher struct{}

// NewConfluenceFetcher creates a new ConfluenceFetcher.
func NewConfluenceFetcher() *ConfluenceFetcher {
	return &ConfluenceFetcher{}
}

// Fetch retrieves ADRs from the configured Confluence space.
func (f *ConfluenceFetcher) Fetch(source config.Source) ([]domain.ADR, error) {
	return nil, fmt.Errorf("confluence fetcher not implemented")
}

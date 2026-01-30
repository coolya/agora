package fetcher

import (
	"agora/pkg/config"
	"agora/pkg/domain"
)

// Fetcher is the interface for fetching ADRs from a source.
type Fetcher interface {
	Fetch(source config.Source) ([]domain.ADR, error)
}

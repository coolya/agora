package fetcher

import "agora-aggregator/pkg/config"

// ADR holds the data of a single Architecture Decision Record.
type ADR struct {
	Title   string
	Status  string
	Content string
	URL     string
}

// Fetcher is the interface for fetching ADRs from a source.
type Fetcher interface {
	Fetch(source config.Source) ([]ADR, error)
}

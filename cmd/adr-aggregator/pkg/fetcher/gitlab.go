package fetcher

import (
	"adr-aggregator/pkg/config"
	"fmt"
)

// GitLabFetcher fetches ADRs from a GitLab repository.
type GitLabFetcher struct{}

// NewGitLabFetcher creates a new GitLabFetcher.
func NewGitLabFetcher() *GitLabFetcher {
	return &GitLabFetcher{}
}

// Fetch retrieves ADRs from the configured GitLab repository.
func (f *GitLabFetcher) Fetch(source config.Source) ([]ADR, error) {
	return nil, fmt.Errorf("gitlab fetcher not implemented")
}

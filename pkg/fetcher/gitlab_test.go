package fetcher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGitLabFetcher(t *testing.T) {
	fetcher, err := NewGitLabFetcher("fake-token", "")
	assert.NoError(t, err)
	assert.NotNil(t, fetcher)
}

func TestGitLabFetcher_Fetch(t *testing.T) {
	t.Skip("Test requires proper mock server setup for GitLab API client library")
	// This test is skipped because it requires a properly mocked GitLab API
	// The actual integration testing should be done with a real GitLab instance
}

func TestParseGitLabURL(t *testing.T) {
	projectID, err := parseGitLabURL("https://gitlab.com/owner/repo")
	assert.NoError(t, err)
	assert.Equal(t, "owner/repo", projectID)
}

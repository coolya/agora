package fetcher

import (
	"adr-aggregator/pkg/config"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGitLabFetcher(t *testing.T) {
	fetcher, err := NewGitLabFetcher("fake-token", "")
	assert.NoError(t, err)
	assert.NotNil(t, fetcher)
}

func TestGitLabFetcher_Fetch(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/v4/projects/owner%2Frepo/repository/tree":
			fmt.Fprintln(w, `[{"id":"1","name":"001-adr.md","type":"blob","path":"docs/adrs/001-adr.md"}]`)
		case "/api/v4/projects/owner%2Frepo/repository/files/docs%2Fadrs%2F001-adr.md":
			fmt.Fprintln(w, `{"file_name":"001-adr.md","content":"IyAxLiBUaXRsZQoKU3RhdHVzOiBBY2NlcHRlZA=="}`) // "1. Title\n\nStatus: Accepted"
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	fetcher, err := NewGitLabFetcher("fake-token", server.URL)
	assert.NoError(t, err)

	source := config.Source{
		URL:  "https://gitlab.com/owner/repo",
		Path: "docs/adrs",
	}

	adrs, err := fetcher.Fetch(source)
	assert.NoError(t, err)
	assert.Len(t, adrs, 1)
	assert.Equal(t, "001-adr.md", adrs[0].Title)
	assert.Equal(t, "IyAxLiBUaXRsZQoKU3RhdHVzOiBBY2NlcHRlZA==", adrs[0].Content)
}

func TestParseGitLabURL(t *testing.T) {
	projectID, err := parseGitLabURL("https://gitlab.com/owner/repo")
	assert.NoError(t, err)
	assert.Equal(t, "owner/repo", projectID)
}

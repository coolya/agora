package fetcher

import (
	"adr-aggregator/pkg/config"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/go-github/v63/github"
	"golang.org/x/oauth2"
)

// GitHubFetcher fetches ADRs from a GitHub repository.
type GitHubFetcher struct {
	client *github.Client
}

// NewGitHubFetcher creates a new GitHubFetcher.
func NewGitHubFetcher(token string) *GitHubFetcher {
	var tc *http.Client
	if token != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc = oauth2.NewClient(context.Background(), ts)
	}
	client := github.NewClient(tc)
	return &GitHubFetcher{client: client}
}

// Fetch retrieves ADRs from the configured GitHub repository.
func (f *GitHubFetcher) Fetch(source config.Source) ([]ADR, error) {
	owner, repo, err := parseGitHubURL(source.URL)
	if err != nil {
		return nil, err
	}

	_, directoryContent, _, err := f.client.Repositories.GetContents(
		context.Background(),
		owner,
		repo,
		source.Path,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get contents of %s: %w", source.Path, err)
	}

	var adrs []ADR
	for _, item := range directoryContent {
		if strings.HasSuffix(*item.Name, ".md") {
			fileContent, _, _, err := f.client.Repositories.GetContents(
				context.Background(),
				owner,
				repo,
				*item.Path,
				nil,
			)
			if err != nil {
				return nil, fmt.Errorf("failed to get content of %s: %w", *item.Path, err)
			}
			content, err := fileContent.GetContent()
			if err != nil {
				return nil, fmt.Errorf("failed to decode content of %s: %w", *item.Path, err)
			}
			adrs = append(adrs, ADR{
				Title:   *item.Name,
				Content: content,
				URL:     *item.HTMLURL,
			})
		}
	}

	return adrs, nil
}

func parseGitHubURL(rawURL string) (owner, repo string, err error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", "", fmt.Errorf("failed to parse url: %w", err)
	}

	if parsedURL.Hostname() != "github.com" {
		return "", "", fmt.Errorf("not a github url: %s", rawURL)
	}

	parts := strings.Split(strings.TrimPrefix(parsedURL.Path, "/"), "/")
	if len(parts) < 2 {
		return "", "", fmt.Errorf("invalid github url path: %s", rawURL)
	}

	owner = parts[0]
	repo = parts[1]
	return
}

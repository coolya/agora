package fetcher

import (
	"adr-aggregator/pkg/config"
	"fmt"
	"net/url"
	"strings"

	"gitlab.com/gitlab-org/api/client-go"
)

// GitLabFetcher fetches ADRs from a GitLab repository.
type GitLabFetcher struct {
	client *gitlab.GitLabClient
}

// NewGitLabFetcher creates a new GitLabFetcher.
func NewGitLabFetcher(token, baseURL string) (*GitLabFetcher, error) {
	var options []gitlab.OptionFunc
	if baseURL != "" {
		options = append(options, gitlab.WithBaseURL(baseURL))
	}
	gclient, err := gitlab.New(token, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to create gitlab client: %w", err)
	}
	return &GitLabFetcher{client: gclient}, nil
}

// Fetch retrieves ADRs from the configured GitLab repository.
func (f *GitLabFetcher) Fetch(source config.Source) ([]ADR, error) {
	projectID, err := parseGitLabURL(source.URL)
	if err != nil {
		return nil, err
	}

	tree, _, err := f.client.Repositories.ListTree(projectID, &gitlab.ListTreeOptions{
		Path: &source.Path,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list tree for %s: %w", projectID, err)
	}

	var adrs []ADR
	for _, node := range tree {
		if strings.HasSuffix(node.Name, ".md") {
			ref := "main" // Assuming main branch
			file, _, err := f.client.RepositoryFiles.GetFile(projectID, node.Path, &gitlab.GetFileOptions{
				Ref: &ref,
			})
			if err != nil {
				return nil, fmt.Errorf("failed to get file %s: %w", node.Path, err)
			}
			adrs = append(adrs, ADR{
				Title:   node.Name,
				Content: file.Content,
				URL:     fmt.Sprintf("%s/blob/main/%s", source.URL, node.Path),
			})
		}
	}

	return adrs, nil
}

func parseGitLabURL(rawURL string) (projectID string, err error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("failed to parse url: %w", err)
	}

	projectID = strings.TrimPrefix(parsedURL.Path, "/")
	return
}

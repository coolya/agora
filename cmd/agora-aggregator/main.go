package main

import (
	"agora/pkg/config"
	"agora/pkg/fetcher"
	"agora/pkg/parser"
	"agora/pkg/storage"
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "agora-aggregator",
	Short: "A tool to aggregate Architecture Decision Records.",
	Run:   run,
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		os.Exit(1)
	}

	// Initialize the database
	err = storage.InitDB("adrs.db")
	if err != nil {
		fmt.Printf("Error initializing database: %v\n", err)
		os.Exit(1)
	}

	var allADRs []fetcher.ADR

	for _, source := range cfg.Sources {
		var f fetcher.Fetcher
		switch source.Type {
		case "github":
			f = fetcher.NewGitHubFetcher(source.Auth.Token)
		case "gitlab":
			parsedURL, err := url.Parse(source.URL)
			if err != nil {
				fmt.Printf("Error parsing GitLab URL %s: %v\n", source.URL, err)
				continue
			}
			baseURL := parsedURL.Scheme + "://" + parsedURL.Host
			f, err = fetcher.NewGitLabFetcher(source.Auth.Token, baseURL)
			if err != nil {
				fmt.Printf("Error creating GitLab fetcher for %s: %v\n", source.URL, err)
				continue
			}
		case "confluence":
			f = fetcher.NewConfluenceFetcher()
		default:
			fmt.Printf("Unknown source type: %s (URL: %s)\n", source.Type, source.URL)
			continue
		}

		adrs, err := f.Fetch(source)
		if err != nil {
			fmt.Printf("Error fetching from %s: %v\n", source.URL, err)
			continue
		}

		for i := range adrs {
			parser.Parse(&adrs[i])
		}

		allADRs = append(allADRs, adrs...)
	}

	// Save ADRs to database
	if len(allADRs) > 0 {
		err = storage.SaveADRs(allADRs)
		if err != nil {
			fmt.Printf("Error saving ADRs to database: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Successfully aggregated %d ADRs to adrs.db\n", len(allADRs))
	} else {
		fmt.Println("No ADRs were fetched from any source")
	}
}

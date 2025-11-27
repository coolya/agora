package main

import (
	"adr-aggregator/pkg/config"
	"adr-aggregator/pkg/fetcher"
	"adr-aggregator/pkg/parser"
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "adr-aggregator",
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
			fmt.Printf("Unknown source type: %s\n", source.Type)
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

	jsonData, err := json.MarshalIndent(allADRs, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling to JSON: %v\n", err)
		os.Exit(1)
	}

	err = os.WriteFile("adrs.json", jsonData, 0644)
	if err != nil {
		fmt.Printf("Error writing to adrs.json: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Successfully aggregated ADRs to adrs.json")
}

package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "agora",
	Short: "A tool to aggregate, query, and serve Architecture Decision Records.",
	Long: `agora is a command-line tool for aggregating Architecture Decision Records (ADRs) 
from various sources (GitHub, GitLab, Confluence), querying them, and serving them via a REST API.`,
}

func init() {
	rootCmd.AddCommand(aggregateCmd)
	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(queryCmd)
}

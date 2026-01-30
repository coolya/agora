package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query ADRs from the database",
	Long:  `Query and filter Architecture Decision Records stored in the database.`,
	Run:   runQuery,
}

func runQuery(cmd *cobra.Command, args []string) {
	fmt.Println("Query command is not yet implemented")
}

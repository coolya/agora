package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the REST API server",
	Long:  `Start a REST API server to query and retrieve ADRs from the database.`,
	Run:   runServe,
}

func runServe(cmd *cobra.Command, args []string) {
	fmt.Println("Serve command is not yet implemented")
}

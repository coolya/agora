This document provides guidance for AI agents working on this project.

## Project Overview

**agora** is a monolithic command-line tool written in Go that provides multiple functionalities for managing Architecture Decision Records (ADRs):
- **Aggregation**: Fetch ADRs from various sources (GitHub, GitLab, Confluence) and store them in a SQLite database
- **Querying**: Query and filter ADRs from the database (planned)
- **Serving**: REST API for accessing ADRs programmatically (planned)

The Go module is named `agora` and resides at the root level.

## Technology Stack

- **Go**: The primary programming language. The Go version is specified in `go.mod`.
- **Cobra**: A library for creating powerful modern CLI applications in Go with subcommands.
- **Viper**: A Go configuration library that handles configuration from files, environment variables, and flags.
- **GORM**: An ORM library for database operations with SQLite backend.

## Project Structure

```
cmd/agora/              - Main monolithic CLI application
  ├── main.go          - Entry point that executes the root command
  ├── root.go          - Root command definition with subcommand registration
  ├── aggregate.go     - "agora aggregate" subcommand implementation
  ├── serve.go         - "agora serve" subcommand (REST API server, planned)
  ├── query.go         - "agora query" subcommand (database querying, planned)
  └── aggregate_test.go - Tests for the aggregate functionality

pkg/                    - Shared packages used across all commands
  ├── config/          - Configuration loading and management
  ├── fetcher/         - Logic for fetching ADRs from different sources
  ├── parser/          - ADR parsing and status extraction
  └── storage/         - SQLite database persistence using GORM

web/                    - Web UI for visualizing and interacting with ADRs (planned)
config.yaml            - Configuration file specifying ADR sources
go.mod, go.sum         - Go module files managing dependencies
```

## Getting Started

### Prerequisites

- Go (version specified in `go.mod`)

### Building the Application

To build the `agora` tool from the repository root:

```bash
go build ./cmd/agora/
```

### Running the Application

The `agora` tool is a multi-purpose CLI with several subcommands, all configured via `config.yaml`.

#### Aggregating ADRs

Before running aggregation, set the `GITHUB_TOKEN` environment variable if using GitHub sources:

```bash
export GITHUB_TOKEN="your_github_token"
```

To aggregate ADRs from configured sources:

```bash
./agora aggregate
```

#### Future Subcommands

- `agora serve`: Start a REST API server (planned)
- `agora query`: Query ADRs from the database (planned)

### Testing

To run all tests from the repository root:

```bash
go test ./...
```

## Common Development Tasks

### Adding a New Data Source

1. **Create a new fetcher**: In the `pkg/fetcher/` directory, create a new file (e.g., `my_source.go`) that implements the `Fetcher` interface.
2. **Implement the `Fetch` method**: The `Fetch` method should contain the logic for fetching ADRs from the new data source.
3. **Update `aggregate.go`**: In `cmd/agora/aggregate.go`, add a new case to the `switch` statement in the `runAggregate` function to handle the new source type.
4. **Update `config.yaml`**: Add a new entry to the `sources` list with configuration for the new data source.
5. **Write tests**: Consider adding unit tests for the new fetcher.

### Adding a New Subcommand

1. **Create a new command file**: In `cmd/agora/`, create a new file (e.g., `mycommand.go`) for your subcommand.
2. **Define the cobra.Command**: Create a `var myCmd = &cobra.Command{...}` with your command logic.
3. **Register the command**: Add `rootCmd.AddCommand(myCmd)` in the `init()` function in `root.go`.
4. **Add tests**: Create corresponding test files (e.g., `mycommand_test.go`).

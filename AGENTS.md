This document provides guidance for AI agents working on this project.

## Project Overview

The `adr-aggregator` is a command-line tool written in Go that aggregates Architecture Decision Records (ADRs) from various sources, such as GitHub, GitLab, and Confluence. The tool fetches ADRs, parses them, and compiles them into a single `adrs.json` file.

## Technology Stack

- **Go**: The primary programming language. The Go version is specified in `cmd/adr-aggregator/go.mod`.
- **Cobra**: A library for creating powerful modern CLI applications in Go.
- **Viper**: A Go configuration library that handles configuration from files, environment variables, and flags.

## Project Structure

- `cmd/adr-aggregator/`: The main application directory.
  - `main.go`: The entry point of the application.
  - `pkg/`: Contains the core logic of the application.
    - `config/`: Configuration loading and management.
    - `fetcher/`: Logic for fetching ADRs from different sources.
    - `parser/`: ADR parsing logic.
  - `go.mod`, `go.sum`: Go module files that manage dependencies.
- `config.yaml`: The configuration file for the `adr-aggregator` tool.

## Getting Started

### Prerequisites

- Go (version specified in `cmd/adr-aggregator/go.mod`)

### Building the Application

To build the `adr-aggregator` tool, navigate to the `cmd/adr-aggregator` directory and run the following command:

```bash
go build
```

### Running the Application

The `adr-aggregator` tool is configured using `config.yaml`. The `config.yaml` file specifies the sources from which to fetch ADRs.

Before running the application, you may need to set the `GITHUB_TOKEN` environment variable to authenticate with the GitHub API:

```bash
export GITHUB_TOKEN="your_github_token"
```

To run the tool, execute the following command from the repository root:

```bash
./cmd/adr-aggregator/adr-aggregator
```

### Testing

To run the tests, navigate to the `cmd/adr-aggregator` directory and run:

```bash
go test ./...
```

## Common Development Tasks

### Adding a New Data Source

1.  **Create a new fetcher**: In the `pkg/fetcher` directory, create a new file (e.g., `my_new_source_fetcher.go`) that implements the `Fetcher` interface.
2.  **Implement the `Fetch` method**: The `Fetch` method should contain the logic for fetching ADRs from the new data source.
3.  **Update `main.go`**: In `cmd/adr-aggregator/main.go`, add a new case to the `switch` statement in the `run` function to handle the new source type.
4.  **Update `config.yaml`**: Add a new entry to the `sources` list in `config.yaml` with the configuration for the new data source.
5.  **Write tests**: Add unit tests for the new fetcher.

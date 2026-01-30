# agora
A tool for aggregating architecture decision records from different places and fostering a conversational decision making process.

## Architecture

The project uses a root module structure (`agora`) with a monolithic CLI application that supports multiple subcommands for different functionalities:

### Directory Structure
- **`cmd/agora/`**: The main monolithic CLI application with subcommands
  - `main.go`: Entry point
  - `root.go`: Root command definition
  - `aggregate.go`: ADR aggregation subcommand
  - `serve.go`: REST API server subcommand (planned)
  - `query.go`: ADR query subcommand (planned)
- **`pkg/config/`**: Configuration loading and management
- **`pkg/fetcher/`**: ADR fetching logic for different sources (GitHub, GitLab, Confluence)
- **`pkg/parser/`**: ADR parsing and status extraction
- **`pkg/storage/`**: SQLite database persistence layer using GORM
- **`web/`**: Web UI for visualizing and interacting with ADRs (planned)

## Getting Started

### Prerequisites

- Go (version specified in `go.mod`)

### Building the Application

To build the `agora` tool, run the following command from the repository root:

```bash
go build ./cmd/agora/
```

### Running the Application

The `agora` tool is a multi-purpose CLI with several subcommands. It is configured using `config.yaml`, which specifies the sources from which to fetch ADRs.

#### Aggregating ADRs

Before running the aggregation, you may need to set the `GITHUB_TOKEN` environment variable to authenticate with the GitHub API:

```bash
export GITHUB_TOKEN="your_github_token"
```

To aggregate ADRs from configured sources:

```bash
./agora aggregate
```

The aggregator will fetch ADRs from the configured sources and store them in a SQLite database (`adrs.db`).

#### Future Subcommands

- `agora serve`: Start a REST API server (planned)
- `agora query`: Query ADRs from the database (planned)

### Testing

To run all tests, execute from the repository root:

```bash
go test ./...
```

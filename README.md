# agora
A tool for aggregating architecture decision records from different places and fostering a conversational decision making process.

## Architecture

The project uses a root module structure (`agora`) that enables code sharing between the aggregator and future API services:

- **`cmd/agora-aggregator/`**: The ADR aggregator CLI application
- **`pkg/config/`**: Configuration loading and management
- **`pkg/fetcher/`**: ADR fetching logic for different sources (GitHub, GitLab, Confluence)
- **`pkg/parser/`**: ADR parsing and status extraction
- **`pkg/storage/`**: SQLite database persistence layer using GORM

## Getting Started

### Prerequisites

- Go (version specified in `go.mod`)

### Building the Application

To build the `agora-aggregator` tool, run the following command from the repository root:

```bash
go build ./cmd/agora-aggregator/
```

### Running the Application

The `agora-aggregator` tool is configured using `config.yaml`. The `config.yaml` file specifies the sources from which to fetch ADRs.

Before running the application, you may need to set the `GITHUB_TOKEN` environment variable to authenticate with the GitHub API:

```bash
export GITHUB_TOKEN="your_github_token"
```

To run the tool, execute the following command from the repository root:

```bash
./cmd/agora-aggregator/agora-aggregator
```

The aggregator will fetch ADRs from the configured sources and store them in a SQLite database (`adrs.db`).

### Testing

To run all tests, execute from the repository root:

```bash
go test ./...
```

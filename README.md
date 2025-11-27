# agora
A tool for aggregating architecture decision records from different places and fostering a conversational  decision making process. 

## Getting Started

### Prerequisites

- Go (version specified in `cmd/agora-aggregator/go.mod`)

### Building the Application

To build the `agora-aggregator` tool, navigate to the `cmd/agora-aggregator` directory and run the following command:

```bash
go build
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

### Testing

To run the tests, navigate to the `cmd/agora-aggregator` directory and run:

```bash
go test ./...
```

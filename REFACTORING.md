# Architecture Refactoring Summary

## Overview
The agora project has been successfully refactored from a single-purpose aggregator CLI to a monolithic multi-command CLI application. This enables code sharing and provides a foundation for planned features like REST API serving and database querying.

## Changes Made

### 1. New Directory Structure

#### `cmd/agora/` - Main CLI Application
Replaced the single `cmd/agora-aggregator/` with a new modular structure:

- **`main.go`**: Entry point that executes the root command
- **`root.go`**: Root command definition with subcommand registration
- **`aggregate.go`**: Implementation of `agora aggregate` command (ported from aggregator)
- **`aggregate_test.go`**: Integration tests for the aggregate functionality
- **`serve.go`**: Stub implementation of `agora serve` command (future REST API)
- **`query.go`**: Stub implementation of `agora query` command (future database querying)

#### `web/README.md`
New directory reserved for future web UI implementation that will consume the REST API.

### 2. Removed Components

- **`cmd/agora-aggregator/`**: Entire directory removed
  - `main.go` - merged into `cmd/agora/`
  - `main_test.go` - migrated to `cmd/agora/aggregate_test.go`
  - Binary no longer needed

### 3. Documentation Updates

#### `README.md`
- Updated architecture section to describe the monolithic CLI structure
- Added explicit directory structure with descriptions
- Updated build command: `go build ./cmd/agora/`
- Updated run command: `./agora aggregate`
- Added documentation for future subcommands (serve, query)

#### `AGENTS.md`
- Fixed incorrect `go.mod` path references (was pointing to `cmd/agora-aggregator/go.mod`)
- Updated project overview to describe the new monolithic multi-command architecture
- Updated technology stack section
- Completely rewrote project structure documentation with ASCII tree
- Updated building, running, and testing instructions
- Added section on "Adding a New Subcommand" for future development
- Updated the "Adding a New Data Source" section to reference correct file paths

### 4. CLI Command Interface

The application now provides a single `agora` binary with subcommands:

```bash
# Aggregate ADRs from configured sources
agora aggregate

# Start REST API server (planned)
agora serve

# Query ADRs from database (planned)
agora query
```

All commands use the existing shared libraries:
- `pkg/config/` - Configuration management
- `pkg/fetcher/` - Data source fetching
- `pkg/parser/` - ADR parsing
- `pkg/storage/` - Database persistence

## Benefits

1. **Single Binary**: Users only need to distribute and run one CLI tool
2. **Code Reuse**: All commands share the same config, fetcher, parser, and storage packages
3. **Extensibility**: New subcommands can be easily added following the established pattern
4. **Clear Architecture**: The monolithic structure with subcommands is a common Go pattern
5. **Foundation for Web UI**: The `web/` directory and planned REST API support visualization
6. **Better Documentation**: AI agents now have clearer guidance on the project structure

## Testing

All existing tests continue to pass:
- `cmd/agora` integration tests
- `pkg/config` tests
- `pkg/fetcher` tests (with GitLab mock skipped as before)
- `pkg/parser` tests
- `pkg/storage` tests

Test execution:
```bash
go test ./...
```

## Build and Deployment

Building the application:
```bash
go build ./cmd/agora/
```

This creates a single `agora` binary that can be deployed to users or CI/CD systems.

## Future Work

1. **REST API Server** (`agora serve`):
   - Implement HTTP server exposing ADR data via REST endpoints
   - Support filtering and search capabilities

2. **Query Command** (`agora query`):
   - Implement CLI-based querying of the database
   - Provide filter options for status, source, date ranges, etc.

3. **Web UI** (`web/`):
   - Build web interface consuming the REST API
   - Provide visual exploration and filtering of ADRs
   - Support collaborative decision-making features

## Migration Notes

Users of the old `agora-aggregator` should:
1. Rebuild using `go build ./cmd/agora/`
2. Update any documentation or scripts to use `./agora aggregate` instead of `./cmd/agora-aggregator/agora-aggregator`
3. All configuration and functionality remains the same

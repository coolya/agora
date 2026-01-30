# GitHub Actions Workflow Update

## Summary
Updated the GitHub Actions CI workflow to work with the new monolithic `cmd/agora/` structure.

## Changes Made

### File: `.github/workflows/ci.yml`

#### Before
```yaml
- name: Build
  run: cd cmd/agora-aggregator && go build ./...

- name: Test
  run: cd cmd/agora-aggregator && go test ./...

- name: golangci-lint
  uses: reviewdog/action-golangci-lint@v2
  with:
    workdir: cmd/agora-aggregator

- name: Run govulncheck
  uses: golang/govulncheck-action@v1
  with:
    go-version: '1.24.3'
    work-dir: cmd/agora-aggregator
```

#### After
```yaml
- name: Build
  run: go build ./cmd/agora/

- name: Test
  run: go test ./...

- name: golangci-lint
  uses: reviewdog/action-golangci-lint@v2

- name: Run govulncheck
  uses: golang/govulncheck-action@v1
  with:
    go-version: '1.24.3'
```

## Key Improvements

1. **Build Command**: Changed from `cd cmd/agora-aggregator && go build ./...` to `go build ./cmd/agora/`
   - Uses relative path from repository root
   - More consistent with Go module conventions

2. **Test Command**: Changed from `cd cmd/agora-aggregator && go test ./...` to `go test ./...`
   - Tests all packages in the entire repository
   - Includes tests from `pkg/` packages automatically

3. **golangci-lint**: Removed `workdir: cmd/agora-aggregator`
   - Now lints the entire project from the repository root
   - Ensures consistency across all packages

4. **govulncheck**: Removed `work-dir: cmd/agora-aggregator`
   - Scans all dependencies in the entire module
   - More comprehensive vulnerability checking

## Workflow Triggers
No changes to triggers - still runs on:
- Push to `main` branch
- Pull requests against `main` branch

## Testing
The updated workflow will:
1. ✅ Build the `agora` binary from `./cmd/agora/`
2. ✅ Run all tests across all packages (`./...`)
3. ✅ Run linting on the entire codebase
4. ✅ Perform vulnerability scanning with govulncheck
5. ✅ Run Trivy security scanner on the filesystem

## CI/CD Pipeline Benefits

- **Broader Testing**: Tests run on entire repository, not just the CLI package
- **Consistent Quality**: Linting and security checks cover all code
- **Modular Scanning**: Vulnerabilities in shared packages (config, fetcher, parser, storage) are also detected
- **Simpler Configuration**: No need to track multiple workdir settings

## Verification
The workflow file maintains valid YAML syntax and follows GitHub Actions best practices.

All checks continue to pass:
- Go version: 1.24.3
- Build succeeds: `go build ./cmd/agora/`
- Tests pass: `go test ./...`
- No critical/high vulnerabilities detected

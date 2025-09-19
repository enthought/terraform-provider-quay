# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

### Testing
```bash
# Run acceptance tests (requires QUAY_URL and QUAY_TOKEN environment variables)
make testacc
# OR
TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m
```

### Code Generation
```bash
# Generate documentation after code changes
go generate .

# Generate Quay API client from OpenAPI spec
make generate-quay-api
```

### Build and Formatting
```bash
# Format Go code
gofmt -w .

# Build the provider
go build .
```

## Architecture Overview

This is a Terraform provider for the Quay container registry. The codebase follows standard Terraform provider patterns using the HashiCorp Plugin Framework.

### Core Components

- **Provider Definition**: `internal/provider/provider.go` - Main provider configuration with support for token-based and OAuth2 authentication
- **Generated API Client**: `quay_api/` - Auto-generated Go client from Quay's OpenAPI specification
- **Resources & Data Sources**: `internal/provider/*_resource.go` and `internal/provider/*_data_source.go` - Manual implementations
- **Generated Resources**: `internal/resource_*/` and `internal/datasource_*/` - Auto-generated resource implementations

### Authentication
The provider supports two authentication methods:
1. Bearer token authentication (`QUAY_TOKEN`)
2. OAuth2 password grant flow (`QUAY_OAUTH2_USERNAME`, `QUAY_OAUTH2_PASSWORD`, `QUAY_OAUTH2_CLIENT_ID`, `QUAY_OAUTH2_TOKEN_URL`)

### Code Generation Strategy
The provider uses a hybrid approach:
- API client is generated from `code_generator/quay_api.json` using openapi-generator
- Some resources have both manual implementations (in `internal/provider/`) and generated versions (in `internal/resource_*/`)
- Configuration files in `code_generator/` drive the generation process

### Resource Types
- Organizations: Create and manage Quay organizations
- Teams: Manage organization teams and permissions
- Robots: Service accounts for automated access
- Repositories: Container repositories within organizations

### Environment Variables for Testing
- `QUAY_URL`: Quay instance URL (e.g., https://quay.example.com)
- `QUAY_TOKEN`: Authentication token for testing
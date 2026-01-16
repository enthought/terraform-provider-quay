# terraform-provider-quay
Terraform provider for the [Quay Project](https://github.com/quay/quay).

Please visit the Terraform registry site for instructions on how to use the provider:

https://registry.terraform.io/providers/enthought/quay/latest/docs

## Developer Instructions
### Build Documentation
The documentation should be updated every time the provider code is changed.
```bash
go generate .
```

### Run Tests
```bash
export QUAY_URL="https://quay.example.com"
export QUAY_TOKEN=""
make testacc
```

### Update Quay API Spec
The Go API client is generated from an OpenAPI 3.0.1 specification stored in `code_generator/quay_api.json`. Quay instances expose their API spec in Swagger 2.0 format at `/api/v1/discovery`, so updating the spec requires fetching and converting it.

**Important:** Different Quay instances expose different API endpoints. Self-hosted Quay instances typically include superuser endpoints that are not available on quay.io. The current spec was generated from a self-hosted instance and the provider relies on these superuser endpoints for some operations (e.g., organization management).

**Prerequisites:** podman (or docker), curl, jq

To update the API spec from a Quay instance:
```bash
make update-quay-api-spec QUAY_INSTANCE=https://quay.example.com
```

This target:
1. Starts a swagger-converter container
2. Fetches the Swagger 2.0 spec from the Quay instance
3. Converts it to OpenAPI 3.0.1
4. Removes OAuth2 security definitions (not used by this provider)
5. Saves the result to `code_generator/quay_api.json`

After updating the spec, regenerate the Go client and verify the provider still builds:
```bash
make generate-quay-api
go build .
```

### Generate Quay API Client
Regenerate the Go client after spec changes:
```bash
make generate-quay-api
```

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

**Prerequisites:** podman (or docker), curl, jq

To update the API spec from quay.io:
```bash
make update-quay-api-spec
```

To update from a different Quay instance:
```bash
make update-quay-api-spec QUAY_INSTANCE=https://quay.example.com
```

This target:
1. Starts a swagger-converter container
2. Fetches the Swagger 2.0 spec from the Quay instance
3. Converts it to OpenAPI 3.0.1
4. Removes OAuth2 security definitions (not used by this provider)
5. Saves the result to `code_generator/quay_api.json`

### Generate Quay API Client
After updating the API spec, regenerate the Go client:
```bash
make generate-quay-api
```

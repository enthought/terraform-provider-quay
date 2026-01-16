# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

# Modifications copyright (c) Enthought, Inc.
# SPDX-License-Identifier:	BSD-3-Clause

default: testacc

# Run acceptance tests
.PHONY: testacc
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

# Generate Quay API
.PHONY: generate-quay-api
generate-quay-api:
	openapi-generator generate --generator-name go --git-user-id enthought --git-repo-id terraform-provider-quay --additional-properties=packageName=quay_api,isGoSubmodule=true --skip-validate-spec --input-spec code_generator/quay_api.json --output quay_api && \
    gofmt -w quay_api && \
    cd quay_api && \
    go mod tidy

# Update Quay API spec from a Quay instance
# This fetches the Swagger 2.0 spec from Quay and converts it to OpenAPI 3.0.1
# Prerequisites: podman (or docker), curl, jq
# Usage: make update-quay-api-spec QUAY_INSTANCE=https://quay.io
QUAY_INSTANCE ?= https://quay.io
.PHONY: update-quay-api-spec
update-quay-api-spec:
	@echo "Starting swagger-converter..."
	podman run -d --rm -p 8080:8080 --name swagger-converter docker.io/swaggerapi/swagger-converter:v1.0.5
	@sleep 2
	@echo "Converting Quay API spec from Swagger 2.0 to OpenAPI 3.0.1..."
	curl --silent --data "$$(curl --silent --location $(QUAY_INSTANCE)/api/v1/discovery)" \
		-H "Content-Type: application/json" \
		http://localhost:8080/api/convert | \
		jq 'del(.. | .security?) | del(.components.securitySchemes)' \
		> code_generator/quay_api.json
	podman stop swagger-converter
	@echo "Done! quay_api.json updated from $(QUAY_INSTANCE)."

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

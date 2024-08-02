#!/bin/bash

openapi-generator generate --generator-name go --git-user-id enthought --git-repo-id terraform-provider-quay --additional-properties=packageName=quay_api,isGoSubmodule=true --skip-validate-spec --input-spec code_generator/quay_api.json --output quay_api
gofmt -w quay_api
cd quay_api
go mod tidy

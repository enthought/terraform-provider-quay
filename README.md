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

### Generate Quay API
```bash
make generate-quay-api
```

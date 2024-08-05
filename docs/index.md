---
page_title: "Quay Provider"
subcategory: ""
description: |-
    Terraform provider for the [Quay Project](https://github.com/quay/quay).
---

# Quay Provider
The Quay provider provides resources and data sources for managing Quay organizations and repositories.

## Authentication
An authentication token can be generated in the Quay application settings page. The Quay documentation refers to it as
an `OAuth 2 Access Token`. Since there is no mechanism in Quay for creating applications outside of an organization, it
is recommended to create a separate organization and application for Terraform. Be sure to create the
`OAuth 2 Access Token` using a service account since the token will be tied to that account.

## Example Usage

```terraform
provider "quay" {
  url = "https://quay.example.com"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `token` (String, Sensitive) Quay token. May also be provided via the QUAY_TOKEN environment variable.
- `url` (String) Quay URL. May also be provided via the QUAY_URL environment variable. Example: https://quay.example.com

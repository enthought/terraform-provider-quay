---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "quay_repository Data Source - quay"
subcategory: ""
description: |-
  
---

# quay_repository (Data Source)



## Example Usage

```terraform
data "quay_organization" "main" {
  name = "main"
}

data "quay_repository" "test" {
  name      = "test"
  namespace = data.quay_organization.main.name
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Repository name
- `namespace` (String) Repository namespace. Should be an organization name or username

### Read-Only

- `description` (String) Markdown description
- `visibility` (String) Repository visibility. Should be private or public.
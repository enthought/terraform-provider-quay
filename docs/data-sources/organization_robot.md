---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "quay_organization_robot Data Source - quay"
subcategory: ""
description: |-
  
---

# quay_organization_robot (Data Source)



## Example Usage

```terraform
resource "quay_organization" "org" {
  name  = "org"
  email = "quay+org@example.com"
}

data "quay_organization_robot" "robot" {
  name    = "robot"
  orgname = quay_organization.org.name
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Robot short name
- `orgname` (String) Organization name

### Read-Only

- `description` (String) Text description
- `fullname` (String) Robot full name

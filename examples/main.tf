terraform {
  required_providers {
    quay = {
      source = "registry.terraform.io/enthought/quay"
    }
  }
}

provider "quay" {
  url = "https://quay.example.com"
}

resource "quay_organization" "main" {
  name  = "main"
  email = "quay+main@example.com"
}

resource "quay_organization_robot" "test" {
  name    = "test"
  orgname = quay_organization.main.name
}

resource "quay_organization_team" "admin" {
  name    = "admin"
  orgname = quay_organization.main.name
  role    = "admin"
  members = [
    quay_organization_robot.test.fullname
  ]
}

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

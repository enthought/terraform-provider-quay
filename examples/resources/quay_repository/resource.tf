resource "quay_organization" "main" {
  name  = "main"
  email = "quay+main@example.com"
}

resource "quay_repository" "test" {
  name        = "test"
  namespace   = quay_organization.main.name
  visibility  = "private"
  description = "test"
}

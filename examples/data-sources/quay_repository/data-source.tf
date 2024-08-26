data "quay_organization" "main" {
  name = "main"
}

data "quay_repository" "test" {
  name      = "test"
  namespace = data.quay_organization.main.name
}

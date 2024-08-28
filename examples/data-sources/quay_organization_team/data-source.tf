resource "quay_organization" "org" {
  name  = "org"
  email = "quay+org@example.com"
}

resource "quay_repository" "repo" {
  name      = "repo"
  namespace = quay_organization.org.name
}

resource "quay_organization_robot" "robot" {
  name    = "robot"
  orgname = quay_organization.org.name
}

data "quay_organization_team" "team" {
  name    = "team"
  orgname = quay_organization.org.name
}

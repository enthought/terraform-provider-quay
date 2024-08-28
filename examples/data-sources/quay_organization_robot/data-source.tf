resource "quay_organization" "org" {
  name  = "org"
  email = "quay+org@example.com"
}

data "quay_organization_robot" "robot" {
  name    = "robot"
  orgname = quay_organization.org.name
}

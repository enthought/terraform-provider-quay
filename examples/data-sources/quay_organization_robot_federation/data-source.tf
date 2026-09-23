resource "quay_organization" "org" {
  name  = "org"
  email = "quay+org@example.com"
}

data "quay_organization_robot_federation" "ci" {
  orgname   = quay_organization.org.name
  robotname = "ci"
}

output "robot_federation" {
  value = data.quay_organization_robot_federation.ci.federation
}

resource "quay_organization" "main" {
  name  = "main"
  email = "quay+main@example.com"
}

resource "quay_organization_robot" "test" {
  name    = "test"
  orgname = quay_organization.main.name
}

output "robot_token" {
  value     = quay_organization_robot.test.token
  sensitive = true
}

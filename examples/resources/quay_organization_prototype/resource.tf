resource "quay_organization" "main" {
  name  = "main"
  email = "quay+main@example.com"
}

resource "quay_organization_robot" "push" {
  name    = "push"
  orgname = quay_organization.main.name
}

resource "quay_organization_robot" "pull" {
  name    = "pull"
  orgname = quay_organization.main.name
}

# Grant the pull robot read access to every new repository created by the push robot.
resource "quay_organization_prototype" "test" {
  orgname         = quay_organization.main.name
  activating_user = quay_organization_robot.push.fullname
  delegate_name   = quay_organization_robot.pull.fullname
  delegate_kind   = "user"
  role            = "read"
}

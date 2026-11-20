resource "quay_organization" "main" {
  name  = "main"
  email = "quay+main@example.com"
}

resource "quay_organization_team" "creators" {
  name    = "creators"
  orgname = quay_organization.main.name
  role    = "creator"

  # quay_organization_team_member manages members outside this resource, so ignore them here.
  # Without this, Terraform plans to remove the members that quay_organization_team_member added.
  lifecycle {
    ignore_changes = [members]
  }
}

resource "quay_organization_robot" "push" {
  name    = "push"
  orgname = quay_organization.main.name
}

# Add a single robot to the team without managing the team's other members.
resource "quay_organization_team_member" "test" {
  orgname    = quay_organization.main.name
  teamname   = quay_organization_team.creators.name
  membername = quay_organization_robot.push.fullname
}

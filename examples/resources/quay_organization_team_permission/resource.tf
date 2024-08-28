resource "quay_organization" "org" {
  name  = "org"
  email = "quay+org@example.com"
}

resource "quay_repository" "repo" {
  name      = "repo"
  namespace = quay_organization.org.name
}

resource "quay_organization_team" "team" {
  name    = "team"
  orgname = quay_organization.org.name
  role    = "member"
}

resource "quay_organization_team_permission" "permission" {
  orgname    = quay_organization.org.name
  reponame   = quay_repository.repo.name
  teamname   = quay_organization_team.team.name
  permission = "read"
}

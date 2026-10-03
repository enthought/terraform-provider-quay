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

# A mirrored repository. The platform assigns the repository state to "mirror",
# pulls the configured tags from the upstream registry on the given interval, and
# uses the given robot account to authenticate. Trigger an immediate sync after
# (re)configuring the mirror by setting sync_on_change to true.
resource "quay_repository" "mirror" {
  name        = "mirror"
  namespace   = quay_organization.main.name
  visibility  = "private"
  description = "mirrored upstream images"
  state       = "mirror"

  mirror {
    upstream_repository = "quay.io/argoproj/argocd"
    robot_user          = "main+ci"
    sync_interval       = 86400
    sync_on_change      = true
    verify_tls          = true
    unsigned_images     = false
    root_rule           = { rule_kind = "tag_glob_csv", rule_value = ["*"] }
  }
}

# resource "quay_organization" "main" {
#   name  = "main"
#   email = "quay+main@example.com"
# }

resource "quay_organization_robot" "ci" {
  name    = "mshtest20260901"
  orgname = "apheris_try"
}

# Allow the robot to be authenticated (keyless) via OIDC tokens from an
# external identity provider instead of using its static token.
resource "quay_organization_robot_federation" "ci" {
  orgname   = "apheris_try"
  robotname = quay_organization_robot.ci.name

  federation = [
    {
      issuer  = "https://token.actions.githubusercontent.com"
      subject = "repo:my-org/my-repo:ref:refs/heads/main"
    },
    {
      issuer  = "https://accounts.google.com"
      subject = "1234567890"
    },
  ]
}

terraform {
  required_providers {
    quay = {
      source = "enthought/quay"
    version = "0.6.1" }
  }
}

provider "quay" {
  url = "https://quay.io"
  # token read from QUAY_TOKEN env var
}
